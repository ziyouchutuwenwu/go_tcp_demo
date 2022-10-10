package go_tcp_pkg_server

import (
	"github.com/gogf/gf/v2/net/gtcp"
	"io"
	"strings"
)

func looper(conn *gtcp.Conn) {
	OnConnected(conn)
	defer conn.Close()

	for {
		data, err := conn.RecvPkg(gtcp.PkgOption{HeaderSize: 2})
		if err == nil {
			OnData(conn, data)
		}
		if err != nil {
			// 自己断开
			if strings.Contains(err.Error(), "use of closed network connection") {
				OnServerKicked(conn)
				break
			}

			// 对方断开
			if err == io.EOF || strings.Contains(err.Error(), "connection reset by peer") {
				OnClientDisconected(conn)
				break
			}

			// 超时
			if strings.Contains(err.Error(), "i/o timeout") {
				OnTimeOut(conn)
				break
			}
		}
	}
}
