package go_tcp_pkg_client

import (
	"github.com/gogf/gf/v2/encoding/gbinary"
	"github.com/gogf/gf/v2/net/gtcp"
)

func Demo() {
	conn, _ := gtcp.NewConn("127.0.0.1:9999")

	data := make([]byte, 0)
	cmd := gbinary.BeEncodeUint16(111)
	data = append(data, cmd...)
	data = append(data, []byte("这是服务器发的测试数据")...)
	conn.SendPkg(data, gtcp.PkgOption{HeaderSize: 2})

	looper(conn)
}
