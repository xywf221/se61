package se61

import (
	"bufio"
	"golang.org/x/net/websocket"
)

type Conn struct {
	conn *websocket.Conn
	brw  *bufio.ReadWriter
}

func NewConn() *Conn {
	return &Conn{conn: &websocket.Conn{}}
}

func (c *Conn) Connect(server string) error {
	config, err := websocket.NewConfig(server, "http://s.61.com")
	if err != nil {
		return err
	}

	config.Header.Set("User-Agent", "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36 Edg/85.0.564.51")
	config.Header.Set("Origin", "http://s.61.com")
	config.Header.Set("Accept-Encoding", "gzip, deflate")
	config.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")

	wc, err := websocket.DialConfig(config)
	if err != nil {
		return err
	}
	//todo 修改为其他的读取方式
	br, bw := bufio.NewReader(wc), bufio.NewWriter(wc)
	c.conn = wc
	c.brw = bufio.NewReadWriter(br, bw)
	return nil
}

func (c *Conn) Write(p []byte) (n int, err error) {
	//todo 修改为一次性吧
	w, err := c.conn.NewFrameWriter(websocket.BinaryFrame)
	if err != nil {
		return 0, err
	}

	return w.Write(p)
}

func (c *Conn) Read(p []byte) (n int, err error) {
	return c.brw.Read(p)
}

func (c *Conn) Flush() error {
	return c.brw.Flush()
}

func (c *Conn) Connected() bool {
	//todo get conn status
	return true
}
