package bitstamp

import "encoding/json"

// TickerResponse used to map results of GetTicker, GetTickerHourly methods
type GetTickerResponse struct {
	High      string `json:"high"`
	Last      string `json:"last"`
	Timestamp string `json:"timestamp"`
	Bid       string `json:"bid"`
	Vwap      string `json:"vwap"`
	Volume    string `json:"volume"`
	Low       string `json:"low"`
	Ask       string `json:"ask"`
	Open      string `json:"open"`
}

// OrderBookResponse used to map results of GetOrderBook methods
type GetOrderBookResponse struct {
	Timestamp      string     `json:"timestamp"`
	Microtimestamp string     `json:"microtimestamp"`
	Bids           [][]string `json:"bids"`
	Asks           [][]string `json:"asks"`
}

// GetTransactionsResponse used by GetTransactions to map response
type GetTransactionResponse struct {
	TID    string `json:"tid"`
	Type   string `json:"type"`
	Amount string `json:"amount"`
	Price  string `json:"price"`
	Date   string `json:"date"`
}

// GetTradingPairInfoResult used to map result of GetTradingPairsInfo method
type GetTradingPairInfoResult struct {
	Trading                string `json:"trading"`
	BaseDecimals           int    `json:"base_decimals"`
	URLSymbol              string `json:"url_symbol"`
	Name                   string `json:"name"`
	InstantAndMarketOrders string `json:"instant_and_market_orders"`
	MinimumOrder           string `json:"minimum_order"`
	CounterDecimals        int    `json:"counter_decimals"`
	Description            string `json:"description"`
}

// GetOHLCDataResponse used to map result of GetOHLCData method
type GetOHLCDataResponse struct {
	Data struct {
		Pair string `json:"pair"`
		Ohlc []struct {
			High      string `json:"high"`
			Timestamp string `json:"timestamp"`
			Volume    string `json:"volume"`
			Low       string `json:"low"`
			Close     string `json:"close"`
			Open      string `json:"open"`
		} `json:"ohlc"`
	} `json:"data"`
}

// GetEURUSDConversionRate used to map result of GetEURUSDConversionRate method
type GetEURUSDConversionRateResult struct {
	Sell string `json:"sell"`
	Buy  string `json:"buy"`
}

