package main

import (
	"log"
	"net"
	// "time"
)

func main() {
	ln, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server is listening on port 8001...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Connection from client is accepted")
		for {
			go handleConnection(conn)
		}
	}	
}
func handleConnection(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("Error while closing connection: %v", err)
		}
	}()
	for {
		b := make([]byte, (1024 * 4))
		nb, err := conn.Read(b)
		if nb == 0 || err != nil {
			log.Println("Connection red error:", err)
			break
		}
		log.Println(string(b))
		_, err = conn.Write([]byte("Json-RPC <--- Yes, I got your message.."))
		if err != nil {
			log.Println("Could not write back error:", err)
			break
		}
		// time.Sleep(time.Second * 4)		
	}
	log.Println("Connection is closed")
}