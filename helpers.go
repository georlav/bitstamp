package bitstamp

import (
	"strings"
)

func GetAllPairs() []Pair {
	var result []Pair

	for k := range pairs {
		result = append(result, k)
	}

	return result
}

func GetEuroPairs() []Pair {
	var result []Pair

	for key := range pairs {
		if strings.HasSuffix(key.String(), "eur") {
			result = append(result, key)
		}
	}

	return result
}

func GetUSDPairs() []Pair {
	var result []Pair

	for key := range pairs {
		if strings.HasSuffix(key.String(), "usd") {
			result = append(result, key)
		}
	}

	return result
}

func GetBTCPairs() []Pair {
	var result []Pair

	for key := range pairs {
		if strings.HasSuffix(key.String(), "btc") {
			result = append(result, key)
		}
	}

	return result
}

func GetGBPPairs() []Pair {
	var result []Pair

	for key := range pairs {
		if strings.HasSuffix(key.String(), "gbp") {
			result = append(result, key)
		}
	}

	return result
}

func GetAllChannels() []Channel {
	var result []Channel

	for key := range channels {
		result = append(result, key)
	}

	return result
}

func GetEuroChannels() []Channel {
	var result []Channel

	for key := range channels {
		if strings.HasSuffix(key.String(), "eur") {
			result = append(result, key)
		}
	}

	return result
}

func GetUSDChannels() []Channel {
	var result []Channel

	for key := range channels {
		if strings.HasSuffix(key.String(), "usd") {
			result = append(result, key)
		}
	}

	return result
}

func GetBTCChannels() []Channel {
	var result []Channel

	for key := range channels {
		if strings.HasSuffix(key.String(), "btc") {
			result = append(result, key)
		}
	}

	return result
}

func GetGBPChannels() []Channel {
	var result []Channel

	for key := range channels {
		if strings.HasSuffix(key.String(), "gbp") {
			result = append(result, key)
		}
	}

	return result
}

func GetLiveTradeChannels() []Channel {
	var result []Channel

	for key := range channels {
		if strings.HasPrefix(key.String(), "live_trades_") {
			result = append(result, key)
		}
	}

	return result
}

func GetLiveOrderChannels() []Channel {
	var result []Channel

	for key := range channels {
		if strings.HasPrefix(key.String(), "live_orders_") {
			result = append(result, key)
		}
	}

	return result
}

func GetOrderBookChannels() []Channel {
	var result []Channel

	for key := range channels {
		if strings.HasPrefix(key.String(), "order_book_") {
			result = append(result, key)
		}
	}

	return result
}

func GetDetailOrderBookChannels() []Channel {
	var result []Channel

	for key := range channels {
		if strings.HasPrefix(key.String(), "detail_order_book_") {
			result = append(result, key)
		}
	}

	return result
}

func GetDiffOrderBookChannels() []Channel {
	var result []Channel

	for key := range channels {
		if strings.HasPrefix(key.String(), "diff_order_book_") {
			result = append(result, key)
		}
	}

	return result
}
