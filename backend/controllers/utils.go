package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	m "github.com/io-m/echo-wedge/backend/models"
	cl "github.com/io-m/echo-wedge/backend/tcpClient"
	"net/http"
)

func WedgeWrite(url string) error {

	in := &m.JSONRPC{
		ID:      uuid.New().String(),
		Jsonrpc: "2.0",
		Method:  http.MethodGet,
		Params: m.Params{
			URL: url,
		},
	}
	b, err := json.Marshal(in)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if err = cl.Backend.Write(b); err != nil {
		return err
	}
	return nil
}

func WedgeCallNetwork(url string) (*m.Network, error) {

	data := &m.JSONResponseNetwork{}
	if err := WedgeWrite(url); err != nil {
		return &m.Network{}, err
	}

	netData, err := cl.Backend.ReadData(data)
	if err != nil {
		return &m.Network{}, err
	}

	jrpc := netData.(*m.JSONResponseNetwork)
	if jrpc.Error == (m.Error{}) {
		return &jrpc.Result.Value, nil
	} else {
		return &m.Network{}, errors.New(jrpc.Error.Message)
	}
}

func WedgeCallDevice(url string) (*m.Device, error) {

	data := &m.JSONResponseDevice{}
	if err := WedgeWrite(url); err != nil {
		return &m.Device{}, err
	}

	devData, err := cl.Backend.ReadData(data)
	if err != nil {
		return &m.Device{}, err
	}

	jrpc := devData.(*m.JSONResponseDevice)
	return &jrpc.Result.Value, nil
}

func WedgeCallItemList(url string) (*m.ItemList, error) {

	data := &m.JSONResponseItemList{}
	if err := WedgeWrite(url); err != nil {
		return &m.ItemList{}, err
	}

	devData, err := cl.Backend.ReadData(data)
	if err != nil {
		return &m.ItemList{}, err
	}

	jrpc := devData.(*m.JSONResponseItemList)
	return &jrpc.Result.Value, nil
}

func WedgeCallValue(url string) (*m.Value, error) {

	data := &m.JSONResponseValue{}
	if err := WedgeWrite(url); err != nil {
		return &m.Value{}, err
	}

	devData, err := cl.Backend.ReadData(data)
	if err != nil {
		return &m.Value{}, err
	}

	jrpc := devData.(*m.JSONResponseValue)
	return &jrpc.Result.Value, nil
}

func WedgeCallState(url string) (*m.State, error) {

	data := &m.JSONResponseState{}
	if err := WedgeWrite(url); err != nil {
		return &m.State{}, err
	}

	devData, err := cl.Backend.ReadData(data)
	if err != nil {
		return &m.State{}, err
	}

	jrpc := devData.(*m.JSONResponseState)
	return &jrpc.Result.Value, nil
}
