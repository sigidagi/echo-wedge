package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	ws "github.com/gorilla/websocket"
	"github.com/io-m/echo-wedge/backend/models"
	"github.com/labstack/echo/v4"
)

type eventController struct{}

var (
	upgrader ws.Upgrader
	getChan = make(chan *models.State)
	// getChan = make(chan *StateWg)

)

// type Meta struct {
//     // state         protoimpl.MessageState
//     // sizeCache     protoimpl.SizeCache
//     // unknownFields protoimpl.UnknownFields

//     Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
//     Type    string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
//     Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
// }
// type NetworkWg struct {
// 	Name      string      `json:"name"`
// 	DevicesWg []*DeviceWg `json:"device"`
// 	Meta      *Meta   `json:"meta"`
// }

// type DeviceWg struct {
// 	Parent   *NetworkWg `json:"-"`
// 	Id       uint32     `json:"-"`
// 	Url      string     `json:"-"`
// 	ValuesWg []*ValueWg `json:"value"`
// 	NodeId   string     `json:"-"`
// 	*models.Device
// }

// type ValueWg struct {
// 	Parent   *DeviceWg  `json:"-"`
// 	Id       uint32     `json:"-"`
// 	Url      string     `json:"-"`
// 	StatesWg []*StateWg `json:"state"`
// 	NodeId   string     `json:"-"`
// 	*models.Value
// }

// type StateWg struct {
// 	Parent *ValueWg `json:"-"`
// 	Id     uint32   `json:"-"`
// 	Url    string   `json:"-"`
// 	NodeId string   `json:"-"`
// 	*models.State
// }

// What we get from API gtw JSON-RPC response
type APIResponse struct {
	Body string // JSON rpc response body
}

func NewEvent() *eventController {
	return &eventController{}
}

func (*eventController) Receive(c echo.Context) error {
	
	// response := &StateWg{}
	response := &models.State{}
	dataBack, _ := ioutil.ReadAll(c.Request().Body)
		if err := json.Unmarshal(dataBack, &response); err != nil {
			log.Println("Could not convert received data from TCP server to responseState struct")
		}
		log.Println("===========================")
		log.Println(response)
	go func(){
		getChan <- response
	}()

	return nil
}

func (*eventController) PushUpdates(c echo.Context) error {
	w := c.Response()
	r := c.Request()
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Cannot establish WS connection: ", err)
	}
	log.Println("Connection established")
	go func(){
		defer conn.Close()
		for {
			response := <-getChan // here is the key -> it blocks as long as it does not get new value of response
			if response.Meta.ID != "" {
				if err := conn.WriteJSON(response); err != nil {
					http.Error(w, err.Error(), http.StatusNotFound)
					break
				}
			} else {
				break
			}
		}
	}()
	return nil
}