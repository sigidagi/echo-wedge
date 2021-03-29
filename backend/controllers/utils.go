package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	client "github.com/io-m/echo-wedge/backend/tcpClient"
)

var (
	tcpClient = client.TCPClient{
		Host: "localhost",
		Port: 8051,
	}
	in = &client.JSONRPC{}
)

// WedgeCallAllNetworks is helper function for making tcp call to API gtw server for data related to network
func WedgeCallAllNetworks(url string) (*client.JSONRPCResponseNetwork, error) {
	tcpClient.Start()
	in = &client.JSONRPC{
		ID:      uuid.New().String(),
		Jsonrpc: "2.0",
		Method:  http.MethodGet,
		Params: client.Params {
			URL: url,
		},
	}
	b, err := json.Marshal(in)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}	
	if err = tcpClient.Write(b); err != nil {
		return nil, err
	}
	allNetworks := new(client.JSONRPCResponseNetwork)
	netData, err := tcpClient.ReadData(allNetworks)
	if err != nil {
		return nil, err
	}
	data, ok := netData.(*client.JSONRPCResponseNetwork)
	if !ok {
		return nil, nil
	}
	return data, nil
}

func WedgeCallOneNetwork(url string) (*client.JSONRPCResponseNetwork, error) {
	tcpClient.Start()
	in = &client.JSONRPC {
		ID:      uuid.New().String(),
		Jsonrpc: "2.0",
		Method:  http.MethodGet,
		Params: client.Params {
			URL: url,
		},
	}
	b, err := json.Marshal(in)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if err = tcpClient.Write(b); err != nil {
		return nil, err
	}
	oneNetwork := new(client.JSONRPCResponseNetwork)
	netData, err := tcpClient.ReadData(oneNetwork)
	if err != nil {
		return nil, err
	}
	data, ok := netData.(*client.JSONRPCResponseNetwork)
	if !ok {
		return nil, nil
	}
	return data, nil
}



// WedgeCallAllDevices is helper function for making tcp call to API gtw server for data related to device
func WedgeCallAllDevices(url string) (*client.JSONResponseValueDevice, error) {
	tcpClient.Start()
	in = &client.JSONRPC{
		ID:      uuid.New().String(),
		Jsonrpc: "2.0",
		Method:  http.MethodGet,
		Params: client.Params {
			URL: url,
		},
	}
	b, err := json.Marshal(in)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if err = tcpClient.Write(b); err != nil {
		return nil, err
	}
	allDevices := new(client.JSONResponseValueDevice)
	devData, err := tcpClient.ReadData(allDevices)
	if err != nil {
		return nil, err
	}
	data, ok := devData.(*client.JSONResponseValueDevice)
	if !ok {
		return nil, nil
	}
	return data, nil
}
	
func WedgeCallOneDevice(url string) (*client.JSONResponseOneDevice, error) {
	in = &client.JSONRPC{
		ID:      uuid.New().String(),
		Jsonrpc: "2.0",
		Method:  http.MethodGet,
		Params: client.Params {
			URL: url,
		},
	}
	b, err := json.Marshal(in)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if err = tcpClient.Write(b); err != nil {
		return nil, err
	}
	oneDevice := new(client.JSONResponseOneDevice)
	devData, err := tcpClient.ReadData(oneDevice)
	if err != nil {
		return nil, err
	}
	data, ok := devData.(*client.JSONResponseOneDevice)
	if !ok {
		return nil, nil
	}
	return data, nil
}

// WedgeCallAllValues is helper function for making tcp call to API gtw server for data related to value
func WedgeCallAllValues(url string) (*client.JSONResponseValueDevice, error) {
	tcpClient.Start()
	in = &client.JSONRPC{
		ID:      uuid.New().String(),
		Jsonrpc: "2.0",
		Method:  http.MethodGet,
		Params: client.Params {
			URL: url,
		},
	}
	b, err := json.Marshal(in)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if err = tcpClient.Write(b); err != nil {
		return nil, err
	}
	allValues := new(client.JSONResponseValueDevice)
	devData, err := tcpClient.ReadData(allValues)
	if err != nil {
		return nil, err
	}
	data, ok := devData.(*client.JSONResponseValueDevice)
	if !ok {
		return nil, nil
	}
	return data, nil
}
func WedgeCallOneValue(url string) (*client.JSONResponseOneValue, error) {
	in = &client.JSONRPC{
		ID:      uuid.New().String(),
		Jsonrpc: "2.0",
		Method:  http.MethodGet,
		Params: client.Params {
			URL: url,
		},
	}
	b, err := json.Marshal(in)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if err = tcpClient.Write(b); err != nil {
		return nil, err
	}
	oneValue := new(client.JSONResponseOneValue)
	devData, err := tcpClient.ReadData(oneValue)
	if err != nil {
		return nil, err
	}
	data, ok := devData.(*client.JSONResponseOneValue)
	if !ok {
		return nil, nil
	}
	return data, nil
}
	
// WedgeCallAllStates is helper function for making tcp call to API gtw server for data related to state
func WedgeCallAllStates(url string) (*client.JSONResponseState, error) {
	tcpClient.Start()
	in = &client.JSONRPC{
		ID:      uuid.New().String(),
		Jsonrpc: "2.0",
		Method:  http.MethodGet,
		Params: client.Params {
			URL: url,
		},
	}
	b, err := json.Marshal(in)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if err = tcpClient.Write(b); err != nil {
		return nil, err
	}
	allStates := new(client.JSONResponseState)
	devData, err := tcpClient.ReadData(allStates)
	if err != nil {
		return nil, err
	}
	data, ok := devData.(*client.JSONResponseState)
	if !ok {
		return nil, nil
	}
	return data, nil
}
func WedgeCallOneState(url string) (*client.JSONResponseOneState, error) {
	in = &client.JSONRPC{
		ID:      uuid.New().String(),
		Jsonrpc: "2.0",
		Method:  http.MethodGet,
		Params: client.Params {
			URL: url,
		},
	}
	b, err := json.Marshal(in)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if err = tcpClient.Write(b); err != nil {
		return nil, err
	}
	oneState := new(client.JSONResponseOneState)
	devData, err := tcpClient.ReadData(oneState)
	if err != nil {
		return nil, err
	}
	data, ok := devData.(*client.JSONResponseOneState)
	if !ok {
		return nil, nil
	}
	return data, nil
}
