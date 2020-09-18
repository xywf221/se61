package socket

import "github.com/xywf221/se61"

type ConnectHeader struct {
	Uid   uint32
	Token string
}

func NewConnectHeader() *ConnectHeader {
	return &ConnectHeader{}
}

func (c ConnectHeader) WriteTo(b *se61.ProtocolBytes) {
	b.WriteUInt32(c.Uid)
	b.WriteString(c.Token)
}
