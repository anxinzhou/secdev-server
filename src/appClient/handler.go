package appClient

import (
	"contract"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"sync"
	"time"
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

func (h *Handler) SigninHandler(data []byte) {
	h.signinHandler(data)
}

func (h *Handler) SignoutHandler(data []byte) {
	h.signoutHandler(data)
}

func (h *Handler) GetWalletsAndMachineHandler(data []byte) {
	h.getWalletsAndMachineHandler(data)
}

func (h *Handler) GetWalletsHandler(data []byte) {
	h.getWalletsHandler(data)
}

func (h *Handler) GetTransactionsHandler(data []byte) {
	h.getTransactionsHandler(data)
}

func (h *Handler) GetExchangeRateHandler(data []byte) {
	h.getExchangeRateHandler(data)
}

func (h *Handler) PostExchangeHandler(data []byte) {
	h.postExchangeHandler(data)
}

func (h *Handler) PostQRCodeHandler(data []byte) {
	h.postRes(PostQRCode)
	time.Sleep(SleepTime)
	h.machineStatusChange(Login)
}

func (h *Handler) MachineLogoutHandler(data []byte) {
	h.postRes(MachineLogout)
	h.machineStatusChange(Logout)
}

func (h *Handler) PostTokenUserOrRewardHandler(data []byte) {
	h.postTokenUserOrRewardHandler(data)
}

func (h *Handler) PostGameStartOrEndHandler(data []byte) {
	h.postGameStartOrEndHandler(data)
}

func (h *Handler) PostTransferHandler(data []byte) {
	h.postTransferHandler(data)
}

func (h *Handler) DisConnect() {
	res := &Disconnect{
		Status: FailedStatus,
		Code:   ServerErrorCode,
		Reason: ServerDisconnect,
	}

	h.wrapperAndSend(DisConnect, res)
}