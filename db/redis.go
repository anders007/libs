package db

import (
	"fmt"
	"os"

	"gopkg.in/vmihailenco/msgpack.v2"

	log "github.com/gonet2/libs/nsq-logger"
	"github.com/fzzy/radix/redis"
)

const (
	DEFAULT_REDIS_HOST = "127.0.0.1:6379"
	ENV_REDIS_HOST     = "REDIS_HOST"
)

type db struct {
	redis_client *redis.Client
}

var Redis *db

func init() {
	Redis = &db{}
	// read redis host
	redis_host := DEFAULT_REDIS_HOST
	if env := os.Getenv(ENV_REDIS_HOST); env != "" {
		redis_host = env
	}
	// start connection to redis cluster
	client, err := redis.Dial("tcp", redis_host)
	if err != nil {
		log.Critical(err)
		os.Exit(-1)
	}
	Redis.redis_client = client

}

func (s *db) Get(tbl string, uid int32, data interface{}) error {
	raw, err := s.redis_client.Cmd("GET", Key(tbl, uid)).Bytes()
	if err != nil {
		log.Critical(err)
		return err
	}
	// unpack message from msgpack format
	err = msgpack.Unmarshal(raw, &data)
	if err != nil {
		log.Critical(err)
		return err
	}
	return nil
}

func (s *db) Set(tbl string, uid int32, data interface{}) error {
	bin, err := msgpack.Marshal(data)
	if err != nil {
		log.Critical(err)
		return err
	}
	_, err = s.redis_client.Cmd("SET", Key(tbl, uid), bin).Str()
	if err != nil {
		log.Critical(err)
		return err
	}
	return nil
}

func Key(tbl string, uid int32) string {
	if tbl == "" || uid == 0 {
		log.Critical("Key error: %+v: %+v", tbl, uid)
		return ""
	}
	return fmt.Sprintf("%s:%s", tbl, uid)
}