// GetAccountBalancesResponse used to map result of GetAccountBalances method
type GetAccountBalancesResponse struct {
	AaveAvailable      string `json:"aave_available,omitempty"`
	AaveBalance        string `json:"aave_balance,omitempty"`
	AaveReserved       string `json:"aave_reserved,omitempty"`
	AaveWithdrawalFee  string `json:"aave_withdrawal_fee,omitempty"`
	AavebtcFee         string `json:"aavebtc_fee,omitempty"`
	AaveeurFee         string `json:"aaveeur_fee,omitempty"`
	AaveusdFee         string `json:"aaveusd_fee,omitempty"`
	AlgoAvailable      string `json:"algo_available,omitempty"`
	AlgoBalance        string `json:"algo_balance,omitempty"`
	AlgoReserved       string `json:"algo_reserved,omitempty"`
	AlgoWithdrawalFee  string `json:"algo_withdrawal_fee,omitempty"`
	AlgobtcFee         string `json:"algobtc_fee,omitempty"`
	AlgoeurFee         string `json:"algoeur_fee,omitempty"`
	AlgousdFee         string `json:"algousd_fee,omitempty"`
	AlphaAvailable     string `json:"alpha_available,omitempty"`
	AlphaBalance       string `json:"alpha_balance,omitempty"`
	AlphaReserved      string `json:"alpha_reserved,omitempty"`
	AlphaWithdrawalFee string `json:"alpha_withdrawal_fee,omitempty"`
	AlphaeurFee        string `json:"alphaeur_fee,omitempty"`
	AlphausdFee        string `json:"alphausd_fee,omitempty"`
	AudioAvailable     string `json:"audio_available,omitempty"`
	AudioBalance       string `json:"audio_balance,omitempty"`
	AudioReserved      string `json:"audio_reserved,omitempty"`
	AudioWithdrawalFee string `json:"audio_withdrawal_fee,omitempty"`
	AudiobtcFee        string `json:"audiobtc_fee,omitempty"`
	AudioeurFee        string `json:"audioeur_fee,omitempty"`
	AudiousdFee        string `json:"audiousd_fee,omitempty"`
	AxsAvailable       string `json:"axs_available,omitempty"`
	AxsBalance         string `json:"axs_balance,omitempty"`
	AxsReserved        string `json:"axs_reserved,omitempty"`
	AxsWithdrawalFee   string `json:"axs_withdrawal_fee,omitempty"`
	AxseurFee          string `json:"axseur_fee,omitempty"`
	AxsusdFee          string `json:"axsusd_fee,omitempty"`
	BatAvailable       string `json:"bat_available,omitempty"`
	BatBalance         string `json:"bat_balance,omitempty"`
	BatReserved        string `json:"bat_reserved,omitempty"`
	BatWithdrawalFee   string `json:"bat_withdrawal_fee,omitempty"`
	BatbtcFee          string `json:"batbtc_fee,omitempty"`
	BateurFee          string `json:"bateur_fee,omitempty"`
	BatusdFee          string `json:"batusd_fee,omitempty"`
	BchAvailable       string `json:"bch_available,omitempty"`
	BchBalance         string `json:"bch_balance,omitempty"`
	BchReserved        string `json:"bch_reserved,omitempty"`
	BchWithdrawalFee   string `json:"bch_withdrawal_fee,omitempty"`
	BchbtcFee          string `json:"bchbtc_fee,omitempty"`
	BcheurFee          string `json:"bcheur_fee,omitempty"`
	BchgbpFee          string `json:"bchgbp_fee,omitempty"`
	BchusdFee          string `json:"bchusd_fee,omitempty"`
	BtcAvailable       string `json:"btc_available,omitempty"`
	BtcBalance         string `json:"btc_balance,omitempty"`
	BtcReserved        string `json:"btc_reserved,omitempty"`
	BtcWithdrawalFee   string `json:"btc_withdrawal_fee,omitempty"`
	BtceurFee          string `json:"btceur_fee,omitempty"`
	BtcgbpFee          string `json:"btcgbp_fee,omitempty"`
	BtcpaxFee          string `json:"btcpax_fee,omitempty"`
	BtcusdFee          string `json:"btcusd_fee,omitempty"`
	BtcusdcFee         string `json:"btcusdc_fee,omitempty"`
	BtcusdtFee         string `json:"btcusdt_fee,omitempty"`
	CelAvailable       string `json:"cel_available,omitempty"`
	CelBalance         string `json:"cel_balance,omitempty"`
	CelReserved        string `json:"cel_reserved,omitempty"`
	CelWithdrawalFee   string `json:"cel_withdrawal_fee,omitempty"`
	CeleurFee          string `json:"celeur_fee,omitempty"`
	CelusdFee          string `json:"celusd_fee,omitempty"`
	ChzAvailable       string `json:"chz_available,omitempty"`
	ChzBalance         string `json:"chz_balance,omitempty"`
	ChzReserved        string `json:"chz_reserved,omitempty"`
	ChzWithdrawalFee   string `json:"chz_withdrawal_fee,omitempty"`
	ChzeurFee          string `json:"chzeur_fee,omitempty"`
	ChzusdFee          string `json:"chzusd_fee,omitempty"`
	CompAvailable      string `json:"comp_available,omitempty"`
	CompBalance        string `json:"comp_balance,omitempty"`
	CompReserved       string `json:"comp_reserved,omitempty"`
	CompWithdrawalFee  string `json:"comp_withdrawal_fee,omitempty"`
	CompbtcFee         string `json:"compbtc_fee,omitempty"`
	CompeurFee         string `json:"compeur_fee,omitempty"`
	CompusdFee         string `json:"compusd_fee,omitempty"`
	CrvAvailable       string `json:"crv_available,omitempty"`
	CrvBalance         string `json:"crv_balance,omitempty"`
	CrvReserved        string `json:"crv_reserved,omitempty"`
	CrvWithdrawalFee   string `json:"crv_withdrawal_fee,omitempty"`
	CrvbtcFee          string `json:"crvbtc_fee,omitempty"`
	CrveurFee          string `json:"crveur_fee,omitempty"`
	CrvusdFee          string `json:"crvusd_fee,omitempty"`
	DaiAvailable       string `json:"dai_available,omitempty"`
	DaiBalance         string `json:"dai_balance,omitempty"`
	DaiReserved        string `json:"dai_reserved,omitempty"`
	DaiWithdrawalFee   string `json:"dai_withdrawal_fee,omitempty"`
	DaiusdFee          string `json:"daiusd_fee,omitempty"`
	EnjAvailable       string `json:"enj_available,omitempty"`
	EnjBalance         string `json:"enj_balance,omitempty"`
	EnjReserved        string `json:"enj_reserved,omitempty"`
	EnjWithdrawalFee   string `json:"enj_withdrawal_fee,omitempty"`
	EnjeurFee          string `json:"enjeur_fee,omitempty"`
	EnjusdFee          string `json:"enjusd_fee,omitempty"`
	Eth2Available      string `json:"eth2_available,omitempty"`
	Eth2Balance        string `json:"eth2_balance,omitempty"`
	Eth2Reserved       string `json:"eth2_reserved,omitempty"`
	Eth2EthFee         string `json:"eth2eth_fee,omitempty"`
	Eth2RAvailable     string `json:"eth2r_available,omitempty"`
	Eth2RBalance       string `json:"eth2r_balance,omitempty"`
	Eth2RReserved      string `json:"eth2r_reserved,omitempty"`
	EthAvailable       string `json:"eth_available,omitempty"`
	EthBalance         string `json:"eth_balance,omitempty"`
	EthReserved        string `json:"eth_reserved,omitempty"`
	EthWithdrawalFee   string `json:"eth_withdrawal_fee,omitempty"`
	EthbtcFee          string `json:"ethbtc_fee,omitempty"`
	EtheurFee          string `json:"etheur_fee,omitempty"`
	EthgbpFee          string `json:"ethgbp_fee,omitempty"`
	EthpaxFee          string `json:"ethpax_fee,omitempty"`
	EthusdFee          string `json:"ethusd_fee,omitempty"`
	EthusdcFee         string `json:"ethusdc_fee,omitempty"`
	EthusdtFee         string `json:"ethusdt_fee,omitempty"`
	EurAvailable       string `json:"eur_available,omitempty"`
	EurBalance         string `json:"eur_balance,omitempty"`
	EurReserved        string `json:"eur_reserved,omitempty"`
	EurtAvailable      string `json:"eurt_available,omitempty"`
	EurtBalance        string `json:"eurt_balance,omitempty"`
	EurtReserved       string `json:"eurt_reserved,omitempty"`
	EurtWithdrawalFee  string `json:"eurt_withdrawal_fee,omitempty"`
	EurteurFee         string `json:"eurteur_fee,omitempty"`
	EurtusdFee         string `json:"eurtusd_fee,omitempty"`
	EurusdFee          string `json:"eurusd_fee,omitempty"`
	FetAvailable       string `json:"fet_available,omitempty"`
	FetBalance         string `json:"fet_balance,omitempty"`
	FetReserved        string `json:"fet_reserved,omitempty"`
	FetWithdrawalFee   string `json:"fet_withdrawal_fee,omitempty"`
	FeteurFee          string `json:"feteur_fee,omitempty"`
	FetusdFee          string `json:"fetusd_fee,omitempty"`
	FttAvailable       string `json:"ftt_available,omitempty"`
	FttBalance         string `json:"ftt_balance,omitempty"`
	FttReserved        string `json:"ftt_reserved,omitempty"`
	FttWithdrawalFee   string `json:"ftt_withdrawal_fee,omitempty"`
	FtteurFee          string `json:"ftteur_fee,omitempty"`
	FttusdFee          string `json:"fttusd_fee,omitempty"`
	GbpAvailable       string `json:"gbp_available,omitempty"`
	GbpBalance         string `json:"gbp_balance,omitempty"`
	GbpReserved        string `json:"gbp_reserved,omitempty"`
	GbpeurFee          string `json:"gbpeur_fee,omitempty"`
	GbpusdFee          string `json:"gbpusd_fee,omitempty"`
	GrtAvailable       string `json:"grt_available,omitempty"`
	GrtBalance         string `json:"grt_balance,omitempty"`
	GrtReserved        string `json:"grt_reserved,omitempty"`
	GrtWithdrawalFee   string `json:"grt_withdrawal_fee,omitempty"`
	GrteurFee          string `json:"grteur_fee,omitempty"`
	GrtusdFee          string `json:"grtusd_fee,omitempty"`
	GusdAvailable      string `json:"gusd_available,omitempty"`
	GusdBalance        string `json:"gusd_balance,omitempty"`
	GusdReserved       string `json:"gusd_reserved,omitempty"`
	GusdWithdrawalFee  string `json:"gusd_withdrawal_fee,omitempty"`
	GusdusdFee         string `json:"gusdusd_fee,omitempty"`
	HbarAvailable      string `json:"hbar_available,omitempty"`
	HbarBalance        string `json:"hbar_balance,omitempty"`
	HbarReserved       string `json:"hbar_reserved,omitempty"`
	HbarWithdrawalFee  string `json:"hbar_withdrawal_fee,omitempty"`
	HbareurFee         string `json:"hbareur_fee,omitempty"`
	HbarusdFee         string `json:"hbarusd_fee,omitempty"`
	KncAvailable       string `json:"knc_available,omitempty"`
	KncBalance         string `json:"knc_balance,omitempty"`
	KncReserved        string `json:"knc_reserved,omitempty"`
	KncWithdrawalFee   string `json:"knc_withdrawal_fee,omitempty"`
	KncbtcFee          string `json:"kncbtc_fee,omitempty"`
	KnceurFee          string `json:"knceur_fee,omitempty"`
	KncusdFee          string `json:"kncusd_fee,omitempty"`
	LinkAvailable      string `json:"link_available,omitempty"`
	LinkBalance        string `json:"link_balance,omitempty"`
	LinkReserved       string `json:"link_reserved,omitempty"`
	LinkWithdrawalFee  string `json:"link_withdrawal_fee,omitempty"`
	LinkbtcFee         string `json:"linkbtc_fee,omitempty"`
	LinkethFee         string `json:"linketh_fee,omitempty"`
	LinkeurFee         string `json:"linkeur_fee,omitempty"`
	LinkgbpFee         string `json:"linkgbp_fee,omitempty"`
	LinkusdFee         string `json:"linkusd_fee,omitempty"`
	LtcAvailable       string `json:"ltc_available,omitempty"`
	LtcBalance         string `json:"ltc_balance,omitempty"`
	LtcReserved        string `json:"ltc_reserved,omitempty"`
	LtcWithdrawalFee   string `json:"ltc_withdrawal_fee,omitempty"`
	LtcbtcFee          string `json:"ltcbtc_fee,omitempty"`
	LtceurFee          string `json:"ltceur_fee,omitempty"`
	LtcgbpFee          string `json:"ltcgbp_fee,omitempty"`
	LtcusdFee          string `json:"ltcusd_fee,omitempty"`
	MaticAvailable     string `json:"matic_available,omitempty"`
	MaticBalance       string `json:"matic_balance,omitempty"`
	MaticReserved      string `json:"matic_reserved,omitempty"`
	MaticWithdrawalFee string `json:"matic_withdrawal_fee,omitempty"`
	MaticeurFee        string `json:"maticeur_fee,omitempty"`
	MaticusdFee        string `json:"maticusd_fee,omitempty"`
	MkrAvailable       string `json:"mkr_available,omitempty"`
	MkrBalance         string `json:"mkr_balance,omitempty"`
	MkrReserved        string `json:"mkr_reserved,omitempty"`
	MkrWithdrawalFee   string `json:"mkr_withdrawal_fee,omitempty"`
	MkrbtcFee          string `json:"mkrbtc_fee,omitempty"`
	MkreurFee          string `json:"mkreur_fee,omitempty"`
	MkrusdFee          string `json:"mkrusd_fee,omitempty"`
	OmgAvailable       string `json:"omg_available,omitempty"`
	OmgBalance         string `json:"omg_balance,omitempty"`
	OmgReserved        string `json:"omg_reserved,omitempty"`
	OmgWithdrawalFee   string `json:"omg_withdrawal_fee,omitempty"`
	OmgbtcFee          string `json:"omgbtc_fee,omitempty"`
	OmgeurFee          string `json:"omgeur_fee,omitempty"`
	OmggbpFee          string `json:"omggbp_fee,omitempty"`
	OmgusdFee          string `json:"omgusd_fee,omitempty"`
	PaxAvailable       string `json:"pax_available,omitempty"`
	PaxBalance         string `json:"pax_balance,omitempty"`
	PaxReserved        string `json:"pax_reserved,omitempty"`
	PaxWithdrawalFee   string `json:"pax_withdrawal_fee,omitempty"`
	PaxeurFee          string `json:"paxeur_fee,omitempty"`
	PaxgbpFee          string `json:"paxgbp_fee,omitempty"`
	PaxusdFee          string `json:"paxusd_fee,omitempty"`
	RgtAvailable       string `json:"rgt_available,omitempty"`
	RgtBalance         string `json:"rgt_balance,omitempty"`
	RgtReserved        string `json:"rgt_reserved,omitempty"`
	RgtWithdrawalFee   string `json:"rgt_withdrawal_fee,omitempty"`
	RgteurFee          string `json:"rgteur_fee,omitempty"`
	RgtusdFee          string `json:"rgtusd_fee,omitempty"`
	SandAvailable      string `json:"sand_available,omitempty"`
	SandBalance        string `json:"sand_balance,omitempty"`
	SandReserved       string `json:"sand_reserved,omitempty"`
	SandWithdrawalFee  string `json:"sand_withdrawal_fee,omitempty"`
	SandeurFee         string `json:"sandeur_fee,omitempty"`
	SandusdFee         string `json:"sandusd_fee,omitempty"`
	SklAvailable       string `json:"skl_available,omitempty"`
	SklBalance         string `json:"skl_balance,omitempty"`
	SklReserved        string `json:"skl_reserved,omitempty"`
	SklWithdrawalFee   string `json:"skl_withdrawal_fee,omitempty"`
	SkleurFee          string `json:"skleur_fee,omitempty"`
	SklusdFee          string `json:"sklusd_fee,omitempty"`
	SnxAvailable       string `json:"snx_available,omitempty"`
	SnxBalance         string `json:"snx_balance,omitempty"`
	SnxReserved        string `json:"snx_reserved,omitempty"`
	SnxWithdrawalFee   string `json:"snx_withdrawal_fee,omitempty"`
	SnxbtcFee          string `json:"snxbtc_fee,omitempty"`
	SnxeurFee          string `json:"snxeur_fee,omitempty"`
	SnxusdFee          string `json:"snxusd_fee,omitempty"`
	StorjAvailable     string `json:"storj_available,omitempty"`
	StorjBalance       string `json:"storj_balance,omitempty"`
	StorjReserved      string `json:"storj_reserved,omitempty"`
	StorjWithdrawalFee string `json:"storj_withdrawal_fee,omitempty"`
	StorjeurFee        string `json:"storjeur_fee,omitempty"`
	StorjusdFee        string `json:"storjusd_fee,omitempty"`
	SushiAvailable     string `json:"sushi_available,omitempty"`
	SushiBalance       string `json:"sushi_balance,omitempty"`
	SushiReserved      string `json:"sushi_reserved,omitempty"`
	SushiWithdrawalFee string `json:"sushi_withdrawal_fee,omitempty"`
	SushieurFee        string `json:"sushieur_fee,omitempty"`
	SushiusdFee        string `json:"sushiusd_fee,omitempty"`
	SxpAvailable       string `json:"sxp_available,omitempty"`
	SxpBalance         string `json:"sxp_balance,omitempty"`
	SxpReserved        string `json:"sxp_reserved,omitempty"`
	SxpWithdrawalFee   string `json:"sxp_withdrawal_fee,omitempty"`
	SxpeurFee          string `json:"sxpeur_fee,omitempty"`
	SxpusdFee          string `json:"sxpusd_fee,omitempty"`
	UmaAvailable       string `json:"uma_available,omitempty"`
	UmaBalance         string `json:"uma_balance,omitempty"`
	UmaReserved        string `json:"uma_reserved,omitempty"`
	UmaWithdrawalFee   string `json:"uma_withdrawal_fee,omitempty"`
	UmabtcFee          string `json:"umabtc_fee,omitempty"`
	UmaeurFee          string `json:"umaeur_fee,omitempty"`
	UmausdFee          string `json:"umausd_fee,omitempty"`
	UniAvailable       string `json:"uni_available,omitempty"`
	UniBalance         string `json:"uni_balance,omitempty"`
	UniReserved        string `json:"uni_reserved,omitempty"`
	UniWithdrawalFee   string `json:"uni_withdrawal_fee,omitempty"`
	UnibtcFee          string `json:"unibtc_fee,omitempty"`
	UnieurFee          string `json:"unieur_fee,omitempty"`
	UniusdFee          string `json:"uniusd_fee,omitempty"`
	UsdAvailable       string `json:"usd_available,omitempty"`
	UsdBalance         string `json:"usd_balance,omitempty"`
	UsdReserved        string `json:"usd_reserved,omitempty"`
	UsdcAvailable      string `json:"usdc_available,omitempty"`
	UsdcBalance        string `json:"usdc_balance,omitempty"`
	UsdcReserved       string `json:"usdc_reserved,omitempty"`
	UsdcWithdrawalFee  string `json:"usdc_withdrawal_fee,omitempty"`
	UsdceurFee         string `json:"usdceur_fee,omitempty"`
	UsdcusdFee         string `json:"usdcusd_fee,omitempty"`
	UsdcusdtFee        string `json:"usdcusdt_fee,omitempty"`
	UsdtAvailable      string `json:"usdt_available,omitempty"`
	UsdtBalance        string `json:"usdt_balance,omitempty"`
	UsdtReserved       string `json:"usdt_reserved,omitempty"`
	UsdtWithdrawalFee  string `json:"usdt_withdrawal_fee,omitempty"`
	UsdteurFee         string `json:"usdteur_fee,omitempty"`
	UsdtusdFee         string `json:"usdtusd_fee,omitempty"`
	XlmAvailable       string `json:"xlm_available,omitempty"`
	XlmBalance         string `json:"xlm_balance,omitempty"`
	XlmReserved        string `json:"xlm_reserved,omitempty"`
	XlmWithdrawalFee   string `json:"xlm_withdrawal_fee,omitempty"`
	XlmbtcFee          string `json:"xlmbtc_fee,omitempty"`
	XlmeurFee          string `json:"xlmeur_fee,omitempty"`
	XlmgbpFee          string `json:"xlmgbp_fee,omitempty"`
	XlmusdFee          string `json:"xlmusd_fee,omitempty"`
	XrpAvailable       string `json:"xrp_available,omitempty"`
	XrpBalance         string `json:"xrp_balance,omitempty"`
	XrpReserved        string `json:"xrp_reserved,omitempty"`
	XrpWithdrawalFee   string `json:"xrp_withdrawal_fee,omitempty"`
	XrpbtcFee          string `json:"xrpbtc_fee,omitempty"`
	XrpeurFee          string `json:"xrpeur_fee,omitempty"`
	XrpgbpFee          string `json:"xrpgbp_fee,omitempty"`
	XrppaxFee          string `json:"xrppax_fee,omitempty"`
	XrpusdFee          string `json:"xrpusd_fee,omitempty"`
	XrpusdtFee         string `json:"xrpusdt_fee,omitempty"`
	YfiAvailable       string `json:"yfi_available,omitempty"`
	YfiBalance         string `json:"yfi_balance,omitempty"`
	YfiReserved        string `json:"yfi_reserved,omitempty"`
	YfiWithdrawalFee   string `json:"yfi_withdrawal_fee,omitempty"`
	YfibtcFee          string `json:"yfibtc_fee,omitempty"`
	YfieurFee          string `json:"yfieur_fee,omitempty"`
	YfiusdFee          string `json:"yfiusd_fee,omitempty"`
	ZrxAvailable       string `json:"zrx_available,omitempty"`
	ZrxBalance         string `json:"zrx_balance,omitempty"`
	ZrxReserved        string `json:"zrx_reserved,omitempty"`
	ZrxWithdrawalFee   string `json:"zrx_withdrawal_fee,omitempty"`
	ZrxbtcFee          string `json:"zrxbtc_fee,omitempty"`
	ZrxeurFee          string `json:"zrxeur_fee,omitempty"`
	ZrxusdFee          string `json:"zrxusd_fee,omitempty"`
}

