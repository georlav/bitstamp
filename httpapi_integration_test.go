package bitstamp_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/georlav/bitstamp"
)

func TestHTTPClient_GetTicker_Integration(t *testing.T) {
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
			input:        bitstamp.NILNIL,
			expectedCode: http.StatusNotFound,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := c.GetTicker(context.Background(), tc.input)
			if err != nil {
				apierr, ok := err.(bitstamp.Error)
				if !ok || apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			if err == nil && result.High == "" {
				t.Fatal("Expected to have a value for high got none")
			}

			_ = result
			// t.Logf("%+v", result)
		})
	}
}

func TestHTTPClient_GetTickerHourly_Integration(t *testing.T) {
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
			input:        bitstamp.NILNIL,
			expectedCode: http.StatusNotFound,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := c.GetTickerHourly(context.Background(), tc.input)
			if err != nil {
				apierr, ok := err.(bitstamp.Error)
				if !ok || apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			if err == nil && result.High == "" {
				t.Fatal("Expected to have a value for high got none")
			}

			_ = result
			// t.Logf("%+v", result)
		})
	}
}

func TestHTTPClient_GetOrderBook_Integration(t *testing.T) {
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
			input:        bitstamp.NILNIL,
			expectedCode: http.StatusNotFound,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := c.GetOrderBook(context.Background(), tc.input)
			if err != nil {
				apierr, ok := err.(bitstamp.Error)
				if !ok || apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			if err == nil && len(result.Asks) == 0 {
				t.Fatal("Expected to have asks got none")
			}
			if err == nil && len(result.Bids) == 0 {
				t.Fatal("Expected to have bids got none")
			}

			_ = result
			// t.Logf("%+v", result)
		})
	}
}

func TestHTTPClient_GetTransactions_Integration(t *testing.T) {
	type input struct {
		pair    bitstamp.Pair
		request bitstamp.GetTransactionsRequest
	}

	testCases := []struct {
		description  string
		input        input
		expectedCode int
	}{
		{
			description: "Should fetch hourly transactions",
			input: input{
				pair: bitstamp.BTCEUR,
				request: bitstamp.GetTransactionsRequest{
					Time: "minute",
				},
			},
			expectedCode: http.StatusOK,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := c.GetTransactions(context.Background(), tc.input.pair, tc.input.request)
			if err != nil {
				apierr, ok := err.(bitstamp.Error)
				if !ok || apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			if err == nil && len(result) == 0 {
				t.Fatal("Expected to have transactions got none")
			}

			_ = result
			// t.Logf("%+v", result)
		})
	}
}

func TestHTTPClient_GetTradingPairsInfo_Integration(t *testing.T) {
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
			result, err := c.GetTradingPairsInfo(context.Background())
			if err != nil {
				apierr, ok := err.(bitstamp.Error)
				if !ok || apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			if len(result) == 0 {
				t.Fatal("Expected to have results got none")
			}

			_ = result
			// t.Logf("%+v", result)
		})
	}
}

func TestHTTPClient_GetOHLCData_Integration(t *testing.T) {
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
			expectedCode: http.StatusBadRequest,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := c.GetOHLCData(context.Background(), tc.input.pair, tc.input.request)
			if err != nil {
				apierr, ok := err.(bitstamp.Error)
				if !ok || apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			if err == nil && len(result.Data.Ohlc) != int(tc.input.request.Limit) {
				t.Fatalf("Expected to have %d results got %d", tc.input.request.Limit, len(result.Data.Ohlc))
			}

			_ = result
			// t.Logf("%+v", result)
		})
	}
}

func TestHTTPClient_GetEURUSDConversionRate_Integration(t *testing.T) {
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
			result, err := c.GetEURUSDConversionRate(context.Background())
			if err != nil {
				apierr, ok := err.(bitstamp.Error)
				if !ok || apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			if result.Buy == "0.0" || result.Buy == "" {
				t.Fatal("Expected to have a buy rate got none")
			}
			if result.Sell == "0.0" || result.Sell == "" {
				t.Fatal("Expected to have a buy rate got none")
			}

			_ = result
			// t.Logf("%+v", result)
		})
	}
}

