package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	ws "github.com/gorilla/websocket"
	m "github.com/io-m/echo-wedge/backend/models"
	"github.com/labstack/echo/v4"
)

type eventController struct{}

var (
	upgrader ws.Upgrader
	getChan  = make(chan *m.State)
)

// What we get from API gtw JSON-RPC response
type APIResponse struct {
	Body string // JSON rpc response body
}

func NewEvent() *eventController {
	return &eventController{}
}

func (*eventController) Receive(c echo.Context) error {

	// response := &StateWg{}
	response := &m.State{}
	dataBack, _ := ioutil.ReadAll(c.Request().Body)
	if err := json.Unmarshal(dataBack, &response); err != nil {
		log.Println("Could not convert received data from TCP server to responseState struct")
	}
	log.Println("===========================")
	b, _ := json.Marshal(response)
	log.Println(string(b))
	go func() {
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
	go func() {
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
