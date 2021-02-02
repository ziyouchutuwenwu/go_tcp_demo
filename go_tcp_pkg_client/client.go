package go_tcp_pkg_client

import (
	"github.com/gogf/gf/net/gtcp"
)

func Demo(){
	conn, _ := gtcp.NewConn("127.0.0.1:8999")
	conn.SendPkg([]byte("我没有来"), gtcp.PkgOption{HeaderSize: 2})

	looper(conn)
}