package appClient

const (
	Signin = 888
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
)

const (
	LoginCode = "6469D2959C1AA0D2AA9A6ACDF35BEA32524B5142684948576C4B31586D317475695364587A425552334A34"
	Guuid = "gu9"
	User = "0x20A40B83a495DD2fbbE33E0b6ad119B09F09151f"
)

const (
	ErrorCode = 400
	FailedStatus = "failed"
	OkStatus = "ok"
)

const (
	InvalidRequestInfo = "Invalid request Information"
)


type PostRes struct {
	Status string `json:"status"`
	Gcuid int `json:"gcuid"`
}

type MachineState struct {
	State int `json:"state"`
	MachineId string `json:"machine_id"`
}

type Transaction struct {
	Type int `json:"type"`
	Amount string `json:"amount"`
	CreatedDate string `json:"created_date"`
}

type Wallet struct {
	Name string `json:"name"`
	Amount string `json:"amount"`
	Id string `json:"id"`
}

type SigninReq struct {
	Gcuid int `json:"gcuid"`
	Act string `json:"act"`
	LoginCode string `json:"login_code"`
}

type SigninRes struct {
	Status string `json:"status"`
	Gcuid int `json:"gcuid"`
	Guuid string `json:"guuid"`
}

type SignoutReq struct {
	Gcuid int `json:"gcuid"`
	Act string `json:"act"`
	Guuid string `json:"guuid"`
}

type SignoutRes struct {
	*PostRes
}

type GetWalletsAndMachineReq struct {
	Gcuid int `json:"gcuid"`
	Act string `json:"act"`
}

type GetWalletsAndMachineRes struct {
	Gcuid int `json:"gcuid"`
	*MachineState
	WalletsCount int `json:"wallets_count"`
	Wallets []*Wallet `json:"wallets"`
}

type GetWalletReq struct {
	Gcuid int `json:"gcuid"`
	Act string `json:"act"`
}

type GetWalletRes struct {
	Gcuid int `json:"gcuid"`
	Count int `json:"count"`
	Wallets []*Wallet `json:"wallets"`
}

type GetTransactionsReq struct {
	Gcuid int `json:"gcuid"`
	Act string `json:"act"`
	After int `json:"after"`
	Amount int `json:"amount"`
	From int `json:"from"`
}

type GetTransactionRes struct {
	Gcuid int `json:"gcuid"`
	Count int `json:"count"`
	After int `json:"after"`
	From int `json:"from"`
	Transactions []*Transaction `json:"transactions"`
}

type GetExchangeRateReq struct {
	Gcuid int `json:"gcuid"`
	Act string `json:"act"`
	ExchangeType int `json:"exchange_type"`
}

type GetExchangeRateRes struct {
	Gcuid int `json:"gcuid"`
	ExchangeType int `json:"exchange_type"`
	Rate string `json:"rate"`
}

type PostExchangeReq struct {
	Gcuid int `json:"gcuid"`
	Act string `json:"act"`
	ExchangeType int `json:"exchange_type"`
	Amount string `json:"amount"`
}

type PostExchangeRes struct {
	*PostRes
}

type PostQRCodeReq struct {
	Gcuid int `json:"gcuid"`
	Act string `json:"act"`
	QRCode string `json:"qrcode"`
}

type PostQRCodeRes struct {
	*PostRes
}

type MachineStatusChangeRes struct {
	Gcuid int `json:"gcuid"`
	State int `json:"state"`
	MachineId string `json:"machine_id"`
	Act string `json:"act"`
	CreatedDate string `json:"created_date"`
}

type TokenCountChangeRes struct {
	Gcuid int `json:"gcuid"`
	State int `json:"state"`
	WalletName string `json:"wallet_name"`
	WalletId string `json:"wallet_id"`
	TokenUpdate string `json:"token_update"`
	TokenTotal string `json:"token_total"`
	Act string `json:"act"`
	CreatedDate string `json:"created_date"`
}

type PostMachineLogoutReq struct {
	Gcuid int `json:"gcuid"`
	Act string `json:"act"`
	MachineId string `json:"machine_id"`
}

type PostMachineLogoutRes struct {
	*PostRes
}

type ExchangeResultRes struct{
	Gcuid int `json:"gcuid"`
	Status string `json:"status"`
	ExchangeType int `json:"exchange_type"`
	Amount string `json:"amount"`
	CreatedDate string `json:"created_date"`
}

type PostTokenUseOrRewardReq struct {
	Gcuid int `json:"gcuid"`
	Type int `json:"string"`
	Amount string `json:"string"`
	Act string `json:"string"`
}

type PostTOkenUserOrRewardRes struct {
	*PostRes
}

type Error struct {
	Status string `json:"status"`
	Code int `json:"code"`
	Gcuid int `json:"gcuid"`
	Reason string `json:"reason"`
}

type Disconnect struct {
	Status string `json:"status"`
	Code int `json:"code"`
	Reason string `json:"reason"`
}