func TestHTTPClient_GetAccountBalance_Integration(t *testing.T) {
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
			description:  "Should fetch account balance for BTC/EUR pair",
			pair:         &[]bitstamp.Pair{bitstamp.BTCEUR}[0],
			expectedCode: http.StatusOK,
		},
		{
			description:  "Should fail to fetch account balance due to invalid pair",
			pair:         &[]bitstamp.Pair{bitstamp.NILNIL}[0],
			expectedCode: http.StatusNotFound,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := c.GetAccountBalance(context.Background(), tc.pair)
			if err != nil {
				apierr, ok := err.(bitstamp.Error)
				if !ok || apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			_ = result
			// t.Logf("%+v", result)
		})
	}
}

func TestHTTPClient_GetUserTransactions_Integration(t *testing.T) {
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
			// API responds with status 200 and error
			expectedCode: http.StatusOK,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := c.GetUserTransactions(context.Background(), tc.input.pair, tc.input.request)
			if err != nil {
				apierr, ok := err.(bitstamp.Error)
				if !ok || apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			_ = result
			// t.Logf("%+v", result)
		})
	}
}

func TestHTTPClient_GetCryptoTransactions_Integration(t *testing.T) {
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
			result, err := c.GetCryptoTransactions(context.Background(), tc.input)
			if err != nil {
				apierr, ok := err.(bitstamp.Error)
				if !ok || apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			_ = result
			// t.Logf("%+v", result)
		})
	}
}

func TestHTTPClient_GetOpenOrders_Integration(t *testing.T) {
	if testing.Short() {
		t.Skipf("Skipping test %s in short mode", t.Name())
	}

	testCases := []struct {
		description  string
		expectedCode int
	}{
		{
			description:  "Should retrieve open orders",
			expectedCode: http.StatusOK,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := c.GetOpenOrders(context.Background())
			if err != nil {
				apierr, ok := err.(bitstamp.Error)
				if !ok || apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			if len(result) > 0 && result[0].CurrencyPair == "" {
				t.Fatal("Order expected to have a currency pair value got none")
			}

			if len(result) > 0 && result[0].ID == "" {
				t.Fatal("Order expected to have a unique identifier value got none")
			}

			_ = result
			// t.Logf("%+v", result)
		})
	}
}

func TestHTTPClient_GetOrderStatus_Integration(t *testing.T) {
	if testing.Short() {
		t.Skipf("Skipping test %s in short mode", t.Name())
	}

	testCases := []struct {
		description  string
		input        bitstamp.GetOrderStatusRequest
		expectedCode int
	}{
		// {
		// 	description: "Should fetch order status (test case requires a valid id)",
		// 	input: bitstamp.GetOrderStatusRequest{
		// 		ID: "0000000000000000",
		// 	},
		// 	expectedCode: http.StatusOK,
		// },
		{
			description: "Should fail to fetch order status due to unknown id",
			input: bitstamp.GetOrderStatusRequest{
				ID: "123456789",
			},
			expectedCode: http.StatusTeapot,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := c.GetOrderStatus(context.Background(), tc.input)
			if err != nil {
				apierr, ok := err.(bitstamp.Error)
				if !ok || apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			_ = result
			// t.Logf("%+v", result)
		})
	}
}

func TestHTTPClient_CancelOrder_Integration(t *testing.T) {
	if testing.Short() {
		t.Skipf("Skipping test %s in short mode", t.Name())
	}

	testCases := []struct {
		description  string
		input        bitstamp.CancelOrderRequest
		expectedCode int
	}{
		// {
		// 	description: "Should cancel order by id (test case requires a valid id)",
		// 	input: bitstamp.CancelOrderRequest{
		// 		ID: "0000000000000000",
		// 	},
		// 	expectedCode: http.StatusOK,
		// },
		{
			description: "Should fail due to invalid id order by id",
			input: bitstamp.CancelOrderRequest{
				ID: "xxx",
			},
			expectedCode: http.StatusTeapot,
		},
		{
			description: "Should fail to cancel order due to unknown id",
			input: bitstamp.CancelOrderRequest{
				ID: "1234567890",
			},
			expectedCode: http.StatusTeapot,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := c.CancelOrder(context.Background(), tc.input)
			if err != nil {
				apierr, ok := err.(bitstamp.Error)
				if !ok || apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			_ = result
			t.Logf("%+v", result)
		})
	}
}

