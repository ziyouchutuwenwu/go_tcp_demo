package go_tcp_pkg_server

import (
	"github.com/gogf/gf/v2/net/gtcp"
)

func Send(conn *gtcp.Conn, data []byte) {
	SetSendDeadline(conn, 1)
	conn.SendPkg(data, gtcp.PkgOption{HeaderSize: 2})
}

func Start() {
	server := gtcp.NewServer("", nil)
	server.SetAddress(":8999")
	server.SetHandler(looper)
	server.Run()
}
