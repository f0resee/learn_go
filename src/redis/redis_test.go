package redis

import (
	"github.com/go-playground/assert/v2"
	"github.com/gomodule/redigo/redis"
	"testing"
)

func TestRedisConn(t *testing.T) {
	t.Run("redis connect", func(t *testing.T) {
		c, err := redis.Dial("tcp", ":6379")
		if err != nil {
			t.Fail()
			t.Logf("redis.Dial() error: %v", err)
			return
		}
		c.Close()
	})

	t.Run("redis set and get", func(t *testing.T) {
		c, err := redis.Dial("tcp", ":6379")
		if err != nil {
			t.Fail()
			t.Logf("redis.Dial() error: %v", err)
			return
		}

		defer c.Close()

		_, err = c.Do("SET", "key1", 998)
		if err != nil {
			t.Fail()
			t.Logf("set error: %v", err)
			return
		}

		r, err := redis.Int(c.Do("GET", "key1"))
		if err != nil {
			t.Logf("set error: %v", err)
		}
		assert.Equal(t, r, 998)

	})

}
