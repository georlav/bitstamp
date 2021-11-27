package bitstamp

// LiveTickerChannel object to map messages from live_trades_[currency_pair] channel
type LiveTickerChannel struct {
	Data struct {
		ID             int     `json:"id"`
		Timestamp      string  `json:"timestamp"`
		Amount         float64 `json:"amount"`
		AmountStr      string  `json:"amount_str"`
		Price          float64 `json:"price"`
		PriceStr       string  `json:"price_str"`
		Type           int     `json:"type"`
		Microtimestamp string  `json:"microtimestamp"`
		BuyOrderID     int64   `json:"buy_order_id"`
		SellOrderID    int64   `json:"sell_order_id"`
	} `json:"data"`
	Channel string `json:"channel"`
	Event   string `json:"event"`
}

// LiveOrdersChannel object to map messages from live_orders_[currency_pair] channel
type LiveOrdersChannel struct {
	Data struct {
		ID             int64   `json:"id"`
		IDStr          string  `json:"id_str"`
		OrderType      int     `json:"order_type"`
		Datetime       string  `json:"datetime"`
		Microtimestamp string  `json:"microtimestamp"`
		Amount         float64 `json:"amount"`
		AmountStr      string  `json:"amount_str"`
		Price          float64 `json:"price"`
		PriceStr       string  `json:"price_str"`
	} `json:"data"`
	Channel string `json:"channel"`
	Event   string `json:"event"`
}

// LiveOrderBookChannel object to map messages from order_book_[currency_pair] channel
type LiveOrderBookChannel struct {
	Data struct {
		Timestamp      string `json:"timestamp"`
		Microtimestamp string `json:"microtimestamp"`
		// List of top 100 bids
		Bids [][]string `json:"bids"`
		// List of top 100 asks
		Asks [][]string `json:"asks"`
	} `json:"data"`
	Channel string `json:"channel"`
	Event   string `json:"event"`
}

// LiveDetailOrderBookChannel object to map messages from detail_order_book_[currency_pair] channel
type LiveDetailOrderBookChannel struct {
	Data struct {
		Timestamp      string `json:"timestamp"`
		Microtimestamp string `json:"microtimestamp"`
		// List of top 100 bids [price, amount, order id].
		Bids [][]string `json:"bids"`
		// List of top 100 asks [price, amount, order id].
		Asks [][]string `json:"asks"`
	} `json:"data"`
	Channel string `json:"channel"`
	Event   string `json:"event"`
}

// LiveFullOrderBook object to map messages from diff_order_book_[currency_pair] channel
type LiveFullOrderBook struct {
	Data struct {
		Timestamp      string `json:"timestamp"`
		Microtimestamp string `json:"microtimestamp"`
		// List of changed bids since last broadcast.
		Bids [][]string `json:"bids"`
		// List of changed asks since last broadcast.
		Asks [][]string `json:"asks"`
	} `json:"data"`
	Channel string `json:"channel"`
	Event   string `json:"event"`
}

// WebSocketMessage object to map any message
type WebSocketMessage struct {
	Event   string      `json:"event"`
	Channel string      `json:"channel"`
	Data    interface{} `json:"data"`
}
