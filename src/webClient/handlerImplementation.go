package webClient

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/websocket"
	"log"
)

func (h *Handler) writeMessage(messageType int, data []byte) {
	h.mutex.Lock()
	h.W.WriteMessage(messageType, data)
	defer h.mutex.Unlock()
}

func (h *Handler) wrapperAndSend(gcuid int64, res interface{}) {
	resWrapper, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
		h.sendError(gcuid, ServerJsonError)
		return
	}
	h.writeMessage(websocket.TextMessage, resWrapper)
}

func (h *Handler) sendSuccess(gcuid int64) {
	res := &reqSuccess{
			Status: success,
			Gcuid:  gcuid,
		}
	h.wrapperAndSend(gcuid, res)
}

func (h *Handler) sendError(gcuid int64, reason string) {
	reqError := &reqError{
		Status: fail,
		Gcuid:  gcuid,
		Reason: reason,
	}
	reqErrorWrapper, err := json.Marshal(reqError)
	if err != nil {
		log.Println(err.Error())
		h.writeMessage(websocket.TextMessage, []byte(ServerJsonError))
		return
	}
	h.writeMessage(websocket.TextMessage, reqErrorWrapper)
}

func (h *Handler) createERC998Handler(data []byte) {
	var req createERC998Req
	err := json.Unmarshal(data, &req)
	if err != nil {
		log.Println(err.Error())
		h.sendError(createERC998, ClientFormatError)
		return
	}

	owner:=common.HexToAddress(req.Account)
	err =h.c.PvcERC998.Mint(owner,req.TokenId,req.Uri)
	if err!=nil {
		log.Println(err.Error())
		h.sendError(createERC998, "can not create998 in private chain")
		return
	}

	h.wrapperAndSend(createERC998, &createERC998Res{
		Gcuid:createERC998,
		Status:success,
		TokenId: req.TokenId,
	})
}

func (h *Handler) createERC721Handler(data []byte) {
	var req createERC721Req
	err := json.Unmarshal(data, &req)
	if err != nil {
		log.Println(err.Error())
		h.sendError(createERC721, ClientFormatError)
		return
	}

	log.Println("parent tokenId ",req.ParentTokenId)
	owner,err:= h.c.PvcERC998.OwnerOf(req.ParentTokenId)
	if err!=nil {
		log.Println(err.Error())
	}
	log.Println("parent Token owner ",owner.String())
	log.Println("create ", req.Object, " At ", PvcERC721.String())
	err =h.c.PvcERC998.MintERC721Child(req.ParentTokenId,PvcERC721,req.ChildTokenId,req.Uri)
	if err!=nil {
		log.Println(err.Error())
		h.sendError(createERC721, "can not create721 in private chain")
		return
	}
	h.wrapperAndSend(createERC721, &createERC721Res{
		Gcuid:createERC721,
		Status:success,
		Object:req.Object,
		ChildTokenId: req.ChildTokenId,
	})
}

