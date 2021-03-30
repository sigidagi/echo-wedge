package tcpServer

import (
	"bytes"
	"encoding/json"
	"fmt"
	m "github.com/io-m/echo-wedge/backend/models"
	"log"
	"net"
	"net/http"
)

var (
	tcpserver *server
)

type server struct {
	Port int
	Host string
}

func New(host string, port int) *server {
	return &server{
		Host: host,
		Port: port,
	}
}

func Setup() error {

	tcpserver = New("127.0.0.1", 8060)
	go tcpserver.serve()
	return nil
}

// Listen is method of Subscripiton object for listening to all incomming tcp client connection (such as API Gtw. incomming data)
func (s *server) serve() {
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		log.Fatalf("Error starting TCP server: %s", err.Error())
	}
	defer ln.Close()

	ch := make(chan *m.JSONRPC)
	for {
		conn, _ := ln.Accept()
		go func(c net.Conn) {
			defer c.Close()

			data := &m.JSONRPC{}
			buffer := make([]byte, 4096)
			bn, err := c.Read(buffer)
			if err != nil {
				log.Fatalf("Could not read from API Gateway : %s", err.Error())
			}
			if err := json.Unmarshal(buffer[:bn], &data); err != nil {
				log.Fatal(err)
			}
			ch <- (data)
		}(conn)
		dataToSend := <-ch

		dispatchForward(dataToSend, 8000)
	}
}

func dispatchForward(sd *m.JSONRPC, port int) {
	bdata, _ := json.Marshal(sd.Params.Data)
	//fmt.Printf("Push state update: %s\n", string(bdata))
	cl := &http.Client{}
	pushedData := bytes.NewBuffer(bdata)
	resp, err := cl.Post(fmt.Sprintf("http://localhost:%d/apigtw", port), "application/json", pushedData)
	if err != nil {
		resp.Body.Close()
		log.Fatal(err)
	}
	resp.Body.Close()
}
