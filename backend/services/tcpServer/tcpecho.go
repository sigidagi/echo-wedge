package tcpServer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	client "github.com/io-m/echo-wedge/backend/tcpClient"
)

// Listen is method of Subscripiton object for listening to all incomming tcp client connection (such as API Gtw. incomming data)
func Listen(port int) {
	strPort := strconv.Itoa(port)
	ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%s", strPort))
	if err != nil {
		panic(err)
	}
	ch := make(chan *client.JSONRPC)
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		
		log.Println("Connection established with : ", conn.RemoteAddr())
		go func(c net.Conn) {
			
			defer c.Close()
			
			data := &client.JSONRPC{}
			buffer := make([]byte, 4096)
			bn, err := c.Read(buffer)
			if err != nil {
				log.Fatalf("Could not read from API Gateway : %s", err.Error())
			}
			if err := json.Unmarshal(buffer[:bn], &data); err != nil {
				log.Fatal(err)
			}
			fmt.Println(data)
			ch <- (data)
		}(conn)
		dataToSend := <- ch

		dispatchForward(dataToSend, 8000)
	}
}

func dispatchForward(sd *client.JSONRPC, port int) {
	bdata, _ := json.Marshal(sd.Params.Data)
	cl := &http.Client{}
	pushedData := bytes.NewBuffer(bdata)
	resp, err := cl.Post(fmt.Sprintf("http://localhost:%d/apigtw", port), "application/json",pushedData)
	if err != nil {
		resp.Body.Close()
		log.Fatal(err)
	}
	resp.Body.Close()
}