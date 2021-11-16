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
