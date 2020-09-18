package socket

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/xywf221/se61"
	"golang.org/x/net/websocket"
)

type AutoConnectSocket struct {
	conn  *se61.Conn
	Uid   uint32
	Token string
	seq   uint32
}

func NewAutoConnectSocket() *AutoConnectSocket {
	return &AutoConnectSocket{conn: se61.NewConn()}
}

func (a *AutoConnectSocket) Connect(u string) error {
	err := a.conn.Connect(u)
	if err != nil {
		return err
	}
	go func() {
		buf := make([]byte, 1024)
		for {
			n, _ := a.conn.Read(buf)
			fmt.Println(string(buf[:n]))
		}
		websocket.JSON.Receive()
	}()
	return nil
}
func (a *AutoConnectSocket) Send(cmd uint32, body interface{}) {
	if a.seq == 255 {
		a.seq = 0
	}
	msg := se61.NewCmdMessage(cmd, a.seq, body)
	a.SendMessage(msg)
}

func (a *AutoConnectSocket) SendMessage(msg *se61.CmdMessage) {
	//todo 如果没连接可以把消息推到一个队列中 然后重新连接
	if a.conn.Connected() {
		a.sendMessage(msg)
	} else {
		//todo 重新连接
	}
}

func (a *AutoConnectSocket) sendMessage(msg *se61.CmdMessage) {

}

//socket连接完毕进入
func (a *AutoConnectSocket) SocketConnected() {
	//判断未发送列表中是否又数据。有的话发送
}

//发送多个消息
func (a *AutoConnectSocket) SendMessages(messages []*se61.CmdMessage, b bool) error {
	t := se61.NewProtocolBytes()
	t.WriteUInt32(0)
	if b {
		t.WriteBoole(true)
		connHeader := NewConnectHeader()
		connHeader.Uid = a.Uid
		connHeader.Token = a.Token
		connHeader.WriteTo(t)
	} else {
		t.WriteBoole(false)
	}
	t.WriteUInt32(uint32(len(messages)))
	for _, msg := range messages {
		if err := msg.WriteTo(t); err != nil {
			return err
		}
	}
	//这步只是为了把len写在buf前
	buf := t.Bytes()
	//lenBytes := buf[len(buf)-4:]
	//buf = buf[0 : len(buf)-4]
	//buf = append(lenBytes, buf...)
	//fmt.Println(buf)
	lenBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(lenBytes, uint32(t.Length()))

	buf = bytes.Join([][]byte{lenBytes, buf[4:]}, nil)
	if _, err := a.conn.Write(buf); err != nil {
		return err
	}
	if err := a.conn.Flush(); err != nil {
		return err
	}
	return nil
}
