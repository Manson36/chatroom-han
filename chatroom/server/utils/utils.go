package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/chatroom-han/chatroom/common"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf [8064]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	buf := make([]byte, 8064)
	_, err = this.Conn.Read(buf[0:4])

	//conn.read 在Conn没有关闭的情况下才会阻塞，如果客户端关闭了conn，则不会阻塞，反复报下面的错误

	if err != nil {
		fmt.Println("conn.Read error", err)
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	//根据pkgLen长度读取消息内容
	n, err := this.Conn.Read(buf[:pkgLen])
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

func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))

	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)

	n, err := this.Conn.Write(buf[:4])
	if n != 4 && err != nil {
		fmt.Println("conn.Write error",err)
		return
	}

	n, err = this.Conn.Write(data)
	if n != int(pkgLen) && err != nil {
		fmt.Println("conn.Write error",err)
		return
	}

	return
}
