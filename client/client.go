package client

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/io-m/echo-wedge/models"
)

// Interface is an interface for TCPClient server communication
type Interface interface {
	Start()
	Write([]byte) error
	ReadNetwork() *JSONRPCResponseNetwork
	Read() *JSONRPCResponse
}

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


// ValueResponse is struct embedded in value response
type ValueResponseNet struct {
        Name string `json:"name"`
        Device []models.Device `json:"device"`
}
// Result is struct from response
type ResultNet struct {
        Value ValueResponseNet `json:"value"`
}



// ===================
type ResultValue struct {
	Child []struct {
		Type    string `json:"type"`
		Version string `json:"version"`
	} `json:"child"`
	Id    []string `json:"id"`
	More  bool     `json:"mode"`
	Count int      `json:"count"`
	Meta  struct {
		Type    string `json:"type"`
		Version string `json:"version"`
	} `json:"meta"`
}


type Result struct {
	Value ResultValue `json:"value"`
	Meta  struct {
		ServerSendTime string `json:"server_send_time"`
	} `json:"meta"`
}

// ===================

// ErrorResponse is struct embedded insinde JSON response
type ErrorResponse struct {
        Code int32 `json:"code"`
        Message string `json:"message"`
}
// JSONRPCResponse is struct for Sending API gtw's response
type JSONRPCResponse struct {
        ID      string `json:"id"`
        Jsonrpc string `json:"jsonrpc"`
        Result Result  `json:"result"`
        Error ErrorResponse `json:"error"`
}
// JSONRPCResponseNetwork is struct for Sending API gtw's response
type JSONRPCResponseNetwork struct {
        ID      string `json:"id"`
        Jsonrpc string `json:"jsonrpc"`
        Result ResultNet  `json:"result"`
        Error ErrorResponse `json:"error"`
}




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
func (c *TCPClient) ReadNetwork() *JSONRPCResponseNetwork {
        data := new(JSONRPCResponseNetwork)
        reply := make([]byte, 4096)
        nb, err := c.Conn.Read(reply)
        newBuff := make([]byte, nb)
        newBuff = reply[:nb]
        if err != nil {
                panic(err)
        }
        if err = json.Unmarshal(newBuff, &data); err != nil {
                fmt.Println("Unmarshalling error:",err)
                return nil
        }
        fmt.Printf("<<-- from server:\n %v\n", string(newBuff))
        return data
}
func (c *TCPClient) Read() *JSONRPCResponse {
        data := new(JSONRPCResponse)
        reply := make([]byte, 4096)
        nb, err := c.Conn.Read(reply)
        newBuff := make([]byte, nb)
        newBuff = reply[:nb]
        if err != nil {
                panic(err)
        }
        if err = json.Unmarshal(newBuff, &data); err != nil {
                fmt.Println("Unmarshalling error:",err)
                return nil
        }
        fmt.Printf("<<-- from server:\n %v\n", string(newBuff))
        return data
}
