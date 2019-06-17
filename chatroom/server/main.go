package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/chatroom-han/chatroom/common"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8064)
	_, err = conn.Read(buf[0:4])

	//conn.read 在Conn没有关闭的情况下才会阻塞，如果客户端关闭了conn，则不会阻塞，反复报下面的错误

	if err != nil {
		fmt.Println("conn.Read error", err)
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	//根据pkgLen长度读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read error", err)
		return
	}

	 err = json.Unmarshal(buf[:pkgLen], &mes) //技术就是一层窗户纸，&mes
	 if err != nil {
		 fmt.Println("json.Unmarshal error", err)
		 return
	 }

	 return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))

	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)

	n, err := conn.Write(buf[:4])
	if n != 4 && err != nil {
		fmt.Println("conn.Write error",err)
		return
	}

	n, err = conn.Write(data)
	if n != int(pkgLen) && err != nil {
		fmt.Println("conn.Write error",err)
		return
	}

	return
}

func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}

	//声明返回信息
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes	message.LoginResMes

	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在"
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal error", err)
		return
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal error", err)
		return
	}

	err = writePkg(conn, data)
	return
}

func process(conn net.Conn) {
	//读取客户发送的消息
	defer conn.Close()

	for {
		mes, err := readPkg(conn)
		if err != nil {
			//解决循环报错
			if err == io.EOF {
				fmt.Println("客户端退出了，我也退出。。。")
				return
			}else {
				fmt.Println("readPkg error", err)
				return
			}
		}

		//fmt.Println("mes=", mes)
		err = serverProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}

}

//编写一个ServerProcessMes函数，根据客户发送的信息类型，调用不同的函数处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
		case message.LoginMesType:
		//处理登陆
		err = serverProcessLogin(conn, mes)
		default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

func main() {
	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	defer listen.Close()

	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}

	for {
		fmt.Println("等待客户端来连接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
			return
		}

		//一旦连接成功，启动一个协程与客户端保持通信
		go process(conn)
	}

}
