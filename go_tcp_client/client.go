package go_tcp_client

import (
	"github.com/gogf/gf/net/gtcp"
)

func Demo(){
	conn, _ := gtcp.NewConn("127.0.0.1:8999")
	conn.Send([]byte("我没有来"))

	looper(conn)
}