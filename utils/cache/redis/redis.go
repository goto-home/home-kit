package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type Operator struct {
	// pool redis.Pool
	cli redis.Conn
}

func NewOperator(url string) (*Operator, error) {
	conn, err := redis.Dial("tcp", url)
	if err != nil {
		return nil, err
	}
	return &Operator{cli: conn}, nil
}

func (o *Operator) Set(key, value string, date time.Duration) error {
	_, err := o.cli.Do("SET", key, value, "EX", date.Seconds())
	return err
}

func (o *Operator) Get(key string) (string, error) {
	return redis.String(o.cli.Do("GET", key))
}

func (o *Operator) Del(key string) error {
	_, err := o.cli.Do("DEL", key)
	return err
}

func (o *Operator) CloseFunc() {
	// todo judge ptr
	_ = o.cli.Close()
	// _ = o.pool.Close()
}
