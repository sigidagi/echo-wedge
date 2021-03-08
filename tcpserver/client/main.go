package main

import(
	"log"
	// "fmt"
	"net"
	"time"
)

func main() {
	// port: 8001
	// host := "localhost:"
	conn, err := net.Dial("tcp", "localhost:8001")
	if err != nil {
		log.Fatal("Cannot connect to server")
	}
	log.Println("Connected to tcp server")
	resChan := make(chan string)
	go func(c chan string){
		defer conn.Close()
		msg := "hi this is message from tcp client"
		if _, err := conn.Write([]byte(msg)); err != nil {
			log.Fatal(err)
		}

		b := make([]byte, (1024 * 4))
		rb, err := conn.Read(b)
		if err != nil || rb == 0 {
			log.Fatal(err)
		}
		c <- string(b)
		time.Sleep(time.Second * 5)
	}(resChan)
		
	result := <- resChan
	log.Println(result)

}