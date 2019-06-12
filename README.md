# gokit-starter-pack
starter pack for creating gokit project

##### Table of Contents  
1. [Event Store](#event_store)  
    * [Publisher](#publisher)
    * [Subscriber](#subscriber)
2. [Vault Client](#vault_client)  
    * [Get Config](#get_config)
    * [Write Ecnrypted](#write_encrypted)
    * [Read Ecnrypted](#read_encrypted)
    
<a name="event_store"/>

## Event Store Lib
Add library event for wrapping publish and subscribe nats

<a name="publisher"/>

### Publisher
Library for publishing event to nats (begin and commit) as a middleware in transport

#### Example

```
//Create publisher
eventPublisher := event.NewPublisher("nats_connection", "logger")

//Implementation in transport or as endpoint Go-Kit
eventPublisher.Store("domain", "model", "eventtype", "topic/subject",
    func(ctx context.Context, request interface{}) (response interface{}, err error) {
        result, err := service.Create(reqData.Name, reqData.Code)
        if err != nil {
            return nil, err
        }
        return result, nil
    },
)
```

Description :

**NewPublisher**

| Param           | Description                            |
|-----------------|:---------------------------------------|
| nats_connection | nats connection type **stan.Conn**     |
| logger          | logger for logging type from gokit log |

**.Store**

| Param         | Description                            |
|---------------|:---------------------------------------|
| domain        | Your domain ex: account, authorization |
| model         | Your model from your domain            |
| eventtype     | Event Type ex: create, update          |
| topic/subject | Topic/Subject for nats                 |
| func          | Enpoint gokit                          |

<a name="subscriber"/>

### Subscriber
Library for subscribe event from nats.

#### Example

```
assessmentApproveSub := event.NewSubscriber("nats_connection", "topic/subject", "qGroup", "durable_name", "startAt", "logger", func(msg *stan.Msg) {
    var tmp map[string]interface{}
    if err := json.Unmarshal(msg.Data, &tmp); err != nil {
        logger.Log(err)
    }
    logger.Log("nats", fmt.Sprintf("Incoming message from topic/subject with data %s", tmp))
}).Subscribe()
```

Description :

**NewSubscriber**

| Param               | Description                                                                                      |
|---------------------|:-------------------------------------------------------------------------------------------------|
| nats_connection     | nats connection type **stan.Conn**                                                               |
| topic/subject       | Topic/Subject for nats                                                                           |
| qGroup              | Queue Group fill this with **your domain name**                                                  |
| durable_name        | Durable subscription ex :**authorization-sub**                                                   |
| startat             | Start at                                                                                         |
|                     | (avaliable option :                                                                              |
|                     | all,                                                                                             |
|                     | seqno ex:**sqno:100**,                                                                           |
|                     | time ex:**time:1559291755**,                                                                     |
|                     | since (click [here](https://golang.org/pkg/time/#ParseDuration) for information) ex:**since:2h** |
| logger              | logger for logging type from gokit log                                                           |
| func(msg *stan.Msg) | Handler incoming message                                                                         |



<a name="vault_client"/>

## Vault Client

<a name="get_config"/>

### Get Configuration

Library for getting configuration.
This library will get k/v from 2 paths in Vault.
1. From config/global
2. From config/{user_defined_path}, existing keys from config/global will be replaced.

#### Example

```
//Defining configuration struct
type Config struct {
	DebugAddress       string
	NatsAddress        string
}

//Defining default configuration
var defaultConfig = &Config{
	DebugAddress:       ":9080",
	NatsAddress:        "nats://localhost:4222",
}

func main() {
   error err

   //Create Vault connection
   vaultConn, err := vault.New()

   //Get Configuration from Vault as map[string]string
   cfg, err := vaultConn.GetEnvOrDefaultConfig("path", "defaultConfig")
}
```

Description :

| Param                             | Description                            |
|-----------------------------------|:---------------------------------------|
| path <string>                     | specific path in config                |
| defaultConfig <map[string]string> | default configuration if k/v not found |


<a name="write_encrypted"/>

### Write Encrypted

Library to write k/v with encrypted value.

#### Example

```
//Create Vault connection
vaultConn, err := vault.New()

//Write k/v
encryptedValue, err := vaultConn.WriteEncrypted("transitkey", "path", "key", "value")
```

Description :

| Param                             | Description                   |
|-----------------------------------|:------------------------------|
| transitkey <string>               | key to encrypt value          |
| path <string>                     | path in Vault                 |
| key <string>                      | secret key                    |
| value <string>                    | secret value                  |
   

<a name="write_encrypted"/>

### Read Encrypted

Library to read k/v with encrypted value.

#### Example

```
//Create Vault connection
vaultConn, err := vault.New()

//Write k/v
value, err := vaultConn.ReadEncrypted("transitkey", "path", "key")
```

Description :

| Param                             | Description                   |
|-----------------------------------|:------------------------------|
| transitkey <string>               | key to decrypt value          |
| path <string>                     | path in Vault                 |
| key <string>                      | secret key                    |
   
