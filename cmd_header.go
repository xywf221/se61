package se61

type CmdHeader struct {
	Len      uint32
	Seq      uint32
	Cmd      uint32
	Ret      int32
	Checksum uint32
	Zip      bool
}

func (h CmdHeader) Write(b *ProtocolBytes) {
	b.WriteUInt32(h.Len)
	b.WriteUInt32(h.Seq)
	b.WriteUInt32(h.Cmd)
	b.WriteInt32(h.Ret)
	b.WriteUInt32(h.Checksum)
	b.WriteBoole(h.Zip)
}

func (h *CmdHeader) Read(b *ProtocolBytes) {
	h.Len = b.ReadUInt32()
	h.Seq = b.ReadUInt32()
	h.Cmd = b.ReadUInt32()
	h.Ret = b.ReadInt32()
	h.Checksum = b.ReadUInt32()
	h.Zip = b.ReadBoolean()
}

func (h CmdHeader) String() string {
	return ""
}
