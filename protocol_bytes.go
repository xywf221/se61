package se61

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type ProtocolBytes struct {
	position int //可能为负数
	buf      *bytes.Buffer
}

func NewProtocolBytes() *ProtocolBytes {
	return &ProtocolBytes{buf: new(bytes.Buffer), position: 0}
}

func (p *ProtocolBytes) WriteUInt32(u uint32) {
	binary.Write(p.buf, binary.LittleEndian, u)
}

func (p *ProtocolBytes) WriteInt32(i int32) {
	binary.Write(p.buf, binary.LittleEndian, i)
}

func (p *ProtocolBytes) WriteBoole(b bool) {
	binary.Write(p.buf, binary.LittleEndian, b)
}

func (p *ProtocolBytes) ReadUInt32() uint32 {
	var u uint32
	binary.Read(p.buf, binary.LittleEndian, &u)
	return u
}

func (p *ProtocolBytes) ReadInt32() int32 {
	var i int32
	binary.Read(p.buf, binary.LittleEndian, &i)
	return i
}

func (p *ProtocolBytes) ReadBoolean() bool {
	var b bool
	binary.Read(p.buf, binary.LittleEndian, b)
	return b
}

func (p *ProtocolBytes) WriteString(s string) {
	p.WriteUInt32(uint32(len(s)))
	p.buf.WriteString(s) //也许后面有些空格...
}

func (p *ProtocolBytes) WriteBytes(b []byte) {
	p.buf.Write(b)
}

func (p *ProtocolBytes) Bytes() []byte {
	return p.buf.Bytes()
}
func (p *ProtocolBytes) Length() int {
	return p.buf.Len()
}
func (p *ProtocolBytes) Test() {
	fmt.Println(p.buf.Bytes())
	p.buf.Grow(0)
	p.WriteUInt32(12)
	fmt.Println(p.buf.Bytes())
}

//func (p *ProtocolBytes) writeBytes(buf []byte) {
//	for i := 0; i < 4; i++ {
//		if len(p.buf) < p.position+i+1 {
//			p.buf = append(p.buf, buf[i])
//		} else {
//			p.buf[p.position+i] = buf[i]
//		}
//	}
//}
