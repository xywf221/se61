package se61

import (
	"encoding/json"
)

type CmdMessage struct {
	body   interface{} //json
	header *CmdHeader
	cmd    uint32
}

func NewCmdMessage(cmd uint32, seq uint32, body interface{}) *CmdMessage {
	m := &CmdMessage{cmd: cmd, body: body}
	h := &CmdHeader{}
	m.header = h
	h.Seq = seq
	return m
}

func (c *CmdMessage) WriteTo(b *ProtocolBytes) error {
	bodyBytes := NewProtocolBytes()
	j, err := json.Marshal(&c.body)
	if err != nil {
		return err
	}
	bodyBytes.WriteUInt32(uint32(len(j)))
	bodyBytes.WriteBytes(j) //写入body
	h := c.header
	headerBytes := NewProtocolBytes()

	h.Len = uint32(bodyBytes.Length() + 21)
	h.Cmd = c.cmd
	h.Checksum = 0
	h.Zip = false
	h.Write(headerBytes)                      //写入header
	headerBytes.WriteBytes(bodyBytes.Bytes()) //写入bodyBytes
	b.WriteBytes(headerBytes.Bytes())
	return nil
}

func (c *CmdMessage) GetCmd() uint32 {
	if c.header.Cmd == 0 {
		return c.cmd
	} else {
		return c.header.Cmd
	}

}

func (c *CmdMessage) GetSeq() uint32 {
	return c.header.Seq
}
