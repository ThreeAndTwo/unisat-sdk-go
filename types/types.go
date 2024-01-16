package types

type Unisat struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Brc20StatusRes struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data Brc20Status `json:"data"`
}

type Brc20Status struct {
	Height int               `json:"height"`
	Total  int               `json:"total"`
	Start  int               `json:"start"`
	Detail []Brc20StatusInfo `json:"detail"`
}

type Brc20StatusInfo struct {
	Ticker                 string `json:"ticker"`
	HoldersCount           int    `json:"holdersCount"`
	HistoryCount           int    `json:"historyCount"`
	InscriptionNumber      int    `json:"inscriptionNumber"`
	InscriptionId          string `json:"inscriptionId"`
	Max                    string `json:"max"`
	Limit                  string `json:"limit"`
	Minted                 string `json:"minted"`
	TotalMinted            string `json:"totalMinted"`
	ConfirmedMinted        string `json:"confirmedMinted"`
	ConfirmedMinted1H      string `json:"confirmedMinted1h"`
	ConfirmedMinted24H     string `json:"confirmedMinted24h"`
	MintTimes              int    `json:"mintTimes"`
	Decimal                int    `json:"decimal"`
	Creator                string `json:"creator"`
	Txid                   string `json:"txid"`
	DeployHeight           int    `json:"deployHeight"`
	DeployBlocktime        int    `json:"deployBlocktime"`
	CompleteHeight         int    `json:"completeHeight"`
	CompleteBlocktime      int    `json:"completeBlocktime"`
	InscriptionNumberStart int    `json:"inscriptionNumberStart"`
	InscriptionNumberEnd   int    `json:"inscriptionNumberEnd"`
}

type Brc20TickerInfoRes struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data Brc20TickerInfo `json:"data"`
}

type Brc20TickerInfo struct {
	Ticker                 string `json:"ticker"`
	HoldersCount           int    `json:"holdersCount"`
	HistoryCount           int    `json:"historyCount"`
	InscriptionNumber      int    `json:"inscriptionNumber"`
	InscriptionId          string `json:"inscriptionId"`
	Max                    string `json:"max"`
	Limit                  string `json:"limit"`
	Minted                 string `json:"minted"`
	TotalMinted            string `json:"totalMinted"`
	ConfirmedMinted        string `json:"confirmedMinted"`
	ConfirmedMinted1H      string `json:"confirmedMinted1h"`
	ConfirmedMinted24H     string `json:"confirmedMinted24h"`
	MintTimes              int    `json:"mintTimes"`
	Decimal                int    `json:"decimal"`
	Creator                string `json:"creator"`
	Txid                   string `json:"txid"`
	DeployHeight           int    `json:"deployHeight"`
	DeployBlocktime        int    `json:"deployBlocktime"`
	CompleteHeight         int    `json:"completeHeight"`
	CompleteBlocktime      int    `json:"completeBlocktime"`
	InscriptionNumberStart int    `json:"inscriptionNumberStart"`
	InscriptionNumberEnd   int    `json:"inscriptionNumberEnd"`
}

type Brc20HoldersRes struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data Brc20Holders `json:"data"`
}

type Brc20Holders struct {
	Height int                  `json:"height"`
	Total  int                  `json:"total"`
	Start  int                  `json:"start"`
	Detail []Brc20HoldersDetail `json:"detail"`
}

type Brc20HoldersDetail struct {
	Address                string `json:"address"`
	OverallBalance         string `json:"overallBalance"`
	TransferableBalance    string `json:"transferableBalance"`
	AvailableBalance       string `json:"availableBalance"`
	AvailableBalanceSafe   string `json:"availableBalanceSafe"`
	AvailableBalanceUnSafe string `json:"availableBalanceUnSafe"`
}

type Brc20SummaryRes struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data Brc20Summary `json:"data"`
}

type Brc20Summary struct {
	Height int                  `json:"height"`
	Total  int                  `json:"total"`
	Start  int                  `json:"start"`
	Detail []Brc20SummaryDetail `json:"detail"`
}

type Brc20SummaryDetail struct {
	Ticker                 string `json:"ticker"`
	OverallBalance         string `json:"overallBalance"`
	TransferableBalance    string `json:"transferableBalance"`
	AvailableBalance       string `json:"availableBalance"`
	AvailableBalanceSafe   string `json:"availableBalanceSafe"`
	AvailableBalanceUnSafe string `json:"availableBalanceUnSafe"`
	Decimal                int    `json:"decimal"`
}
