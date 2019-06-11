package event

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-redis/redis"
	stan "github.com/nats-io/stan.go"
)

// Locker use for locking data using redis
type Locker struct {
	DB       *redis.Client
	NatsConn stan.Conn
	logger   log.Logger
}

//NewLocker create new Connection to RedisDB
func NewLocker(addr, password string, db int, conn stan.Conn, logger log.Logger) (*Locker, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		logger.Log("redis", fmt.Sprintf("err %s", err))
		return nil, err
	}
	return &Locker{
		DB:       client,
		logger:   logger,
		NatsConn: conn,
	}, nil
}

//LockData for locking data
func (locker *Locker) LockData(domain, functionName, id string, data map[string]interface{}, subEvent string, handler stan.MsgHandler) {
	keyDomain := fmt.Sprintf("%s-%s-%s", domain, functionName, id)
	existQueue, err := locker.getQueue(keyDomain)
	if err != nil {
		locker.logger.Log("err", err)
	}
	if len(existQueue) == 0 {

	} else {

	}
	//onprogress
}

func (locker *Locker) getQueue(keyStore string) ([]map[string]interface{}, error) {
	var existLock []map[string]interface{}
	getLock, err := locker.DB.Get(keyStore).Result()
	if err != nil && err != redis.Nil {
		existLock = make([]map[string]interface{}, 0)
		return existLock, err
	}
	if len(getLock) > 0 {
		err = json.Unmarshal([]byte(getLock), &existLock)
		if err != nil {
			return existLock, nil
		}
	} else {
		existLock = make([]map[string]interface{}, 0)
	}
	return existLock, nil
}

//insertLock for insert locked data to redis
func (locker *Locker) insertLock(domain, functionName, id string, data map[string]interface{}) error {
	newLock := make(map[string]interface{})
	newLock["timestamp"] = time.Now().Unix()
	newLock["data"] = data
	newLock["status"] = "locked"
	keyDomain := fmt.Sprintf("%s-%s-%s", domain, functionName, id)

	existLock, err := locker.getQueue(keyDomain)
	if err != nil {
		return err
	}
	if len(existLock) > 0 {
		newLock["status"] = "pending"
	}

	existLock = append(existLock, newLock)
	byteArray, err := json.Marshal(existLock)
	if err != nil {
		return err
	}

	err = locker.DB.Set(keyDomain, byteArray, 0).Err()
	if err != nil {
		return err
	}

	locker.logger.Log("redis", fmt.Sprintf("Queue Locked Data %s : %s", keyDomain, existLock))
	return nil
}
