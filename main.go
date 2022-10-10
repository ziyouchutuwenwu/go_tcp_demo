package main

import (
	"go_tcp_demo/go_tcp_pkg_client"
)

func main() {
	//go go_tcp_pkg_server.Start()
	//time.Sleep(time.Second * 2)

	go_tcp_pkg_client.Demo()
	//go_tcp_server.Start()
}
