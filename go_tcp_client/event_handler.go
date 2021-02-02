package go_tcp_client

import (
	"fmt"
	"github.com/gogf/gf/net/gtcp"
)

func OnConnected(conn *gtcp.Conn){
	//SetRecvDeadline(conn, 2)

	fmt.Println("OnConnected")
}

func OnData(conn *gtcp.Conn, data []byte){
	fmt.Println("OnData ", data)
}

func OnServerKicked(conn *gtcp.Conn){
	fmt.Println("OnServerKicked")
}

func OnClientDisconected(conn *gtcp.Conn){
	fmt.Println("OnClientDisconected")
}

func OnTimeOut(conn *gtcp.Conn){
	fmt.Println("OnTimeOut")
}

