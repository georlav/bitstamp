package bitstamp_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/georlav/bitstamp"
	"github.com/georlav/httprawmock"
)

func TestHTTPClient_GetTicker(t *testing.T) {
	testCases := []struct {
		description  string
		input        bitstamp.Pair
		responseFile string
		expectedCode int
	}{
		{
			description:  "Should fetch info for BTC/EUR",
			input:        bitstamp.BTCEUR,
			responseFile: "testdata/getticker_200.txt",
			expectedCode: http.StatusOK,
		},
		{
			description:  "Should fail to fetch info due to invalid pair",
			input:        bitstamp.NILNIL,
			responseFile: "testdata/getticker_404.txt",
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			b, err := os.ReadFile(tc.responseFile)
			if err != nil {
				t.Fatalf("failed to parse response file `%s`, %s", tc.responseFile, err)
			}

			ts := httprawmock.NewServer(
				httprawmock.NewRoute(http.MethodGet, "/api/v2/ticker/{pair}/", b),
			)
			defer t.Cleanup(ts.Close)

			c := bitstamp.NewHTTPAPI(
				bitstamp.BaseURLOption(ts.URL),
			)

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

func TestHTTPClient_GetTickerHourly(t *testing.T) {
	testCases := []struct {
		description  string
		input        bitstamp.Pair
		responseFile string
		expectedCode int
	}{
		{
			description:  "Should fetch info for ZRX/EUR",
			input:        bitstamp.ZRXEUR,
			responseFile: "testdata/getticker_200.txt",
			expectedCode: http.StatusOK,
		},
		{
			description:  "Should fail to fetch info due to invalid pair",
			input:        bitstamp.NILNIL,
			responseFile: "testdata/getticker_404.txt",
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			b, err := os.ReadFile(tc.responseFile)
			if err != nil {
				t.Fatalf("failed to parse response file `%s`, %s", tc.responseFile, err)
			}

			ts := httprawmock.NewServer(
				httprawmock.NewRoute(http.MethodGet, "/api/v2/ticker_hour/{pair}/", b),
			)
			defer t.Cleanup(ts.Close)

			c := bitstamp.NewHTTPAPI(
				bitstamp.BaseURLOption(ts.URL),
			)

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

func TestHTTPClient_GetAccountBalance(t *testing.T) {
	testCases := []struct {
		description  string
		pair         *bitstamp.Pair
		responseFile string
		expectedCode int
	}{
		{
			description:  "Should fetch all account balances",
			responseFile: "testdata/get_account_balances_200.txt",
			expectedCode: http.StatusOK,
		},
		{
			description:  "Should fetch account balance for BTC/EUR pair",
			pair:         &[]bitstamp.Pair{bitstamp.BTCEUR}[0],
			responseFile: "testdata/get_account_balance_200.txt",
			expectedCode: http.StatusOK,
		},
		{
			description:  "Should fail to fetch account balance due to invalid pair",
			pair:         &[]bitstamp.Pair{bitstamp.NILNIL}[0],
			responseFile: "testdata/not_found_page_404.txt",
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			b, err := os.ReadFile(tc.responseFile)
			if err != nil {
				t.Fatalf("failed to parse response file `%s`, %s", tc.responseFile, err)
			}

			ts := httprawmock.NewServer(
				httprawmock.NewRoute(http.MethodPost, "/api/v2/balance/", b),
				httprawmock.NewRoute(http.MethodPost, "/api/v2/balance/{pair}/", b),
			)
			defer t.Cleanup(ts.Close)

			c := bitstamp.NewHTTPAPI(
				bitstamp.BaseURLOption(ts.URL),
			)

			result, err := c.GetAccountBalance(context.Background(), tc.pair)
			if err != nil {
				apierr, ok := err.(bitstamp.Error)
				if !ok || apierr.StatusCode != tc.expectedCode {
					t.Fatalf("Failed to retrieve data, %s", err)
				}
			}

			if err == nil && result.BtcAvailable != "156.00000000" {
				t.Fatalf("Expected to have %s bitcoins got %s", "156.00000000", result.BtcAvailable)
			}

			_ = result
			// t.Logf("%+v", result)
		})
	}
}

func TestHTTPClient_GetUserTransactions(t *testing.T) {
	type input struct {
		pair    *bitstamp.Pair
		request bitstamp.GetUserTransactionsRequest
	}

	testCases := []struct {
		description  string
		input        input
		responseFile string
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
			responseFile: "testdata/get_user_transactions_200.txt",
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
			responseFile: "testdata/get_user_transactions_200_no_results.txt",
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
			responseFile: "testdata/get_user_transactions_200_with_validation_error.txt",
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			b, err := os.ReadFile(tc.responseFile)
			if err != nil {
				t.Fatalf("failed to parse response file `%s`, %s", tc.responseFile, err)
			}

			ts := httprawmock.NewServer(
				httprawmock.NewRoute(http.MethodPost, "/api/v2/user_transactions/", b),
				httprawmock.NewRoute(http.MethodPost, "/api/v2/user_transactions/{pair}/", b),
			)
			defer t.Cleanup(ts.Close)

			c := bitstamp.NewHTTPAPI(
				bitstamp.BaseURLOption(ts.URL),
			)

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
