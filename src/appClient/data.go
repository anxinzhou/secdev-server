package appClient

import (
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"time"
)


//Gcuid
const (
	Signout = 100+iota
	GetWalletsAndMachine
	GetWallets
	GetTransactions
	GetExchangeRate
	PostExchange
	PostQRCode
	NotifyMachineStatusChange
	NotifyTokenChange
	MachineLogout
	NotifyExchangeResult
	PostTokenUseOrReward
	PostGameStartOrEnd
	PostTransfer
	Signin = 888
	DisConnect = 1000
)


// Constant Value
const (
	SleepTime = 2* time.Second
	LoginCode = "6469D2959C1AA0D2AA9A6ACDF35BEA32524B5142684948576C4B31586D317475695364587A425552334A34"
	Guuid = "gu9"
	User = "0x20A40B83a495DD2fbbE33E0b6ad119B09F09151f"
	MachineId = "0"
	WalletCount = 2
	EthWalletName = "ETH"
	TokenWalletName = "SLOT"
	EtherDecimals = "1000000000000000000"
	TokenDecimals = "1000000000000000000"
)

//Error status
const (
	Logout = 0+iota
	Login
	Fail
	ClientErrorCode = 400
	ServerErrorCode = 500
	FailedStatus = "failed"
	SuccessStatus = "success"
	OkStatus = "ok"
)

//Error Info
const (
	InvalidRequestInfo = "Invalid request Information"
	WrongUser = "User id or password not correct"
	ClientFormatError = "wrong data format"
	ServerJsonError = "Can not marshal json"
	ServerChainError = "Transaction time out or revert"
	ServerDisconnect = "websocket disconenct"
)

// token type
const (
	ETH = iota+1
	Slot
)

//Exchange Type
const (
	EthToSLot = iota +1
	SlotToEth
)

//Update Type
const (
	Used = iota +1
	Reward
)


// Token count change Type
const (
	Increase = iota
	Decrease
)

// Act type
const (
	TokenChange = "notify_token_changes"
	MachineStatusChange = "notify_machine_state"
)

// Transaction type
const (
	Deposit = iota+1
	Withdraw
	Gain
	Spend
	Gas
)

// Game procedure
const (
	GameStart = iota+1
	GameEnd
)

var (
	EtherBase *big.Float
	TokenBase *big.Float
	UserAddr common.Address
)

func init(){
	var err error
	EtherBase,_,err = new(big.Float).Parse(EtherDecimals,10)
	if err!=nil{
		log.Fatalln("can not convert base")
	}
	TokenBase,_,err = new(big.Float).Parse(TokenDecimals,10)
	if err!=nil{
		log.Fatalln("can not convert base")
	}
	UserAddr = common.HexToAddress(User)
}


type TransferInfo struct {
	Name string `json:"name"`
	Amount string `json:"amount"`
}

type PostRes struct {
	Status string `json:"status"`
	Gcuid int64 `json:"gcuid"`
}

type MachineState struct {
	State int64 `json:"state"`
	MachineId string `json:"machine_id"`
}

type Transaction struct {
	Type int64 `json:"type"`
	Amount string `json:"amount"`
	CreatedDate string `json:"created_date"`
}

type Wallet struct {
	Name string `json:"name"`
	Amount string `json:"amount"`
	Id string `json:"id"`
}

type SigninReq struct {
	Gcuid int64 `json:"gcuid"`
	Act string `json:"act"`
	LoginCode string `json:"login_code"`
}

type SigninRes struct {
	Status string `json:"status"`
	Gcuid int64 `json:"gcuid"`
	Guuid string `json:"guuid"`
}

type SignoutReq struct {
	Gcuid int64 `json:"gcuid"`
	Act string `json:"act"`
	Guuid string `json:"guuid"`
}

type SignoutRes struct {
	*PostRes
}

type GetWalletsAndMachineReq struct {
	Gcuid int64 `json:"gcuid"`
	Act string `json:"act"`
}

