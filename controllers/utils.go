package controllers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/io-m/echo-wedge/client"
)

var in = &client.JSONRPC{}

func findExact(id string, slice []string) (string, error) {
	var exactID string
	var err error
	for _, v := range slice {
		if id == v {
			exactID = v
			return exactID, nil
		}
	}
	return "", err
}

// WedgeCall is reusable function for reaching out to WedgeAPI gtw with tcp client calls
func WedgeCall(model string, method string, data *client.StateWg, modelID []string) (*client.JSONRPCResponse, *client.JSONRPCResponseNetwork, error) {
	method = strings.ToUpper(method)
	model = strings.ToLower(model)
	var tcpClient = client.TCPClient{
		Host: "localhost",
		Port: 8051,
	}
	tcpClient.Start()
		
	switch {
	case len(modelID) == 0:
		switch model {
			case "network":
				in = &client.JSONRPC {
					ID:      uuid.New().String(),
					Jsonrpc: "2.0",
					Method:  method,
					Params: client.Params {
						URL: fmt.Sprintf("/"),
					},
				}
				break
			}
		
	case len(modelID) == 1:
		if method == "GET" {
			switch model {
			case "device":
				netID, err := findExact(modelID[0], modelID)
				if err != nil {
					return nil, nil, err
				}
				in = &client.JSONRPC {
					ID:      uuid.New().String(),
					Jsonrpc: "2.0",
					Method:  method,
					Params: client.Params{
						URL: fmt.Sprintf("/network/%s/device", netID),
					},
				}
				break
			case "network":
				netID, err := findExact(modelID[0], modelID)
				if err != nil {
					return nil, nil, err
				}
				in = &client.JSONRPC {
					ID:      uuid.New().String(),
					Jsonrpc: "2.0",
					Method:  method,
					Params: client.Params{
						URL: fmt.Sprintf("/network/%s", netID),
					},
				}
			}
		}	
	case len(modelID) == 2:
		if method == "GET" {
			switch model {
			case "device":
				netID, err := findExact(modelID[0], modelID)
				if err != nil {
					return nil, nil, err
				}
				devID, err := findExact(modelID[1], modelID)
				if err != nil {
					return nil, nil, err
				}
				in = &client.JSONRPC {
					ID:      uuid.New().String(),
					Jsonrpc: "2.0",
					Method:  method,
					Params: client.Params{
						URL: fmt.Sprintf("/network/%s/device/%s", netID, devID),
					},
				}
				break
			case "value":
				
				netID, err := findExact(modelID[0], modelID)
				if err != nil {
					return nil, nil, err
				}
				devID, err := findExact(modelID[1], modelID)
				if err != nil {
					return nil, nil, err
				}
				in = &client.JSONRPC {
					ID:      uuid.New().String(),
					Jsonrpc: "2.0",
					Method:  method,
					Params: client.Params {
						URL: fmt.Sprintf("/network/%s/device/%s/value", netID, devID),
					},
				}
				break
			}
		}
		case len(modelID) == 3:
			if method == "GET" {
				switch model {
					case "value":
					
					netID, err := findExact(modelID[0], modelID)
					if err != nil {
						return nil, nil, err
					}
					devID, err := findExact(modelID[1], modelID)
					if err != nil {
						return nil, nil, err
					}
					valID, err := findExact(modelID[2], modelID)
					if err != nil {
						return nil, nil, err
					}
					in = &client.JSONRPC {
						ID:      uuid.New().String(),
						Jsonrpc: "2.0",
						Method:  method,
						Params: client.Params {
							URL: fmt.Sprintf("/network/%s/device/%s/value/%s", netID, devID, valID),
						},
					}
					break
				case "state":
					
					netID, err := findExact(modelID[0], modelID)
					if err != nil {
						return nil, nil, err
					}
					devID, err := findExact(modelID[1], modelID)
					if err != nil {
						return nil, nil, err
					}
					valID, err := findExact(modelID[2], modelID)
					if err != nil {
						return nil, nil, err
					}
					in = &client.JSONRPC {
						ID:      uuid.New().String(),
						Jsonrpc: "2.0",
						Method:  method,
						Params: client.Params {
							URL: fmt.Sprintf("/network/%s/device/%s/value/%s/state", netID, devID, valID),
						},
					}
					break
				}
			}
		case len(modelID) == 4:
			if method == "GET" {
				switch model {
					case "state":
						
						netID, err := findExact(modelID[0], modelID)
						if err != nil {
							return nil, nil, err
						}
						devID, err := findExact(modelID[1], modelID)
						if err != nil {
							return nil, nil, err
						}
						valID, err := findExact(modelID[2], modelID)
						if err != nil {
							return nil, nil, err
						}
						stateID, err := findExact(modelID[3], modelID)
						if err != nil {
							return nil, nil, err
						}
						in = &client.JSONRPC {
							ID:      uuid.New().String(),
							Jsonrpc: "2.0",
							Method:  method,
							Params: client.Params{
								URL: fmt.Sprintf("/network/%s/device/%s/value/%s/state/%s", netID, devID, valID, stateID),
							},
						}
						break
					}
			} else if method == "PUT" {
				switch model {
					case "state":
					
						netID, err := findExact(modelID[0], modelID)
						if err != nil {
							return nil, nil, err
						}
						devID, err := findExact(modelID[1], modelID)
						if err != nil {
							return nil, nil, err
						}
						valID, err := findExact(modelID[2], modelID)
						if err != nil {
							return nil, nil, err
						}
						stateID, err := findExact(modelID[3], modelID)
						if err != nil {
							return nil, nil, err
						}
						in = &client.JSONRPC {
							ID:      uuid.New().String(),
							Jsonrpc: "2.0",
							Method:  method,
							Params: client.Params{
								URL: fmt.Sprintf("/network/%s/device/%s/value/%s/state/%s", netID, devID, valID, stateID),
								Data: data,
							},
						}
					break
				}
			}
	}
	
	b, err := json.Marshal(in)
		if err != nil {
			fmt.Println(err.Error())
			return nil, nil, err
		}
	
	if err = tcpClient.Write(b); err != nil {
		return nil, nil, err
	}
	if model == "network" {
		netData := tcpClient.ReadNetwork()
		return nil, netData, nil
	}

	recData := tcpClient.Read()

	return recData, nil, nil
}
