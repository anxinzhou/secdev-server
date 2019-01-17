package appClient

import (
	"contract"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type Handler struct {
	W *websocket.Conn
	db      *redis.Client
	machine *MachineState
	c *contract.AssembleContract
	mutex *sync.Mutex
}

func NewHandler(w *websocket.Conn, db *redis.Client, machine *MachineState, c* contract.AssembleContract) *Handler {
	return &Handler{
		W:w,
		db:db,
		machine:machine,
		c:c,
		mutex: &sync.Mutex{},
	}
}

func (h *Handler) DisConnect() {
	res := &Disconnect{
		Status: FailedStatus,
		Code:   ServerErrorCode,
		Reason: ServerDisconnect,
	}

	h.wrapperAndSend(DisConnect, res)
}

func (h *Handler) HandlerRequest() {
	for {
		_, data, err := h.W.ReadMessage()
		var kvs map[string]interface{}
		if err != nil {
			log.Println(err.Error())
			continue
		}
		json.Unmarshal(data, &kvs)
		gid, ok := kvs["gcuid"]
		if !ok {
			log.Println(errors.New("gcuid not exist"))
			continue
		}

		gcuid := int64(gid.(float64))
		log.Println("gcuid:", gcuid)
		switch gcuid {
		case Signin:
			go h.signinHandler(data)
		case Signout:
			go h.signoutHandler(data)
		case GetWalletsAndMachine:
			go h.getWalletsAndMachineHandler(data)
		case GetWallets:
			go h.getWalletsHandler(data)
		case GetTransactions:
			go h.getTransactionsHandler(data)
		case GetExchangeRate:
			go h.getExchangeRateHandler(data)
		case PostExchange:
			go h.postExchangeHandler(data)
		case PostQRCode:
			go h.postQRCodeHandler(data)
		case MachineLogout:
			go h.machineLogoutHandler(data)
		case PostTokenUseOrReward:
			go h.postTokenUserOrRewardHandler(data)
		case PostGameStartOrEnd:
			go h.postGameStartOrEndHandler(data)
		case PostTransfer:
			go h.postTransferHandler(data)
		}
	}
}