// GetUserTransactionResponse used to map response of GetUserTransactions method
type GetUserTransactionResponse struct {
	ID       int     `json:"id"`
	OrderID  int     `json:"order_id"`
	Type     string  `json:"type"`
	Fee      string  `json:"fee"`
	Datetime string  `json:"datetime,omitempty"`
	BtcUsd   string  `json:"btc_usd,omitempty"`
	Usd      float64 `json:"usd,omitempty"`
	Btc      float64 `json:"btc,omitempty"`
	Eur      string  `json:"eur,omitempty"`
	XrpEur   float64 `json:"xrp_eur,omitempty"`
	XrpUsd   float64 `json:"xrp_usd,omitempty"`
	Xrp      string  `json:"xrp,omitempty"`
	Eth      string  `json:"eth,omitempty"`
	EthEur   float64 `json:"eth_eur,omitempty"`
	EthUsd   float64 `json:"eth_usd,omitempty"`
	Xlm      string  `json:"xlm,omitempty"`
	XlmEur   float64 `json:"xlm_eur,omitempty"`
	XlmUsd   float64 `json:"xlm_usd,omitempty"`
	Omg      string  `json:"omg,omitempty"`
	OmgEur   float64 `json:"omg_eur,omitempty"`
	OmgUsd   float64 `json:"omg_usd,omitempty"`
	ZrxEur   float64 `json:"zrx_eur,omitempty"`
	ZrxUsd   float64 `json:"zrx_usd,omitempty"`
	Zrx      string  `json:"zrx,omitempty"`
}

