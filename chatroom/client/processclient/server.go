package processclient

import (
	"fmt"
	"github.com/chatroom-han/chatroom/client/utils"
	"net"
	"os"
)

func ShowMenu() {
	for {
		fmt.Println("--------恭喜xxx登陆成功--------")
		fmt.Println("--------1.显示在线用户列表--------")
		fmt.Println("--------2.发送消息--------")
		fmt.Println("--------3.信息列表--------")
		fmt.Println("--------4.退出系统--------")
		fmt.Print("请选择（1-4）：")

		key := 0
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("显示在线用户列表")
		case 2:
			fmt.Println("发送消息")
		case 3:
			fmt.Println("信息列表")
		case 4:
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}
}

func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}

	for {
		fmt.Println("客户端正在等待服务器发送消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("服务器发送消息错误 err=", err)
			return
		}

		//读取到消息，又是下一步逻辑
		fmt.Println("mes=", mes)
	}
}
