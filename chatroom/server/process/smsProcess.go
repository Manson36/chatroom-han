package process

import (
	"encoding/json"
	"fmt"
	"github.com/chatroom-han/chatroom/common"
	"github.com/chatroom-han/chatroom/server/utils"
	"net"
)

type SmsProcess struct {

}


func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	//取出mes的内容 SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err", err)
		return
	}

	//将mes序列化
	data, err := json.Marshal(mes) //json.Marshal()填写的必须是一个结构体
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}

	//遍历服务器的onlineUsers，将消息转发给每个在线的用户
	for id, up := range userMgr.onlineUsers {
		//这里，发送的消息，就不要发送给自己了
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	//创建一个tf实例，发送data
	tf := &utils.Transfer{
		Conn: conn,
	}

	err :=  tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败， err=", err)
	}
}