// GetCryptoTransactionsResponse used to map response of GetCryptoTransactions method
type GetCryptoTransactionsResponse struct {
	Deposits []struct {
		Currency           string `json:"currency,"`
		DestinationAddress string `json:"destination_address,"`
		TXID               string `json:"txid,"`
		Ammount            string `json:"ammount,"`
		Datetime           string `json:"datetime,"`
	} `json:"deposits"`
	Withdrawals []struct {
		Currency           string `json:"currency,"`
		DestinationAddress string `json:"destination_address,"`
		TXID               string `json:"txid,"`
		Ammount            string `json:"ammount,"`
		Datetime           string `json:"datetime,"`
	} `json:"withdrawals"`
}

// GetOpenOrderResponse used to map response of GetOpenOrders method
type GetOpenOrderResponse struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	Price        string `json:"price"`
	CurrencyPair string `json:"currency_pair"`
	Datetime     string `json:"datetime"`
	Amount       string `json:"amount"`
}

// GetOpenOrderResponse used to map response of GetOpenOrders method
type GetOrderStatusResponse struct {
	ID           int64  `json:"id"`
	Status       string `json:"status"`
	Transactions []struct {
		TID      string `json:"tid"`
		USD      string `json:"usd"`
		Price    string `json:"price"`
		Fee      string `json:"fee"`
		Btc      string `json:"btc"`
		Datetime string `json:"datetime"`
		// (0 - deposit; 1 - withdrawal; 2 - market trade).
		Type string `json:"type"`
	} `json:"transactions"`
	AmountRemaining string `json:"amount_remaining"`
	// Client order id. (Only returned if order was placed with client order id parameter.)
	ClientOrderID *string `json:"client_order_id"`
}

