package bitstamp

type Sort string

const (
	SortASC  Sort = "asc"
	SortDESC Sort = "desc"
)

// GetTransactionsRequest used by GetTransactions method to map outgoing request data
type GetTransactionsRequest struct {
	// The time interval from which we want the transactions to be returned. Possible values are minute, hour (default) or day.
	Time string `schema:"time,omitempty"`
}

// GetOHLCDataRequest used by GetOHLCData method to map result
type GetOHLCDataRequest struct {
	// Unix timestamp from when OHLC data will be started. (optional)
	Start int64 `schema:"start,omitempty"`
	// Unix timestamp to when OHLC data will be shown. (optional)
	End int64 `schema:"end,omitempty"`
	// Timeframe in seconds. Possible options are 60, 180, 300, 900, 1800, 3600, 7200, 14400, 21600, 43200, 86400, 259200
	Step int64 `schema:"step,required"`
	// Limit OHLC results (minimum: 1; maximum: 1000)
	Limit int64 `schema:"limit,required"`
}

// GetUserTransactionsRequest used by GetUserTransactions method to map outgoing request data
type GetUserTransactionsRequest struct {
	// Skip that many transactions before returning results (default: 0, maximum: 200000).
	// If you need to export older history contact support OR use combination of limit and since_id parameters
	Offset int64 `schema:"offset,required"`
	// Limit result to that many transactions (default: 100; maximum: 1000).
	Limit int64 `schema:"limit,required"`
	// Sorting by date and time: asc - ascending; desc - descending (default: desc).
	Sort Sort `schema:"sort,required"`
	// Show only transactions from unix timestamp (for max 30 days old). (optional)
	SinceTimestamp int64 `schema:"since_timestamp,omitempty"`
	// Show only transactions from specified transaction id. If since_id parameter is used, limit parameter is set to 1000. (optional)
	SinceID int64 `schema:"since_id,omitempty"`
}

// GetCryptoTransactionsRequest used by GetCryptoTransactions method to map outgoing request data
type GetCryptoTransactionsRequest struct {
	// Limit result to that many transactions (default: 100; minimum: 1; maximum: 1000).
	Limit int64 `schema:"limit,omitempty"`
	// Skip that many transactions before returning results (default: 0, maximum: 200000).
	Offset int64 `schema:"offset,omitempty"`
	// True - shows also ripple IOU transactions.
	IncludeIOUS bool `schema:"include_ious ,omitempty"`
}

// GetOrderStatusRequest used by GetOrderStatus method to map outgoing request data
type GetOrderStatusRequest struct {
	// Order ID.
	ID string `schema:"id,omitempty"`
	// Client order id. (Can only be used if order was placed with client order id parameter.)
	ClientOrderID string `schema:"client_order_id,omitempty"`
}

// CancelOrderRequest used by CancelOrder method to map outgoing request data
type CancelOrderRequest struct {
	// Order ID.
	ID string `schema:"id,omitempty"`
}

// CreateBuyLimitOrderRequest used by CreateBuyLimitOrder method to map outgoing request data
type CreateBuyLimitOrderRequest struct {
	Amount string `schema:"amount,omitempty"`
	Price  string `schema:"price,omitempty"`
	// if the order gets executed, a new sell order will be placed, with "limit_price" as its price.
	LimitPrice string `schema:"limit_price,omitempty"`
	// Opens buy limit order which will be canceled at 0:00 UTC unless it already has been executed. Possible value: True
	DailyOrder bool `schema:"daily_order,omitempty"`
	// An Immediate-Or-Cancel (IOC) order is an order that must be executed immediately.
	// Any portion of an IOC order that cannot be filled immediately will be cancelled. Possible value: True
	IOCOrder bool `schema:"ioc_order,omitempty"`
	// A Fill-Or-Kill (FOK) order is an order that must be executed immediately in its entirety.
	// If the order cannot be immediately executed in its entirety, it will be cancelled.
	//Possible value: True
	FOCOrder bool `schema:"foc_order,omitempty"`
	// Unique client order id set by client. Client order id needs to be unique string. Client order id value can only be used once.
	ClientOrderID string `schema:"client_order_id,omitempty"`
}

// CreateBuyInstantOrderRequest used by CreateBuyInstantOrder method to map outgoing request data
type CreateBuyInstantOrderRequest struct {
	// Amount in counter currency (Example: For BTC/USD pair, amount is quoted in USD)
	Amount string `schema:"amount"`
}

// CreateBuyMarketOrderRequest used by CreateBuyMarketOrder method to map outgoing request data
type CreateBuyMarketOrderRequest struct {
	// Amount in base currency (Example: For BTC/USD pair, amount is quoted in BTC)
	Amount string `schema:"amount"`
}

// CreateSellLimitOrderRequest used by CreateSellLimitOrder method to map outgoing request data
type CreateSellLimitOrderRequest struct {
	Amount string `schema:"amount,omitempty"`
	Price  string `schema:"price,omitempty"`
	// If the order gets executed, a new buy order will be placed, with "limit_price" as its price.
	LimitPrice string `schema:"limit_price,omitempty"`
	// Opens sell limit order which will be canceled at 0:00 UTC unless it already has been executed. Possible value: True
	DailyOrder bool `schema:"daily_order,omitempty"`
	// An Immediate-Or-Cancel (IOC) order is an order that must be executed immediately.
	// Any portion of an IOC order that cannot be filled immediately will be cancelled. Possible value: True
	IOCOrder bool `schema:"ioc_order,omitempty"`
	// A Fill-Or-Kill (FOK) order is an order that must be executed immediately in its entirety.
	// If the order cannot be immediately executed in its entirety, it will be cancelled.
	//Possible value: True
	FOCOrder bool `schema:"foc_order,omitempty"`
	// Unique client order id set by client. Client order id needs to be unique string. Client order id value can only be used once.
	ClientOrderID string `schema:"client_order_id,omitempty"`
}

// CreateSellInstantOrderRequest used by CreateSellInstantOrder method to map outgoing request data
type CreateSellInstantOrderRequest struct {
	// Amount in base currency (Example: For BTC/USD pair, amount is quoted in BTC)
	Amount string `schema:"amount"`
	// Instant sell orders allow you to sell an amount of the base currency determined by the value
	// of it in the counter-currency. Amount_in_counter sets the amount parameter to refer to the counter
	// currency instead of the base currency of the selected trading pair. Possible value: True
	AmountInCounter bool `schema:"amount_in_counter,omitempty"`
	// Unique client order id set by client. Client order id needs to be unique string. Client order id value can only be used once.
	ClientOrderID string `schema:"client_order_id,omitempty"`
}
