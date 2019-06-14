package vault

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/google/uuid"
	"github.com/hashicorp/vault/api"
	rError "github.com/johnjerrico/gokit-starter-pack/pkg/error"
)

// Vault ...
type Vault struct {
	*api.Client
}

// Response ...
type Response struct {
	Data map[string]interface{}
}

func getEnvOrDefault(env string, defaultVal string) string {
	e := os.Getenv(env)
	if e == "" {
		return defaultVal
	}
	return e
}

// New create new Vault connection
func New() (*Vault, error) {
	c, err := api.NewClient(nil)

	if err != nil {
		return nil, err
	}

	addr := getEnvOrDefault("VAULT_ADDR", "http://127.0.0.1:8200")
	token := getEnvOrDefault("VAULT_TOKEN", "")

	c.SetAddress(addr)
	c.SetToken(token)

	return &Vault{c}, nil
}

// GetEnvOrDefaultConfig get configuration from Vault
func (c *Vault) GetEnvOrDefaultConfig(path string, def interface{}) (map[string]string, error) {
	var err error
	res := make(map[string]string)

	if c == nil {
		return nil, rError.New(err, rError.Enum.INTERNALSERVERERROR, "client_has_not_been_initiated")
	}

	// Parsing interface to map
	s := reflect.ValueOf(def).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		res[typeOfT.Field(i).Name] = f.Interface().(string)
	}

	// Read Config from config/global vault
	kv, err := c.Logical().Read("config/global")

	if err != nil {
		return nil, err
	}

	for k := range res {
		if kv.Data[k] != nil {
			res[k] = kv.Data[k].(string)
		}
	}

	// Read Config from config/{path} vault
	kv, err = c.Logical().Read("config/" + path)

	if err != nil {
		return nil, err
	}

	for k := range res {
		if kv.Data[k] != nil {
			res[k] = kv.Data[k].(string)
		}
	}

	return res, nil
}

// WriteEncrypted write k/v with encrypted value in vault
func (c *Vault) WriteEncrypted(transitkey, path string, value []byte) (string, error) {
	var err error
	var response Response

	if c == nil {
		return "", rError.New(err, rError.Enum.INTERNALSERVERERROR, "client_has_not_been_initiated")
	}

	// Create request for vault encryption
	reqBody, _ := json.Marshal(map[string]string{
		"plaintext": base64.StdEncoding.EncodeToString(value),
	})

	uri := fmt.Sprintf("/v1/transit/encrypt/%v", transitkey)

	req := c.NewRequest("POST", uri)
	req.Body = bytes.NewBuffer(reqBody)

	res, err := c.RawRequest(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	// Pass a pointer of type Response
	err = json.Unmarshal(body, &response)

	if err != nil {
		return "", err
	}

	// Get Encrypted value from Response
	ciphertext := response.Data["ciphertext"].(string)

	data := map[string]interface{}{
		"value": ciphertext,
	}

	// Create random-UUID key and write k/v to Vault
	id, _ := uuid.NewUUID()

	_, err = c.Logical().Write(fmt.Sprintf(`%v/%v`, path, id.String()), data)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

// ReadEncrypted read k/v with encrypted value in Vault
func (c *Vault) ReadEncrypted(transitkey, path string) ([]byte, error) {
	var err error
	var resp Response

	if c == nil {
		return nil, rError.New(err, rError.Enum.INTERNALSERVERERROR, "client_has_not_been_initiated")
	}

	// Read k/v from Vault
	data, err := c.Logical().Read(path)

	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, err
	}

	// Decrypt value with given key
	reqBody, _ := json.Marshal(map[string]string{
		"ciphertext": data.Data["value"].(string),
	})

	uri := fmt.Sprintf("/v1/transit/decrypt/%v", transitkey)

	req := c.NewRequest("POST", uri)
	req.Body = bytes.NewBuffer(reqBody)

	res, err := c.RawRequest(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	// Pass a pointer of type Response
	err = json.Unmarshal(body, &resp)

	if err != nil {
		return nil, err
	}

	// Get Plaintext from Response
	plaintextByte, err := base64.StdEncoding.DecodeString(resp.Data["plaintext"].(string))

	if err != nil {
		return nil, err
	}

	return plaintextByte, err
}
