package main

import (
	"fmt"
	"github.com/chatroom-han/chatroom/client/processclient"
	"os"
)

var userId int
var userPwd string

func main() {
	var key int
	//var loop = true

	for {
		fmt.Println("---------欢迎登陆多人聊天系统--------")
		fmt.Println("          1.登陆聊天室")
		fmt.Println("          2.注册用户")
		fmt.Println("          3.退出系统")
		fmt.Print("请选择（1-3）：")

		fmt.Scanln(&key)

		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户Id")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户密码")
			fmt.Scanln(&userPwd)

			//完成登录
			up := &processclient.UserProcess{}
			up.Login(userId, userPwd)

		case 2:
			fmt.Println("注册用户")
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}
}
