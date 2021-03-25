package main

import "github.com/io-m/echo-wedge/backend/services/tcpServer"

func main() {
	tcpServer.Listen(8060)
}