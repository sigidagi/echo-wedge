package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

var (
	Backend *TCPClient
)

// TCPClient struct
type TCPClient struct {
	Host string
	Port int
	Conn *net.TCPConn
}

// NewClient is constructor function for new instances of TCP client
func newClient(host string, port int) *TCPClient {
	client := &TCPClient{
		Host: host,
		Port: port,
	}
	return client
}

func Setup() error {

	port := 8051
	fmt.Printf("Connecting to Wedge on %d port", port)

	Backend = newClient("localhost", port)
	return nil
}

// Start TCPClient
func (c *TCPClient) connect() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}
	c.Conn = conn
	return nil
}

func (c *TCPClient) Write(message []byte) error {

	if err := c.connect(); err != nil { // error
		fmt.Printf("Can not connecto to the Wedge: %s", err.Error())
		return errors.New("No connection with Wedge")
	}

	fmt.Printf("-->> json\n %s\n", string(message))
	_, err := c.Conn.Write(message)
	if err != nil {
		return err
	}
	return nil
}

// ReadNetwork is method for sending Network details to frontend client
func (c *TCPClient) ReadData(data interface{}) (interface{}, error) {
	reply := make([]byte, 4096)
	nb, err := c.Conn.Read(reply)
	newBuff := make([]byte, nb)
	newBuff = reply[:nb]
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(newBuff, &data); err != nil {
		fmt.Println("Unmarshalling error:", err)
		return data, err
	}
	fmt.Printf("<<-- from server:\n %v\n", string(newBuff))
	c.Conn.Close()
	return data, nil
}
