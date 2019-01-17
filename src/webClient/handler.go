package webClient

import (
	"contract"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type Handler struct {
	W *websocket.Conn
	c *contract.AssembleContract
	mutex *sync.Mutex
}

func NewHandler(w *websocket.Conn, c* contract.AssembleContract) *Handler {
	return &Handler{
		W:w,
		c:c,
		mutex: &sync.Mutex{},
	}
}

func (h *Handler) HandlerRequest() {
	for {
		_, data, err := h.W.ReadMessage()
		var kvs map[string]interface{}
		if err != nil {
			log.Println(err.Error())
			continue
		}
		err =json.Unmarshal(data, &kvs)
		if err!=nil {
			log.Println(err.Error())
		}
		gid, ok := kvs["gcuid"]
		if !ok {
			log.Println(errors.New("gcuid not exist"))
			continue
		}

		gcuid := int64(gid.(float64))
		log.Println("request gcuid:", gcuid)
		switch gcuid {
		case createERC998:
			go h.createERC998Handler(data)
		case createERC721:
			go h.createERC721Handler(data)
		case transferERC998:
			go h.transferERC998Handler(data)
		case transferERC721:
			go h.transferERC721Handler(data)
		}
	}
}