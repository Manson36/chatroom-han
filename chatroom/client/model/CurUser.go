package model

import (
	"github.com/chatroom-han/chatroom/common"
	"net"
)
//在客户端要多次使用CurUser，所以创建一个全局的，放在了userMgr

type CurUser struct {
	Conn net.Conn
	message.User
}
