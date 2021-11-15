package bitstamp_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/georlav/bitstamp"
)

func TestHTTPClient_GetTicker(t *testing.T) {
	testCases := []struct {
		description  string
		input        bitstamp.Pair
		expectedCode int
	}{
		{
			description:  "Should fetch info for BTC/EUR",
			input:        bitstamp.BTCEUR,
			expectedCode: http.StatusOK,
		},
		{
			description:  "Should fail to fetch info due to invalid pair",
			input:        bitstamp.Pair(123456789),
			expectedCode: http.StatusNotFound,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			resp, err := c.GetTicker(context.Background(), tc.input)
			if err != nil {
				if apierr, ok := err.(*bitstamp.APIError); ok && apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve pair `%s`, %s", tc.input, err)
				}
			}

			if err == nil && resp.High == "" {
				t.Fatal("Expected to have a value for high got none")
			}
		})
	}
}

func TestHTTPClient_GetTickerHourly(t *testing.T) {
	testCases := []struct {
		description  string
		input        bitstamp.Pair
		expectedCode int
	}{
		{
			description:  "Should fetch info for ZRX/EUR",
			input:        bitstamp.ZRXEUR,
			expectedCode: http.StatusOK,
		},
		{
			description:  "Should fail to fetch info due to invalid pair",
			input:        bitstamp.Pair(123456789),
			expectedCode: http.StatusNotFound,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			resp, err := c.GetTickerHourly(context.Background(), tc.input)
			if err != nil {
				if apierr, ok := err.(*bitstamp.APIError); ok && apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve pair `%s`, %s", tc.input, err)
				}
			}

			if err == nil && resp.High == "" {
				t.Fatal("Expected to have a value for high got none")
			}
		})
	}
}

func TestHTTPClient_GetOrderBook(t *testing.T) {
	testCases := []struct {
		description  string
		input        bitstamp.Pair
		expectedCode int
	}{
		{
			description:  "Should fetch order book for BTC/EUR",
			input:        bitstamp.BTCEUR,
			expectedCode: http.StatusOK,
		},
		{
			description:  "Should fail to fetch order book due to invalid pair",
			input:        bitstamp.Pair(123456789),
			expectedCode: http.StatusNotFound,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			resp, err := c.GetOrderBook(context.Background(), tc.input)
			if err != nil {
				if apierr, ok := err.(*bitstamp.APIError); ok && apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve pair `%s`, %s", tc.input, err)
				}
			}

			if err == nil && len(resp.Asks) == 0 {
				t.Fatal("Expected to have asks got none")
			}
			if err == nil && len(resp.Bids) == 0 {
				t.Fatal("Expected to have bids got none")
			}
		})
	}
}

func TestHTTPClient_GetTradingPairsInfo(t *testing.T) {
	testCases := []struct {
		description  string
		expectedCode int
	}{
		{
			description:  "Should fetch trading pairs information",
			expectedCode: http.StatusOK,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			resp, err := c.GetTradingPairsInfo(context.Background())
			if err != nil {
				t.Fatalf("Failed to retrieve trading pairs information %s", err)
			}

			if len(resp) == 0 {
				t.Fatal("Expected to have results got none")
			}
		})
	}
}

func TestHTTPClient_GetOHLCData(t *testing.T) {
	type input struct {
		pair    bitstamp.Pair
		request bitstamp.GetOHLCDataRequest
	}

	testCases := []struct {
		description  string
		input        input
		expectedCode int
	}{
		{
			description: "Should fetch OHLC datafor BTC/EUR",
			input: input{
				pair: bitstamp.BTCEUR,
				request: bitstamp.GetOHLCDataRequest{
					Limit: 10,
					Step:  7200,
				},
			},
			expectedCode: http.StatusOK,
		},
		{
			description: "Should fail to fetch order book due to validation error",
			input: input{
				pair: bitstamp.BTCEUR,
				request: bitstamp.GetOHLCDataRequest{
					Limit: 1000000,
					Step:  7200,
				},
			},
			expectedCode: http.StatusNotFound,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			results, err := c.GetOHLCData(context.Background(), tc.input.pair, tc.input.request)
			if err != nil {
				if apierr, ok := err.(*bitstamp.APIError); ok && apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve OHLC data for `%s`, %s", tc.input.pair, err)
				}
			}

			t.Log(err)

			if err == nil && len(results.Data.Ohlc) != int(tc.input.request.Limit) {
				t.Fatalf("Expected to have %d results got %d", tc.input.request.Limit, len(results.Data.Ohlc))
			}
		})
	}
}

