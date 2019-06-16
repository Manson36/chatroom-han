package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis dial error", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("HMSet", "user2", "name", "Tom", "age", "33")
	if err != nil {
		fmt.Println("HMSet error", err)
		return
	}

	c, err := redis.Strings(conn.Do("HMGet", "user2", "name", "age"))
	if err != nil {
		fmt.Println("HMGet error", err)
		return
	}

	for i, v := range c {
		fmt.Printf("c[%d]%v\n", i, v)
	}
}
