package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool {
		MaxIdle: 8,
		MaxActive:0,
		IdleTimeout: 100,
		Dial:func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "套母猫")
	if err != nil {
		fmt.Println("Set error", err)
		return
	}

	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("Get error", err)
		return
	}

	fmt.Println("get ", r)
}