func TestHTTPClient_GetEURUSDConversionRate(t *testing.T) {
	testCases := []struct {
		description  string
		expectedCode int
	}{
		{
			description:  "Should fetch trading pairs information",
			expectedCode: http.StatusOK,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			resp, err := c.GetEURUSDConversionRate(context.Background())
			if err != nil {
				t.Fatalf("Failed to retrieve trading pairs information %s", err)
			}

			if resp.Buy == "0.0" || resp.Buy == "" {
				t.Fatal("Expected to have a buy rate got none")
			}
			if resp.Sell == "0.0" || resp.Sell == "" {
				t.Fatal("Expected to have a buy rate got none")
			}
		})
	}
}

func TestHTTPClient_GetAccountBalance(t *testing.T) {
	if testing.Short() {
		t.Skipf("Skipping test %s in short mode", t.Name())
	}

	testCases := []struct {
		description  string
		pair         *bitstamp.Pair
		expectedCode int
	}{
		{
			description:  "Should fetch all account balances",
			expectedCode: http.StatusOK,
		},
		{
			description:  "Should fetch account balance for ZRX/EUR pair",
			pair:         &[]bitstamp.Pair{bitstamp.ZRXEUR}[0],
			expectedCode: http.StatusOK,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			_, err := c.GetAccountBalance(context.Background(), tc.pair)
			if err != nil {
				t.Fatalf("Failed to retrieve account balance, %s", err)
			}
		})
	}
}

func TestHTTPClient_GetUserTransactions(t *testing.T) {
	if testing.Short() {
		t.Skipf("Skipping test %s in short mode", t.Name())
	}

	type input struct {
		pair    *bitstamp.Pair
		request bitstamp.GetUserTransactionsRequest
	}

	testCases := []struct {
		description  string
		input        input
		expectedCode int
	}{
		{
			description: "Should fetch user transactions",
			input: input{
				pair: nil,
				request: bitstamp.GetUserTransactionsRequest{
					Limit:  1000,
					Offset: 0,
					Sort:   bitstamp.SortASC,
				},
			},
			expectedCode: http.StatusOK,
		},
		{
			description: "Should fetch user transactions for BTC/EUR",
			input: input{
				pair: &[]bitstamp.Pair{bitstamp.BTCEUR}[0],
				request: bitstamp.GetUserTransactionsRequest{
					Limit:  10,
					Offset: 0,
					Sort:   bitstamp.SortASC,
				},
			},
			expectedCode: http.StatusOK,
		},
		{
			description: "Should fail to fetch user transactions due to validation error",
			input: input{
				pair: nil,
				request: bitstamp.GetUserTransactionsRequest{
					Limit:  1000000,
					Offset: 0,
					Sort:   bitstamp.SortASC,
				},
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			_, err := c.GetUserTransactions(context.Background(), tc.input.pair, tc.input.request)
			if err != nil {
				if apierr, ok := err.(*bitstamp.APIError); ok && apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve user transactions, pair: `%s` request: `%+v`, %s", tc.input.pair, tc.input.request, err)
				}
			}
		})
	}
}

func TestHTTPClient_GetCryptoTransactions(t *testing.T) {
	if testing.Short() {
		t.Skipf("Skipping test %s in short mode", t.Name())
	}

	testCases := []struct {
		description  string
		input        bitstamp.GetCryptoTransactionsRequest
		expectedCode int
	}{
		{
			description: "Should fetch crypto transactions",
			input: bitstamp.GetCryptoTransactionsRequest{
				Limit:       1000,
				Offset:      0,
				IncludeIOUS: true,
			},
			expectedCode: http.StatusOK,
		},
		{
			description: "Should fail to fetch crypto transactions due to validation error",
			input: bitstamp.GetCryptoTransactionsRequest{
				Limit:       100000000,
				Offset:      0,
				IncludeIOUS: true,
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			_, err := c.GetCryptoTransactions(context.Background(), tc.input)
			if err != nil {
				if apierr, ok := err.(*bitstamp.APIError); ok && apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve crypto transactions, request: `%+v`, %s", tc.input, err)
				}
			}
		})
	}
}

func TestHTTPClient_GetWebsocketsToken(t *testing.T) {
	t.Skipf("Skipping test %s", t.Name())

	testCases := []struct {
		description  string
		input        bitstamp.Pair
		expectedCode int
	}{
		{
			description:  "Should fetch a websockets token",
			expectedCode: http.StatusOK,
		},
	}

	c := bitstamp.NewHTTPAPI(bitstamp.EnableDebugOption())

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			resp, err := c.GetWebsocketsToken(context.Background())
			if err != nil {
				if apierr, ok := err.(*bitstamp.APIError); ok && apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve pair `%s`, %s", tc.input, err)
				}
			}

			if err == nil && resp.Token == "" {
				t.Fatal("Expected to have a token got none")
			}
			if err == nil && resp.ValidSeconds == "" {
				t.Fatal("Expected to have a seconds value got none")
			}
		})
	}
}
