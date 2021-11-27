package bitstamp_test

import (
	"strings"
	"testing"

	"github.com/georlav/bitstamp"
)

func TestGetAllPairs(t *testing.T) {
	pairs := bitstamp.GetAllPairs()
	if len(pairs) == 0 {
		t.Fatal("Failed to retrieve pairs")
	}
}

func TestGetEuroPairs(t *testing.T) {
	pairs := bitstamp.GetEuroPairs()
	if len(pairs) == 0 {
		t.Fatal("Failed to retrieve euro pairs")
	}
}

func TestGetUSDPairs(t *testing.T) {
	pairs := bitstamp.GetUSDPairs()
	if len(pairs) == 0 {
		t.Fatal("Failed to retrieve usd pairs")
	}
}

func TestGetBTCPairs(t *testing.T) {
	pairs := bitstamp.GetBTCPairs()
	if len(pairs) == 0 {
		t.Fatal("Failed to retrieve btc pairs")
	}
}

func TestGetAllChannels(t *testing.T) {
	channels := bitstamp.GetAllChannels()
	if len(channels) == 0 {
		t.Fatal("Failed to retrieve channels")
	}
}

func TestGetEuroChannels(t *testing.T) {
	channels := bitstamp.GetEuroChannels()
	if len(channels) == 0 {
		t.Fatal("Failed to retrieve euro channels")
	}

	for i := range channels {
		if !strings.HasSuffix(channels[i].String(), "eur") {
			t.Fatalf("Invalid euro channel `%s`", channels)
		}
	}
}

func TestGetUSDChannels(t *testing.T) {
	channels := bitstamp.GetUSDChannels()
	if len(channels) == 0 {
		t.Fatal("Failed to retrieve usd channels")
	}

	for i := range channels {
		if !strings.HasSuffix(channels[i].String(), "usd") {
			t.Fatalf("Invalid usd channel `%s`", channels)
		}
	}
}

func TestGetBTCChannels(t *testing.T) {
	channels := bitstamp.GetBTCChannels()
	if len(channels) == 0 {
		t.Fatal("Failed to retrieve bitcoin channels")
	}

	for i := range channels {
		if !strings.HasSuffix(channels[i].String(), "btc") {
			t.Fatalf("Invalid bitcoin channel `%s`", channels)
		}
	}
}

func TestGetGBPChannels(t *testing.T) {
	channels := bitstamp.GetGBPChannels()
	if len(channels) == 0 {
		t.Fatal("Failed to retrieve GBP channels")
	}

	for i := range channels {
		if !strings.HasSuffix(channels[i].String(), "gbp") {
			t.Fatalf("Invalid GBP channel `%s`", channels)
		}
	}
}

func TestGetLiveTradeChannel(t *testing.T) {
	channel := bitstamp.GetLiveTradeChannel(bitstamp.BTCEUR)
	if expected := "live_trades_btceur"; expected != channel.String() {
		t.Fatalf("Invalid live trade channel expected `%s` got `%s`", expected, channel)
	}
}

func TestGetLiveTradesChannels(t *testing.T) {
	channels := bitstamp.GetLiveTradeChannels()
	if len(channels) == 0 {
		t.Fatal("Failed to retrieve live trade channels")
	}

	for i := range channels {
		if !strings.HasPrefix(channels[i].String(), "live_trades_") {
			t.Fatalf("Invalid live trade channel `%s`", channels[i])
		}
	}
}

func TestGetLiveOrderChannel(t *testing.T) {
	channels := bitstamp.GetLiveOrderChannels()
	if len(channels) == 0 {
		t.Fatal("Failed to retrieve live order channels")
	}

	for i := range channels {
		if !strings.HasPrefix(channels[i].String(), "live_orders_") {
			t.Fatalf("Invalid live trade channel `%s`", channels[i])
		}
	}
}

func TestGetLiveOrderChannels(t *testing.T) {
	channels := bitstamp.GetLiveOrderChannels()
	if len(channels) == 0 {
		t.Fatal("Failed to retrieve live order channels")
	}

	for i := range channels {
		if !strings.HasPrefix(channels[i].String(), "live_orders_") {
			t.Fatalf("Invalid live trade channel `%s`", channels[i])
		}
	}
}

func TestGetOrderBookChannels(t *testing.T) {
	channels := bitstamp.GetOrderBookChannels()
	if len(channels) == 0 {
		t.Fatal("Failed to retrieve order book channels")
	}

	for i := range channels {
		if !strings.HasPrefix(channels[i].String(), "order_book_") {
			t.Fatalf("Invalid order book channel `%s`", channels[i])
		}
	}
}

func TestGetDetailOrderBookChannels(t *testing.T) {
	channels := bitstamp.GetDetailOrderBookChannels()
	if len(channels) == 0 {
		t.Fatal("Failed to retrieve detail order book channels")
	}

	for i := range channels {
		if !strings.HasPrefix(channels[i].String(), "detail_order_book_") {
			t.Fatalf("Invalid detail order book channel `%s`", channels[i])
		}
	}
}

func TestGetDiffOrderBookChannels(t *testing.T) {
	channels := bitstamp.GetDiffOrderBookChannels()
	if len(channels) == 0 {
		t.Fatal("Failed to retrieve diff order book channels")
	}

	for i := range channels {
		if !strings.HasPrefix(channels[i].String(), "diff_order_book_") {
			t.Fatalf("Invalid diff order book channel `%s`", channels[i])
		}
	}
}
