package bitstamp

//go:generate go run tools/generatepairs/generatepairs.go
//go:generate go run tools/generatechannels/generatechannels.go

// nolint:deadcode,varcheck
const (
	// Public API urls
	tickerURL               = `/api/v2/ticker/%s/`
	tickerHourlyURL         = `/api/v2/ticker_hour/%s/`
	orderBookURL            = `/api/v2/order_book/%s/`
	transactions            = `/api/v2/transactions/%s/`
	tradingPairsInfoURL     = `/api/v2/trading-pairs-info/`
	ohlcDataURL             = `/api/v2/ohlc/%s/`
	eurusdConversionRateURL = `/api/v2/eur_usd/`

	// Private API urls
	accountBalanceURL        = `/api/v2/balance/`
	userTransactionsURL      = `/api/v2/user_transactions/`
	cryptoTransactionsURL    = `/api/v2/crypto-transactions/`
	openOrdersURL            = `/api/v2/open_orders/all/`
	orderStatusURL           = `/api/v2/order_status/`
	cancelOrderURL           = `/api/v2/cancel_order/`
	cancelAllOrdersURL       = `/api/v2/cancel_all_orders/`
	buyLimitOrderURL         = `/api/v2/buy/%s/`
	buyMarketOrderURL        = `/api/v2/buy/market/%s/`
	buyInstantOrderURL       = `/api/v2/buy/instant/%s/`
	sellLimitOrderURL        = `/api/v2/sell/%s/`
	sellMarketOrderURL       = `/api/v2/sell/market/%s/`
	sellInstantOrderURL      = `/api/v2/sell/instant/%s/`
	withdrawalRequestsURL    = `/api/v2/withdrawal-requests/`
	bitcoinCashWithdrawalURL = `/api/v2/bch_withdrawal/`
	bitcoinWithdrawalURL     = `/api/v2/btc_withdrawal/`
	websocketsTokenURL       = `/api/v2/websockets_token/`
)
