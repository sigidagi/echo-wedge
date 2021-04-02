package client

import (
	"echo-wedge/backend/config"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"net"
)

var (
	Backend *TCPClient
)

// TCPClient struct
type TCPClient struct {
	Url  string
	Conn *net.TCPConn
}

// NewClient is constructor function for new instances of TCP client
func newClient(host string) *TCPClient {
	client := &TCPClient{
		Url: host,
	}
	return client
}

func Setup(c config.Config) error {

	log.WithFields(log.Fields{
		"url": c.Gateway.Url,
	}).Info("api: setup tcp gateway client..")

	Backend = newClient(c.Gateway.Url)
	return nil
}

// Start TCPClient
func (cl *TCPClient) connect() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", cl.Url)
	if err != nil {
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}
	cl.Conn = conn
	return nil
}

func (cl *TCPClient) Write(message []byte) error {

	if err := cl.connect(); err != nil { // error
		log.Errorf("Can not connecto to the Wedge: %s", err.Error())
		return errors.New("No connection with Wedge")
	}

	//fmt.Printf("-->> json\n %s\n", string(message))
	_, err := cl.Conn.Write(message)
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
		log.Warnf("Unmarshalling error:", err)
		return data, err
	}
	//fmt.Printf("<<-- from server:\n %v\n", string(newBuff))
	c.Conn.Close()
	return data, nil
}
