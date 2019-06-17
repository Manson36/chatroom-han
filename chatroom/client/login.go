package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/chatroom-han/chatroom/common"
	"net"
)


func login(userId int, userPwd string) (err error) {
	//fmt.Printf("你输入的Id：%v, 密码：%v\n", userId, userPwd)
	//return nil

	//1 连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial error=", err)
		return
	}

	defer conn.Close()

	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data, err := json.Marshal(loginMes) //data 是byte类型
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) error =", err)
		return
	}

	//发送data的长度
	var pkgLen uint32
	pkgLen = uint32(len(data)) //len()返回的是int类型

	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen) //将Uint32转换为byte类型

	n, err := conn.Write(buf[0:4])
	if err !=nil || n != 4 {
		fmt.Println("conn.Write err", err)
		return
	}

	//fmt.Printf("客户端消息长度发送完成 长度：%v， 内容%v\n", len(data), string(data))

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) error", err)
		return
	}

	//time.Sleep(20* time.Second)
	//fmt.Println("客户端休眠了20s")
	//这里需要处理服务器返回的消息
	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg err =", err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	if loginResMes.Code ==200 {
		fmt.Println("登陆成功")
	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}