package main

import (
	"fmt"
	"github.com/chatroom-han/chatroom/common"
	process3 "github.com/chatroom-han/chatroom/server/process"
	"github.com/chatroom-han/chatroom/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes函数，根据客户发送的信息类型，调用不同的函数处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登陆
		//创建一个userProcess实例
		up := &process3.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

func (this *Processor) process2() (err error) {
	for {
		tf := utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			//解决循环报错
			if err == io.EOF {
				fmt.Println("客户端退出了，我也退出。。。")
				return err
			}else {
				fmt.Println("readPkg error", err)
				return err
			}
		}

		//fmt.Println("mes=", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}