func TestHTTPClient_CancelAllOrders_Integration(t *testing.T) {
	if testing.Short() {
		t.Skipf("Skipping test %s in short mode", t.Name())
	}

	testCases := []struct {
		description  string
		input        *bitstamp.Pair
		expectedCode int
	}{
		{
			description:  "Should cancel all orders",
			expectedCode: http.StatusOK,
		},
		{
			description:  "Should cancel BTC/EUR orders",
			input:        &[]bitstamp.Pair{bitstamp.BTCEUR}[0],
			expectedCode: http.StatusOK,
		},
		{
			description:  "Should fail to cancel orders",
			input:        &[]bitstamp.Pair{bitstamp.NILNIL}[0],
			expectedCode: http.StatusNotFound,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := c.CancelAllOrders(context.Background(), tc.input)
			if err != nil {
				apierr, ok := err.(bitstamp.Error)
				if !ok || apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			_ = result
			// t.Logf("%+v", result)
		})
	}
}

func TestHTTPClient_CreateBuyLimitOrder_Integration(t *testing.T) {
	if testing.Short() {
		t.Skipf("Skipping test %s in short mode", t.Name())
	}

	type input struct {
		pair    bitstamp.Pair
		request bitstamp.CreateBuyLimitOrderRequest
	}

	testCases := []struct {
		description  string
		input        input
		expectedCode int
	}{
		// {
		// 	description: "Should create a buy limit order (test case might trigger an actual buy)",
		// 	input: input{
		// 		pair: bitstamp.ZRXEUR,
		// 		request: bitstamp.CreateBuyLimitOrderRequest{
		// 			// buy 250 zrx
		// 			Amount: "250",
		// 			// At 0.86 euro
		// 			Price: "0.86",
		// 			// Sell if price reaches 2.011 euro
		// 			LimitPrice: "2.011",
		// 		},
		// 	},
		// 	expectedCode: http.StatusOK,
		// },
		{
			description: "Should fail to create a buy limit order",
			input: input{
				pair: bitstamp.ZRXEUR,
				request: bitstamp.CreateBuyLimitOrderRequest{
					Amount:     "-10",
					Price:      "-100.00",
					LimitPrice: "0.3999",
				},
			},
			expectedCode: http.StatusTeapot,
		},
		{
			description: "Should fail to create a buy limit order due to invalid pair",
			input: input{
				pair: bitstamp.NILNIL,
				request: bitstamp.CreateBuyLimitOrderRequest{
					Amount: "0.0",
				},
			},
			expectedCode: http.StatusNotFound,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := c.CreateBuyLimitOrder(context.Background(), tc.input.pair, tc.input.request)
			if err != nil {
				apiErr, ok := err.(bitstamp.Error)
				if !ok || apiErr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			_ = result
			// t.Logf("%+v", result)
		})
	}
}

func TestHTTPClient_CreateBuyInstantOrder_Integration(t *testing.T) {
	if testing.Short() {
		t.Skipf("Skipping test %s in short mode", t.Name())
	}

	type input struct {
		pair    bitstamp.Pair
		request bitstamp.CreateBuyInstantOrderRequest
	}

	testCases := []struct {
		description  string
		input        input
		expectedCode int
	}{
		// {
		// 	description: "Should create a buy instant order (Warning: test case might trigger an actual buy)",
		// 	input: input{
		// 		pair: bitstamp.BTCEUR,
		// 		request: bitstamp.CreateBuyInstantOrderRequest{
		// 			Amount: "20.01",
		// 		},
		// 	},
		// 	expectedCode: http.StatusOK,
		// },
		{
			description: "Should create a buy instant order (Warning: test case might trigger an actual buy)",
			input: input{
				pair: bitstamp.BTCEUR,
				request: bitstamp.CreateBuyInstantOrderRequest{
					Amount: "10.01",
				},
			},
			expectedCode: http.StatusOK,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := c.CreateBuyInstantOrder(context.Background(), tc.input.pair, tc.input.request)
			if err != nil {
				apiErr, ok := err.(bitstamp.Error)
				if !ok || apiErr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			_ = result
			// t.Logf("%+v", result)
		})
	}
}

