package client

import (
	"encoding/json"
	"fmt"
	"net"
)

// Interface is an interface for TCPClient server communication
type Interface interface {
	Start()
	Write([]byte) error
	ReadData(interface{}) (interface{}, error)
}

// ================================
// JSON RPC Client data

// Meta is struct for metadata
type Meta struct {
        Type    string
        Version string
        ID      string
}

// StateWg is struct for data state
type StateWg struct {
        Data interface{} `json:"data"`
        Meta Meta   `json:"meta"`
}

// Params is struct for params inside JSONRPC
type Params struct {
        URL  string      `json:"url"`
        Data interface{} `json:"data"`
}

// JSONRPC is main struct for exchanging data
type JSONRPC struct {
        ID      string `json:"id"`
        Jsonrpc string `json:"jsonrpc"`
        Method  string `json:"method"`
        Params  Params `json:"params"`
}
// =======================================

// TCPClient struct
type TCPClient struct {
        Host string
        Port int
	// JSON JSONRPC
        Conn *net.TCPConn
}

// NewClient is constructor function for new instances of TCP client
func NewClient () Interface {
	return &TCPClient{}
}



// Start TCPClient
func (c *TCPClient) Start() {
        tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port))
        if err != nil {
                panic(err)
        }

        conn, err := net.DialTCP("tcp", nil, tcpAddr)
        //defer conn.Close()
        if err != nil {
                panic(err)
        }
        c.Conn = conn
	
}

func (c *TCPClient) Write(message []byte) error {
        fmt.Printf("-->> json\n %s\n", string(message))
        _, err := c.Conn.Write(message)
        if err != nil {
                return err
        }
        return nil
}

func (c *TCPClient) ReadData(data interface{}) (interface{}, error) {
        reply := make([]byte, 4096)
        nb, err := c.Conn.Read(reply)
        buff := reply[:nb]
        if err != nil {
            return data, err
        }   
        if err = json.Unmarshal(buff, &data); err != nil {
            fmt.Println("Unmarshalling error:", err)
            return data, err
        }   
        fmt.Printf("<<-- from server data:\t %v\n", string(buff))
        return data, nil
} 