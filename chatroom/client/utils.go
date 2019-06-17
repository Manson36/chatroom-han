package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/chatroom-han/chatroom/common"
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
