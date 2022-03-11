package redis

import (
	"github.com/gomodule/redigo/redis"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	// new
	operator, err := NewOperator("localhost:6379")
	if err != nil {
		t.Fail()
		return
	}
	// set 5*seconds
	if err := operator.Set("demo", "one", 5*time.Second); err != nil {
		t.Fail()
		return
	}
	// get
	str, err := operator.Get("demo")
	if err != nil {
		t.Fail()
		return
	}
	if str != "one" {
		t.Fail()
		return
	}
	time.Sleep(6*time.Second)
	str, err = operator.Get("demo")
	switch err {
	case nil:
		t.Fail()
		return
	case redis.ErrNil:
	default:
		t.Fail()
		return
	}
	if str == "one" {
		t.Fail()
		return
	}
}
