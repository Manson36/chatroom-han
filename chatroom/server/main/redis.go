package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

//定义一个全局的pool
var pool *redis.Pool

func initpool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:maxIdle,
		MaxActive: maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn,error) {
			return redis.Dial("tcp", address)
		},
	}
}