type GetWalletsAndMachineRes struct {
	Gcuid int64 `json:"gcuid"`
	Machine *MachineState `json:"machine"`
	WalletsCount int64 `json:"wallets_count"`
	Wallets []*Wallet `json:"wallets"`
	Transfer *TransferInfo `json:"transfer"`
}

type GetWalletReq struct {
	Gcuid int64 `json:"gcuid"`
	Act string `json:"act"`
}

type GetWalletRes struct {
	Gcuid int64 `json:"gcuid"`
	Count int64 `json:"count"`
	Wallets []*Wallet `json:"wallets"`
	Transfer *TransferInfo `json:"transfer"`
}

type GetTransactionsReq struct {
	Gcuid int64 `json:"gcuid"`
	Act string `json:"act"`
	After int64 `json:"after"`
	Amount int64 `json:"amount"`
	From int64 `json:"from"`
}

type GetTransactionRes struct {
	Gcuid int64 `json:"gcuid"`
	Count int64 `json:"count"`
	After int64 `json:"after"`
	From int64 `json:"from"`
	Transactions []*Transaction `json:"transactions"`
}

type GetExchangeRateReq struct {
	Gcuid int64 `json:"gcuid"`
	Act string `json:"act"`
	ExchangeType int64 `json:"exchange_type"`
}

type GetExchangeRateRes struct {
	Gcuid int64 `json:"gcuid"`
	ExchangeType int64 `json:"exchange_type"`
	Rate string `json:"rate"`
}

type PostExchangeReq struct {
	Gcuid int64 `json:"gcuid"`
	Act string `json:"act"`
	ExchangeType int64 `json:"exchange_type"`
	Amount string `json:"amount"`
}

type PostExchangeRes struct {
	*PostRes
}

type PostQRCodeReq struct {
	Gcuid int64 `json:"gcuid"`
	Act string `json:"act"`
	QRCode string `json:"qrcode"`
}

type PostQRCodeRes struct {
	*PostRes
}

type MachineStatusChangeRes struct {
	Gcuid int64 `json:"gcuid"`
	State int64 `json:"state"`
	MachineId string `json:"machine_id"`
	Act string `json:"act"`
	CreatedDate string `json:"created_date"`
}

type TokenCountChangeRes struct {
	Gcuid int64 `json:"gcuid"`
	State int64 `json:"state"`
	WalletName string `json:"wallet_name"`
	WalletId string `json:"wallet_id"`
	TokenUpdate string `json:"token_update"`
	TokenTotal string `json:"token_total"`
	Act string `json:"act"`
	CreatedDate string `json:"created_date"`
}

type PostMachineLogoutReq struct {
	Gcuid int64 `json:"gcuid"`
	Act string `json:"act"`
	MachineId string `json:"machine_id"`
}

type PostMachineLogoutRes struct {
	*PostRes
}

type ExchangeResultRes struct{
	Gcuid int64 `json:"gcuid"`
	Status string `json:"status"`
	ExchangeType int64 `json:"exchange_type"`
	Amount string `json:"amount"`
	CreatedDate string `json:"created_date"`
}

type PostTokenUseOrRewardReq struct {
	Gcuid int64 `json:"gcuid"`
	Type int64 `json:"type"`
	Amount string `json:"amount"`
	Act string `json:"act"`
}

type PostTokenUserOrRewardRes struct {
	*PostRes
}

type PostGameStartOrEndReq struct {
	Gcuid int64 `json:"gcuid"`
	Type int64 `json:"type"`
	Amount string `json:"amount"`
	Act string `json:"act"`
}

type PostGameStartOrEndRes struct {
	*PostRes
}

type Error struct {
	Status string `json:"status"`
	Code int64 `json:"code"`
	Gcuid int64 `json:"gcuid"`
	Reason string `json:"reason"`
}

type Disconnect struct {
	Status string `json:"status"`
	Code int64 `json:"code"`
	Reason string `json:"reason"`
}

type PostEosTokenUpdateReq struct {
	Gcuid int64 `json:"gcuid"`
	Type int64 `json:"type"`
	Amount string `json:"amount"`
	Account string `json:"account"`
}