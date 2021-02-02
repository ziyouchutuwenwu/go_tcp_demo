package go_tcp_server

import (
	"github.com/gogf/gf/net/gtcp"
)

func Send(conn *gtcp.Conn, data []byte){
	SetSendDeadline(conn, 1)
	conn.Send(data)
}

func Start() {
	server := gtcp.NewServer("", nil)
	server.SetAddress(":8999")
	server.SetHandler(looper)
	server.Run()
}