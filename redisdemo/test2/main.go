package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis dial err", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("HSet", "user1", "name", "john")
	if err != nil {
		fmt.Println("HSet name error", err)
		return
	}

	_, err = conn.Do("HSet", "user1", "age", "23")
	if err != nil {
		fmt.Println("HSet age error", err)
		return
	}

	r1, err := redis.String(conn.Do("HGet", "user1", "name"))
	if err != nil {
		fmt.Println("HGet name error", err)
		return
	}

	r2, err := redis.Int(conn.Do("HGet", "user1", "age"))
	if err != nil {
		fmt.Println("HGet name error", err)
		return
	}

	fmt.Printf("操作OK，r1=%v, r2=%v", r1, r2)
}
