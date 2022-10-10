package go_tcp_pkg_server

import (
	"github.com/gogf/gf/v2/net/gtcp"
	"time"
)

func SetRecvDeadline(conn *gtcp.Conn, second time.Duration) {
	conn.SetReadDeadline(time.Now().Add(time.Second * second))
}

func SetSendDeadline(conn *gtcp.Conn, second time.Duration) {
	conn.SetWriteDeadline(time.Now().Add(time.Second * second))
}

func SetDeadline(conn *gtcp.Conn, second time.Duration) {
	conn.SetDeadline(time.Now().Add(time.Second * second))
}
