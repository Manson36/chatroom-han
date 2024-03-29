package processclient

import (
	"encoding/json"
	"fmt"
	"github.com/chatroom-han/chatroom/common"
)

func outputGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err", err)
		return
	}

	//显示信息
	info := fmt.Sprintf("用户id: %d 对大家说：%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
}
