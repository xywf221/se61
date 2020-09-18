package main

import (
	"fmt"
	"github.com/xywf221/se61"
	"github.com/xywf221/se61/socket"
	"time"
)

func main() {
	type AutoGenerated struct {
		Action string `json:"action"`
		Params struct {
			Game     int    `json:"game"`
			Vericode string `json:"vericode"`
			Account  string `json:"account"`
			Passwd   string `json:"passwd"`
			Operator int    `json:"operator"`
			Entrance int    `json:"entrance"`
			Channel  string `json:"channel"`
		} `json:"params"`
	}

	s := AutoGenerated{
		Action: "auth",
	}
	s.Params.Game = 657
	s.Params.Vericode = ""
	s.Params.Account = "993651481@qq.com"
	s.Params.Passwd = "0c785ab0bc768016c022d69d3a9834c9"
	s.Params.Operator = 1
	s.Params.Entrance = 1
	s.Params.Channel = "other"
	msg := se61.NewCmdMessage(338, 1, s)
	sz := socket.NewAutoConnectSocket()
	//sz.SendMessages([]*se61.CmdMessage{msg}, true)
	err := sz.Connect("ws://115.159.142.36:21200")
	if err != nil {
		fmt.Println("连接服务器出现错误")
	} else {
		fmt.Println("连接服务器成功!")
	}
	err = sz.SendMessages([]*se61.CmdMessage{msg}, true)
	if err != nil {
		panic(err)
	}
	fmt.Println("消息发送完毕")
	time.Sleep(time.Second * 20)
}
