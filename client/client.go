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
	ReadNetwork() *JSONRPCResponseNetwork
	ReadDeviceOne() *JSONResponseOneDevice
	ReadDeviceValue() *JSONResponseValueDevice
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

// ReadNetwork is method for sending Network details to frontend client
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

// ReadDeviceValue is method for sending Device or Value list to frontend client
func (c *TCPClient) ReadDeviceValue() *JSONResponseValueDevice {
        data := new(JSONResponseValueDevice)
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

// ReadDeviceOne is method for sending Device detail to frontend client
func (c *TCPClient) ReadDeviceOne() *JSONResponseOneDevice {
        data := new(JSONResponseOneDevice)
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

// ReadValueOne is method for sending Network details to frontend client
func (c *TCPClient) ReadValueOne() *JSONResponseOneValue{
        data := new(JSONResponseOneValue)
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
        fmt.Printf("<<-- from server VALUE data:\t %v\n", string(newBuff))
        return data
}

// ReadState is method for sending State details to frontend client
func (c *TCPClient) ReadState() *JSONResponseState{
        data := new(JSONResponseState)
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
        fmt.Printf("<<-- from server VALUE data:\t %v\n", string(newBuff))
        return data
}
// ReadStateOne is method for sending State details to frontend client
func (c *TCPClient) ReadStateOne() *JSONResponseOneState{
        data := new(JSONResponseOneState)
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
        fmt.Printf("<<-- from server VALUE data:\t %v\n", string(newBuff))
        return data
}

