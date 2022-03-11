package cache

import "time"

type ICache interface {
	Set(key, value string, date time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
}
