// Package main provides ...
package main

import (
	app "appClient"
	"contract"
	"db"
	"encoding/json"
	"errors"
	"fmt"
	h "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func appRequestHandler( h *app.Handler) {

	for {
		_, data, err := h.W.ReadMessage()
		var kvs map[string]interface{}
		if err != nil {
			log.Println(err.Error())
			return
		}
		json.Unmarshal(data, &kvs)
		gid, ok := kvs["gcuid"]
		if !ok {
			log.Println(errors.New("gcuid not exist"))
			return
		}

		gcuid := int64(gid.(float64))
		log.Println("gcuid:", gcuid)
		switch gcuid {
		case app.Signin:
			go h.SigninHandler(data)
		case app.Signout:
			go h.SignoutHandler(data)
		case app.GetWalletsAndMachine:
			go h.GetWalletsAndMachineHandler(data)
		case app.GetWallets:
			go h.GetWalletsHandler(data)
		case app.GetTransactions:
			go h.GetTransactionsHandler(data)
		case app.GetExchangeRate:
			go h.GetExchangeRateHandler(data)
		case app.PostExchange:
			go h.PostExchangeHandler(data)
		case app.PostQRCode:
			go h.PostQRCodeHandler(data)
		case app.NotifyMachineStatusChange:
			go h.PostQRCodeHandler(data)
		case app.MachineLogout:
			go h.MachineLogoutHandler(data)
		case app.PostTokenUseOrReward:
			go h.PostTokenUserOrRewardHandler(data)
		case app.PostGameStartOrEnd:
			go h.PostGameStartOrEndHandler(data)
		case app.PostTransfer:
			go h.PostTransferHandler(data)
		}
	}
}

func requestParserWrapper(clientKind string) (func (http.ResponseWriter, *http.Request)) {
	return func(w http.ResponseWriter, r *http.Request){
		log.Println("receive a reqeust")
		var upgrader = websocket.Upgrader{}
		upgrader.CheckOrigin = func(rq *http.Request) bool { return true }
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}

		defer c.Close()
		switch clientKind {
		case "app":
			h := appHandler(c)
			appRequestHandler(h)
			defer h.DisConnect()
		case "web":
		}
	}
}

func appHandler(c *websocket.Conn) *app.Handler{
	const (
		Pbc998ConfigJson  = "./etc/PbcERC998.json"
		Pvc998ConfigJson = "./etc/PvcERC998.json"
		redisConfigJson  = "./etc/redisConfig.json"
	)
	dbClient:= db.NewDBFromJSON(redisConfigJson)
	machine := &app.MachineState{
		State:     app.Logout,
		MachineId: app.MachineId,
	}
	pbcERC998:= contract.NewPbcERC998FromJSON(Pbc998ConfigJson)
	pvcERC998:= contract.NewPvcERC998FromJSON(Pvc998ConfigJson)
	assembledContract := contract.NewAssembleContract(pbcERC998,nil,pvcERC998,nil)
	return app.NewHandler(c, dbClient, machine , assembledContract)
}



func main() {
	// log
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	r := mux.NewRouter()

	const applicationKind = "app"     // app or web
	r.HandleFunc("/", requestParserWrapper(applicationKind)).Methods("GET")

	fmt.Println("Running http server")
	http.ListenAndServe(
		":4000",
		h.CORS(
			h.AllowedMethods([]string{"get", "options", "post", "put", "head"}),
			h.AllowedOrigins([]string{"*"}),
			h.AllowedHeaders([]string{"Content-Type"}),
		)(r),
	)
}
