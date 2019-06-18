package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//定义一个全局的UserDao
var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

//工厂模式创建一个UserDao的实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))//三小时的挣扎，id不加引号
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}

	//将获取到的res反序列化为User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json Unmarshal err=", err)
		return
	}
	return
}

//完成登录的校验 Login
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()

	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	if userPwd != user.UserPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}