// CancelOrderResponse used to map response of CancelOrder method
type CancelOrderResponse struct {
	ID     int64   `json:"id"`
	Type   int     `json:"type"`
	Amount float64 `json:"amount"`
	Price  int     `json:"price"`
}

// CancelAllOrdersResponse used to map response of CancelAllOrders method
type CancelAllOrdersResponse struct {
	Canceled []struct {
		ID           int64   `json:"id"`
		CurrencyPair string  `json:"currency_pair"`
		Type         int     `json:"type"`
		Amount       float64 `json:"amount"`
		Price        int     `json:"price"`
	} `json:"canceled"`
	Success bool `json:"success"`
}

// CreateSellLimitOrderResponse used to map response of CreateSellLimitOrder method
type CreateSellLimitOrderResponse struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Price    string `json:"price"`
	Amount   string `json:"amount"`
	Datetime string `json:"datetime"`
}

// GetWebsocketTokenResponse use to map response of GetWebsocketToken method
type GetWebsocketTokenResponse struct {
	Token        string `json:"token"`
	ValidSeconds string `json:"valid_sec"`
}

// GenericErrorResponse errors are not using a unified format trying to map them all kind of error responses at a generic object
type GenericErrorResponse struct {
	Error  string          `json:"error"`
	Errors json.RawMessage `json:"errors"`
	Reason json.RawMessage `json:"reason"`
	Status string          `json:"status"`
	Code   string          `json:"code"`
}
