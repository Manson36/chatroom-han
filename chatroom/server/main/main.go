package main

import (
	"fmt"
	"github.com/chatroom-han/chatroom/server/model"
	"net"
	"time"
)

func process(conn net.Conn) {
	//读取客户发送的消息
	defer conn.Close()

	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端与服务端的通信出问题了", err)
		return
	}
}

func init() {
	//当服务器启动时，初始化一个redis的链接池
	initpool("127.0.0.1:6379", 16, 0, 300 *time.Second)
	initUserDao()
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
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
