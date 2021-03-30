package main

import (
	er "github.com/io-m/echo-wedge/backend/router/echo"
	serv "github.com/io-m/echo-wedge/backend/services/tcpServer"
	cl "github.com/io-m/echo-wedge/backend/tcpClient"
)

func main() {
	if err := cl.Setup(); err != nil {
		panic(err)
	}

	if err := serv.Setup(); err != nil {
		panic(err)
	}

	er.RunApp()
}