func TestHTTPClient_CreateSellInstantOrder_Integration(t *testing.T) {
	if testing.Short() {
		t.Skipf("Skipping test %s in short mode", t.Name())
	}

	type input struct {
		pair    bitstamp.Pair
		request bitstamp.CreateSellInstantOrderRequest
	}

	testCases := []struct {
		description  string
		input        input
		expectedCode int
	}{
		{
			description: "Should create a sell instant order (test case might trigger an actual sell)",
			input: input{
				pair: bitstamp.BTCEUR,
				request: bitstamp.CreateSellInstantOrderRequest{
					// sell 10 btc
					Amount: "10",
				},
			},
			expectedCode: http.StatusOK,
		},
		{
			description: "Should fail to create a sell instant order due to invalid amount",
			input: input{
				pair: bitstamp.BTCEUR,
				request: bitstamp.CreateSellInstantOrderRequest{
					Amount: "-10.66",
				},
			},
			expectedCode: http.StatusOK,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := c.CreateSellInstantOrder(context.Background(), tc.input.pair, tc.input.request)
			if err != nil {
				apiErr, ok := err.(bitstamp.Error)
				if !ok || apiErr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			_ = result
			// t.Logf("%+v", result)
		})
	}
}

func TestHTTPClient_CreateSellLimitOrder_Integration(t *testing.T) {
	if testing.Short() {
		t.Skipf("Skipping test %s in short mode", t.Name())
	}

	type input struct {
		pair    bitstamp.Pair
		request bitstamp.CreateSellLimitOrderRequest
	}

	testCases := []struct {
		description  string
		input        input
		expectedCode int
	}{
		// {
		// 	description: "Should create a sell limit order (test case might trigger an actual sell)",
		// 	input: input{
		// 		pair: bitstamp.ZRXEUR,
		// 		request: bitstamp.CreateSellLimitOrderRequest{
		// 			// Sell 20 ZRXEUR
		// 			Amount: "10",
		// 			// At 100 euro
		// 			Price: "100.00",
		// 			// Buy again if price falls to 0.39 euro
		// 			LimitPrice: "0.3999",
		// 		},
		// 	},
		// 	expectedCode: http.StatusOK,
		// },
		{
			description: "Should fail to create a sell limit order",
			input: input{
				pair: bitstamp.ZRXEUR,
				request: bitstamp.CreateSellLimitOrderRequest{
					Amount:     "-10",
					Price:      "-100.00",
					LimitPrice: "0.3999",
				},
			},
			expectedCode: http.StatusTeapot,
		},
		{
			description: "Should fail to create a sell limit order due to invalid pair",
			input: input{
				pair: bitstamp.NILNIL,
				request: bitstamp.CreateSellLimitOrderRequest{
					Amount: "0.0",
				},
			},
			expectedCode: http.StatusNotFound,
		},
	}

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := c.CreateSellLimitOrder(context.Background(), tc.input.pair, tc.input.request)
			if err != nil {
				apiErr, ok := err.(bitstamp.Error)
				if !ok || apiErr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			_ = result
			// t.Logf("%+v", result)
		})
	}
}

func TestHTTPClient_GetWebsocketsToken_Integration(t *testing.T) {
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

	c := bitstamp.NewHTTPAPI()

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := c.GetWebsocketsToken(context.Background())
			if err != nil {
				apierr, ok := err.(bitstamp.Error)
				if !ok || apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			if err == nil && result.Token == "" {
				t.Fatal("Expected to have a token got none")
			}
			if err == nil && result.ValidSeconds == "" {
				t.Fatal("Expected to have a seconds value got none")
			}

			_ = result
			// t.Logf("%+v", result)
		})
	}
}
