package bitstamp

import (
	"fmt"
	"strings"
)

func GetAllPairs() []Pair {
	var result []Pair

	for k := range getPairs() {
		result = append(result, k)
	}

	return result
}

func GetEuroPairs() []Pair {
	var result []Pair

	for key := range getPairs() {
		if strings.HasSuffix(key.String(), "eur") {
			result = append(result, key)
		}
	}

	return result
}

func GetUSDPairs() []Pair {
	var result []Pair

	for key := range getPairs() {
		if strings.HasSuffix(key.String(), "usd") {
			result = append(result, key)
		}
	}

	return result
}

func GetBTCPairs() []Pair {
	var result []Pair

	for key := range getPairs() {
		if strings.HasSuffix(key.String(), "btc") {
			result = append(result, key)
		}
	}

	return result
}

func GetGBPPairs() []Pair {
	var result []Pair

	for key := range getPairs() {
		if strings.HasSuffix(key.String(), "gbp") {
			result = append(result, key)
		}
	}

	return result
}

func GetAllChannels() []Channel {
	var result []Channel

	for key := range getChannels() {
		result = append(result, key)
	}

	return result
}

func GetEuroChannels() []Channel {
	var result []Channel

	for key := range getChannels() {
		if strings.HasSuffix(key.String(), "eur") {
			result = append(result, key)
		}
	}

	return result
}

func GetUSDChannels() []Channel {
	var result []Channel

	for key := range getChannels() {
		if strings.HasSuffix(key.String(), "usd") {
			result = append(result, key)
		}
	}

	return result
}

func GetBTCChannels() []Channel {
	var result []Channel

	for key := range getChannels() {
		if strings.HasSuffix(key.String(), "btc") {
			result = append(result, key)
		}
	}

	return result
}

func GetGBPChannels() []Channel {
	var result []Channel

	for key := range getChannels() {
		if strings.HasSuffix(key.String(), "gbp") {
			result = append(result, key)
		}
	}

	return result
}

// GetLiveTradeChannel get live trades channel for a pair. (live_trades_[currency_pair])
func GetLiveTradeChannel(p Pair) Channel {
	return getChannel("live_trades_", p)
}

// GetLiveTradeChannel get all live trade channels. (live_trades_[*])
func GetLiveTradeChannels() []Channel {
	return getChannelsByPrefix("live_trades_")
}

// GetLiveTradeChannel get live orders channel for a pair. (live_orders_[currency_pair])
func GetLiveOrderChannel(p Pair) Channel {
	return getChannel("live_orders_", p)
}

// GetLiveOrderChannels get all live order channels. (live_orders_[*])
func GetLiveOrderChannels() []Channel {
	return getChannelsByPrefix("live_orders_")
}

// GetOrderBookChannel get order book channel for a pair (order_book_[currency_pair])
func GetOrderBookChannel(p Pair) Channel {
	return getChannel("order_book_", p)
}

// GetOrderBookChannel get all order book channels order_book_[*]
func GetOrderBookChannels() []Channel {
	return getChannelsByPrefix("order_book_")
}

// GetDetailOrderBookChannel get detail order book for a pair (detail_order_book_[currency_pair])
func GetDetailOrderBookChannel(p Pair) Channel {
	return getChannel("detail_order_book_", p)
}

// GetDetailOrderBookChannels get all detail order book channels (detail_order_book_[*])
func GetDetailOrderBookChannels() []Channel {
	return getChannelsByPrefix("detail_order_book_")
}

// GetDiffOrderBookChannel get diff order book channel for a pair (diff_order_book_[currency_pair])
func GetDiffOrderBookChannel(p Pair) Channel {
	return getChannel("diff_order_book_", p)
}

// GetDiffOrderBookChannels get all diff order book channels (diff_order_book_[*])
func GetDiffOrderBookChannels() []Channel {
	return getChannelsByPrefix("diff_order_book_")
}

func getChannel(prefix string, suffix Pair) Channel {
	chans := getChannels()
	ch := fmt.Sprintf("%s%s", prefix, suffix)

	for i := range chans {
		if v, ok := chans[i]; ok && v == ch {
			return i
		}
	}

	return 0
}

func getChannelsByPrefix(p string) []Channel {
	var result []Channel

	for key := range getChannels() {
		if strings.HasPrefix(key.String(), p) {
			result = append(result, key)
		}
	}

	return result
}
