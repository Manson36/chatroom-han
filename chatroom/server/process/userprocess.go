package process

import (
	"encoding/json"
	"fmt"
	"github.com/chatroom-han/chatroom/common"
	"github.com/chatroom-han/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}