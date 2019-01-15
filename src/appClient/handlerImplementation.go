package appClient

import (
	"contract"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"math/big"
	"strings"
	"utils"
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
		h.sendError(gcuid, ServerErrorCode, ServerJsonError)
		return
	}
	h.writeMessage(websocket.TextMessage, resWrapper)
}

func (h *Handler) sendError(gcuid int64, errorCode int64, reason string) {
	reqError := &Error{
		Status: FailedStatus,
		Code:   errorCode,
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

func (h *Handler) signinHandler(data []byte) {

	var signin SigninReq
	err := json.Unmarshal(data, &signin)
	if err != nil {
		log.Println(err.Error())
		h.sendError(Signin, ClientErrorCode, ClientFormatError)
		return
	}
	if signin.LoginCode != LoginCode {
		h.sendError(Signin, ClientErrorCode, WrongUser)
		return
	}

	res := &SigninRes{
		Status: OkStatus,
		Gcuid:  Signin,
		Guuid:  Guuid,
	}

	h.wrapperAndSend(Signin, res)
}

func (h *Handler) signoutHandler(data []byte) {
	h.postRes(Signout)
}

func (h *Handler) getWallet() ([]*Wallet, *TransferInfo, error) {
	ether, err := h.c.PbcERC20.GetEther(UserAddr)
	if err != nil {
		return nil, nil, err
	}
	token, err := h.c.PbcERC20.GetToken(UserAddr)
	if err != nil {
		log.Println(err.Error())
		return nil, nil, err
	}

	tokenWallet := &Wallet{
		Name:   TokenWalletName,
		Amount: new(big.Float).Quo(new(big.Float).SetInt(token), TokenBase).String(),
		Id:     User,
	}
	etherWallet := &Wallet{
		Name:   EthWalletName,
		Amount: new(big.Float).Quo(new(big.Float).SetInt(ether), EtherBase).String(),
		Id:     User,
	}

	wallets := make([]*Wallet, 0, WalletCount)
	wallets = append(wallets, etherWallet, tokenWallet)

	rawTransferAmount, err := h.c.PvcERC20.GetToken(UserAddr)
	if err != nil {
		log.Println(err.Error())
		return nil, nil, err
	}
	transferAmount := new(big.Float).Quo(new(big.Float).SetInt(rawTransferAmount), TokenBase).String()
	toTransfer := &TransferInfo{
		Amount: transferAmount,
		Name:   TokenWalletName,
	}
	return wallets, toTransfer, nil
}

func (h *Handler) postRes(gcuid int64) {
	res := &PostTokenUserOrRewardRes{
		PostRes: &PostRes{
			Status: OkStatus,
			Gcuid:  gcuid,
		},
	}
	h.wrapperAndSend(gcuid, res)
}

func (h *Handler) getWalletsAndMachineHandler(data []byte) {

	wallets, toTransfer, err := h.getWallet()
	if err != nil {
		log.Println(err.Error())
		h.sendError(GetWalletsAndMachine, ServerErrorCode, ServerChainError)
		return
	}
	res := &GetWalletsAndMachineRes{
		Gcuid: GetWalletsAndMachine,
		Machine: &MachineState{
			State:     h.machine.State,
			MachineId: h.machine.MachineId,
		},
		WalletsCount: WalletCount,
		Wallets:      wallets,
		Transfer:     toTransfer,
	}

	h.wrapperAndSend(GetWalletsAndMachine, res)
}

func (h *Handler) getWalletsHandler(data []byte) {
	wallets, toTransfer, err := h.getWallet()
	if err != nil {
		log.Println(err.Error())
		h.sendError(GetWallets, ServerErrorCode, ServerChainError)
		return
	}
	res := &GetWalletRes{
		Gcuid:    GetWallets,
		Count:    WalletCount,
		Wallets:  wallets,
		Transfer: toTransfer,
	}

	h.wrapperAndSend(GetWallets, res)
}

func (h *Handler) getTransactionsHandler(data []byte) {
	var req GetTransactionsReq
	err := json.Unmarshal(data, &req)
	if err != nil {
		h.sendError(GetTransactions, ClientErrorCode, ClientFormatError)
		return
	}

	start := req.After
	stop := req.After + req.Amount
	var vals []string
	var collection string
	switch req.From {
	case ETH:
		collection = strings.ToLower(User) + ":" + string(ETH)
	case Slot:
		collection = strings.ToLower(User) + ":" + string(Slot)
	}
	totalLen, err := h.db.LLen(collection).Result()
	log.Println("transaction len:", totalLen)
	if err != nil {
		log.Println(err.Error())
		h.sendError(GetTransactions, ClientErrorCode, err.Error())
		return
	}

	count := totalLen - start
	if count > req.Amount {
		count = req.Amount
	}

	if count < 0 {
		count = 0
	}

	vals, err = h.db.LRange(collection, start, stop).Result()
	log.Println(vals)
	if err != nil {
		log.Println(err.Error())
		h.sendError(GetTransactions, ClientErrorCode, err.Error())
		return
	}
	txs := make([]*Transaction, len(vals))
	for i, val := range vals {
		var tx Transaction
		err = json.Unmarshal([]byte(val), &tx)
		if err != nil {
			h.sendError(GetTransactions, ServerErrorCode, ServerJsonError)
			return
		}
		txs[i] = &tx
	}

	res := &GetTransactionRes{
		Gcuid:        GetTransactions,
		Count:        count,
		After:        req.After,
		From:         req.From,
		Transactions: txs,
	}

	h.wrapperAndSend(GetTransactions, res)
}

func (h *Handler) getExchangeRateHandler(data []byte) {
	var req GetExchangeRateReq
	err := json.Unmarshal(data, &req)
	if err != nil {
		h.sendError(GetExchangeRate, ClientErrorCode, ClientFormatError)
		return
	}

	rate, err := h.c.PbcERC20.GetExchangeRate()
	if err != nil {
		log.Println(err.Error())
		h.sendError(GetExchangeRate, ServerErrorCode, ServerChainError)
		return
	}
	switch req.ExchangeType {
	case SlotToEth:
		rateWrapper, _, _ := new(big.Float).Parse(rate, 10)
		rate = new(big.Float).Quo(big.NewFloat(1), rateWrapper).String()
	}

	res := &GetExchangeRateRes{
		Gcuid:        GetExchangeRate,
		ExchangeType: req.ExchangeType,
		Rate:         rate,
	}

	h.wrapperAndSend(GetExchangeRate, res)
}

func (h *Handler) exchangeResultHandler(req *PostExchangeReq, err error) {

	var status string

	if err != nil {
		log.Println(err.Error())
		status = FailedStatus
	} else {
		status = SuccessStatus
	}

	if status == SuccessStatus {
		amountWrapper, _, _ := new(big.Float).Parse(req.Amount, 10)
		rawExchangeRate, _ := h.c.PbcERC20.GetExchangeRate()
		exchangeRate, _, _ := new(big.Float).Parse(rawExchangeRate, 10)
		// "2006-01-02 15:04:05 is the birth day of golang, fixed format"
		date := utils.GetCurrentTime()
		switch req.ExchangeType {
		case EthToSLot:
			withdrawTokenAmount := new(big.Float).Mul(amountWrapper, exchangeRate)
			ethTx := Transaction{
				Type:        Withdraw,
				Amount:      req.Amount,
				CreatedDate: date,
			}
			slotTx := Transaction{
				Type:        Deposit,
				Amount:      withdrawTokenAmount.String(),
				CreatedDate: date,
			}
			slotTxWrapper, err := json.Marshal(&slotTx)
			ethTxWrapper, err := json.Marshal(&ethTx)
			_, err = h.db.LPush(strings.ToLower(User)+":"+string(ETH), ethTxWrapper).Result()
			if err != nil {
				log.Println(err.Error())
			}
			_, err = h.db.LPush(strings.ToLower(User)+":"+string(Slot), slotTxWrapper).Result()
			if err != nil {
				log.Println(err.Error())
			}
		case SlotToEth:
			withdrawEthAmount := new(big.Float).Quo(amountWrapper, exchangeRate)
			ethTx := Transaction{
				Type:        Deposit,
				Amount:      withdrawEthAmount.String(),
				CreatedDate: date,
			}
			slotTx := Transaction{
				Type:        Withdraw,
				Amount:      req.Amount,
				CreatedDate: date,
			}
			slotTxWrapper, _ := json.Marshal(&slotTx)
			ethTxWrapper, _ := json.Marshal(&ethTx)
			_, err = h.db.LPush(strings.ToLower(User)+":"+string(ETH), ethTxWrapper).Result()
			if err != nil {
				log.Println(err.Error())
			}
			_, err = h.db.LPush(strings.ToLower(User)+":"+string(Slot), slotTxWrapper).Result()
			if err != nil {
				log.Println(err.Error())
			}
		}

		feeTx := Transaction{
			Type:        Gas,
			Amount:      contract.TxFee,
			CreatedDate: date,
		}

		feeTxWrapper, _ := json.Marshal(&feeTx)
		_, err = h.db.LPush(strings.ToLower(User)+":"+string(ETH), feeTxWrapper).Result()
		if err != nil {
			log.Println(err.Error())
		}
	}

	resultRes := &ExchangeResultRes{
		Gcuid:        NotifyExchangeResult,
		Status:       status,
		ExchangeType: req.ExchangeType,
		Amount:       req.Amount,
		CreatedDate:  utils.GetCurrentTime(),
	}

	h.wrapperAndSend(NotifyTokenChange, resultRes)
}

func (h *Handler) postExchangeHandler(data []byte) {
	var req PostExchangeReq
	err := json.Unmarshal(data, &req)
	if err != nil {
		log.Println(err.Error())
		h.sendError(PostExchange, ClientErrorCode, ClientFormatError)
		return
	}

	amount, _, err := new(big.Float).Parse(req.Amount, 10)
	if err != nil {
		log.Println(err.Error())
		h.sendError(PostExchange, ClientErrorCode, ClientFormatError)
		return
	}

	h.postRes(PostExchange)
	switch req.ExchangeType {
	case EthToSLot:
		amountWrapper, _ := new(big.Float).Mul(amount, EtherBase).Int(nil)
		err = h.c.PbcERC20.ExchangeForToken(UserAddr, amountWrapper)
	case SlotToEth:
		amountWrapper, _ := new(big.Float).Mul(amount, TokenBase).Int(nil)
		err = h.c.PbcERC20.ExchangeForEther(UserAddr, amountWrapper)
	default:
		err = errors.New("unknown exchange type")
	}
	h.exchangeResultHandler(&req, err)
}

func (h *Handler) machineStatusChange(state int64) {
	res := &MachineStatusChangeRes{
		Gcuid:       NotifyMachineStatusChange,
		State:       state,
		MachineId:   MachineId,
		Act:         MachineStatusChange,
		CreatedDate: utils.GetCurrentTime(),
	}

	h.wrapperAndSend(NotifyMachineStatusChange, res)
}

func (h *Handler) postTokenUserOrRewardHandler(data []byte) {
	var req PostTokenUseOrRewardReq
	err := json.Unmarshal(data, &req)
	if err != nil {
		log.Println(err.Error())
		h.sendError(PostTokenUseOrReward, ClientErrorCode, ClientFormatError)
	}
	amount, _, err := new(big.Float).Parse(req.Amount, 10)
	if err != nil {
		log.Println(err.Error())
		h.sendError(PostTokenUseOrReward, ClientErrorCode, ClientFormatError)
		return
	}

	amountWrapper, _ := new(big.Float).Mul(amount, TokenBase).Int(nil)
	switch req.Type {
	case Used:
		err = h.c.PvcERC20.Consume(UserAddr, amountWrapper)
	case Reward:
		err = h.c.PvcERC20.Reward(UserAddr, amountWrapper)
	}
	if err != nil {
		log.Println(err.Error())
		h.sendError(PostTokenUseOrReward, ServerErrorCode, err.Error())
		return
	}

	h.postRes(PostTokenUseOrReward)
}

func (h *Handler) postGameStartOrEndHandler(data []byte) {
	var req PostGameStartOrEndReq
	err := json.Unmarshal(data, &req)
	if err != nil {
		log.Println(err.Error())
		h.sendError(PostGameStartOrEnd, ClientErrorCode, ClientFormatError)
	}
	amount, _, err := new(big.Float).Parse(req.Amount, 10)
	if err != nil {
		log.Println(err.Error())
		h.sendError(PostTokenUseOrReward, ClientErrorCode, ClientFormatError)
		return
	}

	amountWrapper, _ := new(big.Float).Mul(amount, TokenBase).Int(nil)
	switch req.Type {
	case GameStart:
		err = h.c.PvcERC20.Mint(UserAddr, amountWrapper)
		if err != nil {
			log.Println(err.Error())
			h.sendError(PostGameStartOrEnd, ServerErrorCode, err.Error())
			return
		}
		h.postRes(PostGameStartOrEnd)

		err := h.c.PbcERC20.Burn(UserAddr, amountWrapper)
		if err != nil {
			log.Println(err.Error())
			h.c.PvcERC20.Burn(UserAddr, amountWrapper)
		}

		date := utils.GetCurrentTime()

		feeTx := Transaction{
			Type:        Gas,
			Amount:      contract.TxFee,
			CreatedDate: date,
		}

		feeTxWrapper, _ := json.Marshal(&feeTx)
		_, err = h.db.LPush(strings.ToLower(User)+":"+string(ETH), feeTxWrapper).Result()
		if err != nil {
			log.Println(err.Error())
			return
		}

		tokenUpdate := amount.String()
		tx := Transaction{
			Type:        Spend,
			Amount:      tokenUpdate,
			CreatedDate: date,
		}
		txWrapper, _ := json.Marshal(tx)
		_, err = h.db.LPush(strings.ToLower(User)+":"+string(Slot), txWrapper).Result()
		if err != nil {
			log.Println(err.Error())
			return
		}

		rawTokenTotal, err := h.c.PbcERC20.GetToken(UserAddr)
		if err != nil {
			log.Println(err.Error())
			return
		}
		tokenTotal := new(big.Float).Quo(new(big.Float).SetInt(rawTokenTotal), TokenBase).String()

		resultRes := &TokenCountChangeRes{
			Gcuid:       NotifyTokenChange,
			State:       Decrease,
			WalletName:  TokenWalletName,
			WalletId:    User,
			TokenUpdate: tokenUpdate,
			TokenTotal:  tokenTotal,
			Act:         TokenChange,
			CreatedDate: date,
		}

		h.wrapperAndSend(NotifyTokenChange, resultRes)
	case GameEnd:

		h.postRes(PostGameStartOrEnd)
	}
}

func (h *Handler) postTransferHandler(data []byte) {

	rawToken, err := h.c.PvcERC20.GetToken(UserAddr)
	if err != nil {
		log.Println(err.Error())
		h.sendError(PostTransfer, ServerErrorCode, err.Error())
		return
	}

	err = h.c.PvcERC20.Burn(UserAddr, rawToken)
	if err != nil {
		log.Println(err.Error())
		h.sendError(PostTransfer, ServerErrorCode, err.Error())
		return
	}
	h.postRes(PostTransfer)

	err = h.c.PbcERC20.Mint(UserAddr, rawToken)
	if err != nil {
		log.Println(err.Error())
		h.c.PvcERC20.Mint(UserAddr, rawToken)
		return
	}

	date := utils.GetCurrentTime()

	feeTx := Transaction{
		Type:        Gas,
		Amount:      contract.TxFee,
		CreatedDate: date,
	}

	feeTxWrapper, _ := json.Marshal(&feeTx)
	_, err = h.db.LPush(strings.ToLower(User)+":"+string(ETH), feeTxWrapper).Result()
	if err != nil {
		log.Println(err.Error())
		return
	}

	outGameSlot := new(big.Float).Quo(new(big.Float).SetInt(rawToken), TokenBase)

	tokenUpdate := outGameSlot.String()
	tx := Transaction{
		Type:        Gain,
		Amount:      tokenUpdate,
		CreatedDate: date,
	}
	txWrapper, _ := json.Marshal(tx)
	_, err = h.db.LPush(strings.ToLower(User)+":"+string(Slot), txWrapper).Result()
	if err != nil {
		log.Println(err.Error())
		return
	}

	rawTokenTotal, err := h.c.PbcERC20.GetToken(UserAddr)
	if err != nil {
		log.Println(err.Error())
		return
	}
	tokenTotal := new(big.Float).Quo(new(big.Float).SetInt(rawTokenTotal), TokenBase).String()

	resultRes := &TokenCountChangeRes{
		Gcuid:       NotifyTokenChange,
		State:       Increase,
		WalletName:  TokenWalletName,
		WalletId:    User,
		TokenUpdate: tokenUpdate,
		TokenTotal:  tokenTotal,
		Act:         TokenChange,
		CreatedDate: date,
	}

	h.wrapperAndSend(NotifyTokenChange, resultRes)
}