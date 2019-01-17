// Package main provides ...
package main

import (
	"appClient"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"webClient"
	"contract"
	"db"
	"fmt"
	h "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

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
			h.HandlerRequest()
			defer h.DisConnect()
		case "web":
			h:= webHandler(c)
			h.HandlerRequest()
		default:
			log.Fatalln("wrong client kind")
		}
	}
}

func appHandler(c *websocket.Conn) *appClient.Handler{
	const (
		Pbc998ConfigJson  = "./etc/PbcERC998.json"
		Pvc998ConfigJson = "./etc/PvcERC998.json"
		redisConfigJson  = "./etc/redisConfig.json"
	)
	dbClient:= db.NewDBFromJSON(redisConfigJson)
	machine := &appClient.MachineState{
		State:     appClient.Logout,
		MachineId: appClient.MachineId,
	}
	pbcERC998:= contract.NewPbcERC998FromJSON(Pbc998ConfigJson)
	pvcERC998:= contract.NewPvcERC998FromJSON(Pvc998ConfigJson)
	assembledContract := contract.NewAssembleContract(pbcERC998,nil,pvcERC998,nil)
	return appClient.NewHandler(c, dbClient, machine , assembledContract)
}

// only for test child contract
func loadChildAddress(configFile string) common.Address {
	// read child contract Address, only for test
	type basicChainConfig struct {
		Address  string `json:"address"`
	}
	// initial contract
	var c basicChainConfig
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Println(err.Error())
		panic("can not read chain config file")
	}
	json.Unmarshal(data, &c)
	return common.HexToAddress(c.Address)
}

func webHandler(c *websocket.Conn) *webClient.Handler {
	const (
		Pbc998ConfigJson  = "./etc/PbcERC998.json"
		Pbc721ConfigJson = "./etc/PbcERC721.json"
		Pvc998ConfigJson = "./etc/PvcERC998.json"
		Pvc721ConfigJson = "./etc/PvcERC721.json"
	)

	webClient.PvcERC721 = loadChildAddress(Pvc721ConfigJson)
	webClient.PbcERC721 = loadChildAddress(Pbc721ConfigJson)
	pbcERC998:= contract.NewPbcERC998FromJSON(Pbc998ConfigJson)
	pvcERC998:= contract.NewPvcERC998FromJSON(Pvc998ConfigJson)
	assembledContract := contract.NewAssembleContract(pbcERC998,nil,pvcERC998,nil)
	return webClient.NewHandler(c, assembledContract)
}

func main() {
	// log
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	r := mux.NewRouter()

	const applicationKind = "web"     // app or web
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
