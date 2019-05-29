package vault

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"

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

// Init ...
func (v *Vault) Init(addr string, token string) {
	v.Client, _ = api.NewClient(nil)
	v.SetAddress(addr)
	v.SetToken(token)
}

func (v *Vault) Write(path string, data map[string]interface{}) error {
	var err error

	if v.Client == nil {
		return rError.New(err, rError.Enum.INTERNALSERVERERROR, "client_has_not_been_initiated")
	}

	_, err = v.Logical().Write(path, data)

	if err != nil {
		return err
	}

	return nil
}

func (v *Vault) Read(path string) (map[string]interface{}, error) {
	var err error

	if v.Client == nil {
		return nil, rError.New(err, rError.Enum.INTERNALSERVERERROR, "client_has_not_been_initiated")
	}

	kv, err := v.Logical().Read(path)

	if err != nil {
		return nil, err
	}

	return kv.Data, err
}

// List ...
func (v *Vault) List(path string) (map[string]interface{}, error) {
	var err error

	if v.Client == nil {
		return nil, rError.New(err, rError.Enum.INTERNALSERVERERROR, "client_has_not_been_initiated")
	}

	kv, err := v.Logical().List(path)

	if err != nil {
		return nil, err
	}

	return kv.Data, err
}

// Encrypt ...
func (v *Vault) Encrypt(key, data string) (string, error) {
	var err error
	var resp Response

	if v.Client == nil {
		return "", rError.New(err, rError.Enum.INTERNALSERVERERROR, "client_has_not_been_initiated")
	}

	reqBody, _ := json.Marshal(map[string]string{
		"plaintext": base64.StdEncoding.EncodeToString([]byte(data)),
	})

	path := fmt.Sprintf("/v1/transit/encrypt/%v", key)

	req := v.NewRequest("POST", path)
	req.Body = bytes.NewBuffer(reqBody)

	res, err := v.RawRequest(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	// Pass a pointer of type Response and Go'll do the rest
	err = json.Unmarshal(body, &resp)

	if err != nil {
		return "", err
	}

	// Get Ciphertext from Response
	ct := resp.Data["ciphertext"].(string)

	return ct, err
}

// Decrypt ...
func (v *Vault) Decrypt(key, data string) (string, error) {
	var err error
	var resp Response

	if v.Client == nil {
		return "", rError.New(err, rError.Enum.INTERNALSERVERERROR, "client_has_not_been_initiated")
	}

	reqBody, _ := json.Marshal(map[string]string{
		"ciphertext": data,
	})

	path := fmt.Sprintf("/v1/transit/decrypt/%v", key)

	req := v.NewRequest("POST", path)
	req.Body = bytes.NewBuffer(reqBody)

	res, err := v.RawRequest(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	// Pass a pointer of type Response and Go'll do the rest
	err = json.Unmarshal(body, &resp)

	if err != nil {
		return "", err
	}

	// Get Plaintext from Response
	ptb, err := base64.StdEncoding.DecodeString(resp.Data["plaintext"].(string))

	pt := string(ptb)

	if err != nil {
		return "", err
	}

	return pt, err
}