func (h *Handler) transferERC998Handler(data []byte) {
	var req transferERC998Req
	err := json.Unmarshal(data, &req)
	if err != nil {
		log.Println(err.Error())
		h.sendError(transferERC998, ClientFormatError)
		return
	}

	switch req.Chain {
	case public:
		log.Println("unlock in private chain")
		err=h.c.PvcERC998.Unlock(req.TokenId)
		if err!=nil {
			log.Println(err.Error())
			h.sendError(transferERC998,"can not unlock erc998 in private chain")
			return
		}

		log.Println("lock in public chain")
		err=h.c.PbcERC998.Lock(req.TokenId)
		if err!=nil{
			log.Println(err.Error())
			h.sendError(transferERC998, "can not lock erc998 in public chain")
			return
		}
	case private:
		log.Println("lock in private chain")
		err = h.c.PvcERC998.Lock(req.TokenId)
		if err!=nil {
			log.Println(err.Error())
			h.sendError(transferERC998,"can not lock erc998 in private chain")
			return
		}

		// if token not exists in public chain mint first
		//

		if status,err:=h.c.PbcERC998.Exists(req.TokenId); err==nil {
			if !status {
				// not exist  begin to mint
				log.Println("token not exists in public chain, mint first")
				owner,err:=h.c.PvcERC998.OwnerOf(req.TokenId)
				if err!=nil {
					log.Println(err.Error())
					h.sendError(transferERC998,"can not get token owner in private chain")
					return
				}
				uri,err:= h.c.PvcERC998.TokenURI(req.TokenId)
				if err!=nil {
					log.Println(err.Error())
					h.sendError(transferERC998,"can not get token uri in private chain")
					return
				}

				err = h.c.PbcERC998.Mint(owner,req.TokenId,uri)
				if err!=nil {
					log.Println(err.Error())
					h.sendError(transferERC998,"can not mint token in public chain")
					return
				}
			}
		} else {
			log.Println(err.Error())
			h.sendError(transferERC998,"can not get token state")
			return
		}

		log.Println("unlock in public chain")
		err = h.c.PbcERC998.Unlock(req.TokenId)
		if err!=nil{
			log.Println(err.Error())
			h.sendError(transferERC998, "can not unlock erc998 in public chain")
			return
		}
	default:
		log.Println("unknown request type")
		h.sendError(transferERC998,"unknown request type")
		return
	}
	h.wrapperAndSend(transferERC998,&transferERC998Res{
		Gcuid:transferERC998,
		Status:success,
		Chain: req.Chain,
		TokenId:req.TokenId,
	})
}

func (h *Handler) transferERC721Handler(data []byte) {
	var req transferERC721Req
	err := json.Unmarshal(data, &req)
	if err != nil {
		log.Println(err.Error())
		h.sendError(transferERC721, ClientFormatError)
		return
	}

	log.Println("parent tokenId ",req.ParentTokenId)
	owner,err:= h.c.PbcERC998.OwnerOf(req.ParentTokenId)
	if err!=nil {
		log.Println(err.Error())
	}
	log.Println("parent Token owner in public chain ",owner.String())
	log.Println("transfer ", req.Object, " from or to ", PbcERC721.String())

	switch req.Chain {
	case public:
		log.Println("mint erc721 in private chain")
		err=h.c.PvcERC998.MintERC721Child(req.ParentTokenId,PvcERC721,req.ChildTokenId,req.Uri)
		if err!=nil {
			log.Println(err.Error())
			h.sendError(transferERC721,"can not mint erc721 in private chain")
			return
		}
		log.Println("burn erc721 in public chain")
		err=h.c.PbcERC998.BurnERC721Child(req.ParentTokenId,PbcERC721,req.ChildTokenId)
		if err!=nil{
			log.Println(err.Error())
			h.sendError(transferERC721, "can not burn erc721 in public chain")
			return
		}
	case private:
		log.Println("burn erc721 in private chain")
		err = h.c.PvcERC998.BurnERC721Child(req.ParentTokenId,PvcERC721,req.ChildTokenId)
		if err!=nil {
			log.Println(err.Error())
			h.sendError(transferERC721,"can not burn erc721 in private chain")
			return
		}
		log.Println("mint erc721 in public chain")
		err=h.c.PbcERC998.MintERC721Child(req.ParentTokenId,PbcERC721,req.ChildTokenId,req.Uri)
		if err!=nil{
			log.Println(err.Error())
			h.sendError(transferERC721, "can not mint erc721 in public chain")
			return
		}
	default:
		log.Println("unknown request type")
		h.sendError(transferERC721,"unknown request type")
		return
	}

	h.wrapperAndSend(transferERC721,&transferERC721Res{
		Gcuid:transferERC721,
		Status:success,
		Object: req.Object,
		Chain: req.Chain,
		ChildTokenId:req.ChildTokenId,
	})
}