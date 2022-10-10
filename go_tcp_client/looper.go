package go_tcp_client

import (
	"fmt"
	"github.com/gogf/gf/v2/net/gtcp"
	"io"
	"strings"
)

func looper(conn *gtcp.Conn) {
	OnConnected(conn)
	defer conn.Close()

	for {

		data, err := conn.Recv(-1)
		if err == nil {
			OnData(conn, data)
		}
		if err != nil {
			fmt.Println(err)

			// 自己断开
			if strings.Contains(err.Error(), "use of closed network connection") {
				OnClientDisconected(conn)
				break
			}

			// 对方断开
			if err == io.EOF || strings.Contains(err.Error(), "connection reset by peer") {
				OnServerKicked(conn)
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
