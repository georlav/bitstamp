// package bitstamp
//
// REQUEST LIMITS
// Do not do more than 8000 requests per 10 minutes or your IP address will be banned.
// For real time data use websocket API.
package bitstamp

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/schema"
)

type HTTPAPI struct {
	baseURL string
	key     string
	secret  string
	debug   bool
	handle  *http.Client
}

func NewHTTPAPI(options ...option) *HTTPAPI {
	api := HTTPAPI{
		baseURL: "https://www.bitstamp.net",
		handle:  http.DefaultClient,
		debug:   false,
	}

	// override defaults via available functional options
	for i := range options {
		options[i](&api)
	}

	//  override API key and secret using env variable
	if key := os.Getenv("BITSTAMP_KEY"); key != "" {
		api.key = key
	}
	if secret := os.Getenv("BITSTAMP_SECRET"); secret != "" {
		api.secret = secret
	}

	return &api
}

// Ticker retrieves information for a pair
// Docs https://www.bitstamp.net/api/#ticker
func (h *HTTPAPI) GetTicker(ctx context.Context, p Pair) (*GetTickerResponse, error) {
	resp, err := h.doRequest(ctx, http.MethodGet, fmt.Sprintf(tickerURL, p), nil, false)
	if err != nil {
		return nil, err
	}

	var result GetTickerResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTickerHourly retrieves information for a pair
// Docs https://www.bitstamp.net/api/#ticker-hour
func (h *HTTPAPI) GetTickerHourly(ctx context.Context, p Pair) (*GetTickerResponse, error) {
	resp, err := h.doRequest(ctx, http.MethodGet, fmt.Sprintf(tickerHourlyURL, p), nil, false)
	if err != nil {
		return nil, err
	}

	var result GetTickerResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetOrderBook retrieve information for all supported trading pairs
// Docs https://www.bitstamp.net/api/#order-book
func (h *HTTPAPI) GetOrderBook(ctx context.Context, p Pair) (*GetOrderBookResponse, error) {
	resp, err := h.doRequest(ctx, http.MethodGet, fmt.Sprintf(orderBookURL, p), nil, false)
	if err != nil {
		return nil, err
	}

	var result GetOrderBookResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTransactions retrieve transactions
// Docs https://www.bitstamp.net/api/#transactions
func (h *HTTPAPI) GetTransactions(ctx context.Context, p Pair, r GetTransactionsRequest) ([]GetTransactionResponse, error) {
	params := url.Values{}
	if err := schema.NewEncoder().Encode(r, params); err != nil {
		return nil, err
	}

	resp, err := h.doRequest(ctx, http.MethodGet, fmt.Sprintf(transactions+"?"+params.Encode(), p), nil, false)
	if err != nil {
		return nil, err
	}

	var result []GetTransactionResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetTradingPairsInfo retrieve information for all supported trading pairs
// Docs https://www.bitstamp.net/api/#trading-pairs-info
func (h *HTTPAPI) GetTradingPairsInfo(ctx context.Context) ([]GetTradingPairInfoResult, error) {
	resp, err := h.doRequest(ctx, http.MethodGet, tradingPairsInfoURL, nil, false)
	if err != nil {
		return nil, err
	}

	var result []GetTradingPairInfoResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetOHLCData retrieve information about open,high,low,close values
// Docs https://www.bitstamp.net/api/#ohlc_data
func (h *HTTPAPI) GetOHLCData(ctx context.Context, p Pair, r GetOHLCDataRequest) (*GetOHLCDataResponse, error) {
	params := url.Values{}
	if err := schema.NewEncoder().Encode(r, params); err != nil {
		return nil, err
	}

	resp, err := h.doRequest(ctx, http.MethodGet, fmt.Sprintf(ohlcDataURL+"?"+params.Encode(), p), nil, false)
	if err != nil {
		return nil, err
	}

	var result GetOHLCDataResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetEURUSDConversionRate retrieves EUR/USD convension rate
// Docs https://www.bitstamp.net/api/#conversion-rate
func (h *HTTPAPI) GetEURUSDConversionRate(ctx context.Context) (*GetEURUSDConversionRateResult, error) {
	resp, err := h.doRequest(ctx, http.MethodGet, eurusdConversionRateURL, nil, false)
	if err != nil {
		return nil, err
	}

	var result GetEURUSDConversionRateResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAccountBalance retrieves balances of given pair, if pair is nil returns all account balances
// Docs https://www.bitstamp.net/api/#account-balance
func (h *HTTPAPI) GetAccountBalance(ctx context.Context, p *Pair) (*GetAccountBalancesResponse, error) {
	u := accountBalanceURL
	if p != nil {
		u += p.String() + "/"
	}

	resp, err := h.doRequest(ctx, http.MethodPost, u, nil, true)
	if err != nil {
		return nil, err
	}

	var result GetAccountBalancesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetUserTransactions retrieves current account transactions, if pair is nil returns all user transactions
// Docs https://www.bitstamp.net/api/#user-transactions
func (h *HTTPAPI) GetUserTransactions(ctx context.Context, p *Pair, r GetUserTransactionsRequest) ([]GetUserTransactionResponse, error) {
	params := url.Values{}
	if err := schema.NewEncoder().Encode(r, params); err != nil {
		return nil, err
	}

	u := userTransactionsURL
	if p != nil {
		u += p.String() + "/"
	}

	resp, err := h.doRequest(ctx, http.MethodPost, u, bytes.NewBuffer([]byte(params.Encode())), true)
	if err != nil {
		return nil, err
	}

	var result []GetUserTransactionResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		// handle status 200 with error
		return nil, newErrorFromResponse(resp)
	}

	return result, nil
}

// GetCryptoTransactions retrieves data for all cryptocurrency deposits and withdrawals.
// Docs https://www.bitstamp.net/api/#crypto-transactions
func (h *HTTPAPI) GetCryptoTransactions(ctx context.Context, r GetCryptoTransactionsRequest) (*GetCryptoTransactionsResponse, error) {
	params := url.Values{}
	if err := schema.NewEncoder().Encode(r, params); err != nil {
		return nil, err
	}

	resp, err := h.doRequest(ctx, http.MethodPost, cryptoTransactionsURL, bytes.NewBuffer([]byte(params.Encode())), true)
	if err != nil {
		return nil, err
	}

	var result GetCryptoTransactionsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetOpenOrders retrieves open orders, after call data is cached for 10 seconds
// Docs https://www.bitstamp.net/api/#open-orders
func (h *HTTPAPI) GetOpenOrders(ctx context.Context) ([]GetOpenOrderResponse, error) {
	resp, err := h.doRequest(ctx, http.MethodPost, openOrdersURL, nil, true)
	if err != nil {
		return nil, err
	}

	var result []GetOpenOrderResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetOrderStatus retrieves order status
// Docs https://www.bitstamp.net/api/#order-status
func (h *HTTPAPI) GetOrderStatus(ctx context.Context, r GetOrderStatusRequest) (*GetOrderStatusResponse, error) {
	params := url.Values{}
	if err := schema.NewEncoder().Encode(r, params); err != nil {
		return nil, err
	}

	resp, err := h.doRequest(ctx, http.MethodPost, orderStatusURL, bytes.NewBuffer([]byte(params.Encode())), true)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	teeReader := io.TeeReader(resp.Body, &buf)

	var result GetOrderStatusResponse
	if err := json.NewDecoder(teeReader).Decode(&result); err != nil {
		return nil, err
	}

	// handle status 200 with error
	if result.ID == 0 {
		resp.Body = ioutil.NopCloser(bytes.NewReader(buf.Bytes()))
		resp.StatusCode = http.StatusTeapot
		return nil, newErrorFromResponse(resp)
	}

	return &result, nil
}

// CancelOrder cancels an order by id
// Docs https://www.bitstamp.net/api/#cancel-order
func (h *HTTPAPI) CancelOrder(ctx context.Context, r CancelOrderRequest) (*CancelOrderResponse, error) {
	params := url.Values{}
	if err := schema.NewEncoder().Encode(r, params); err != nil {
		return nil, err
	}

	resp, err := h.doRequest(ctx, http.MethodPost, cancelOrderURL, bytes.NewBuffer([]byte(params.Encode())), true)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	teeReader := io.TeeReader(resp.Body, &buf)

	var result CancelOrderResponse
	if err := json.NewDecoder(teeReader).Decode(&result); err != nil {
		return nil, err
	}

	// handle status 200 with error
	if result.ID == 0 {
		resp.Body = ioutil.NopCloser(bytes.NewReader(buf.Bytes()))
		resp.StatusCode = http.StatusTeapot
		return nil, newErrorFromResponse(resp)
	}

	return &result, nil
}

// CancelAllOrders if a pair is nil cancels all orders on an account, else cancels all orders for the specified pair
// Docs https://www.bitstamp.net/api/#cancel-all-orders
func (h *HTTPAPI) CancelAllOrders(ctx context.Context, p *Pair) (*CancelAllOrdersResponse, error) {
	u := cancelAllOrdersURL
	if p != nil {
		u += p.String() + "/"
	}

	resp, err := h.doRequest(ctx, http.MethodPost, u, nil, true)
	if err != nil {
		return nil, err
	}

	var result CancelAllOrdersResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateBuyLimitOrder creates a buy limit order
// Docs https://www.bitstamp.net/api/#buy-order
func (h *HTTPAPI) CreateBuyLimitOrder(ctx context.Context, p Pair, r CreateBuyLimitOrderRequest) (*CreateBuyLimitOrderResponse, error) {
	params := url.Values{}
	if err := schema.NewEncoder().Encode(r, params); err != nil {
		return nil, err
	}

	resp, err := h.doRequest(ctx, http.MethodPost, fmt.Sprintf(buyLimitOrderURL, p), bytes.NewBuffer([]byte(params.Encode())), true)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	teeReader := io.TeeReader(resp.Body, &buf)

	var result CreateBuyLimitOrderResponse
	if err := json.NewDecoder(teeReader).Decode(&result); err != nil {
		return nil, err
	}

	// handle status 200 with error
	if result.ID == "" {
		resp.Body = ioutil.NopCloser(bytes.NewReader(buf.Bytes()))
		resp.StatusCode = http.StatusTeapot
		return nil, newErrorFromResponse(resp)
	}

	return &result, nil
}

// CreateSellLimitOrder creates a sell limit order
// Docs https://www.bitstamp.net/api/#sell-order
func (h *HTTPAPI) CreateSellLimitOrder(ctx context.Context, p Pair, r CreateSellLimitOrderRequest) (*CreateSellLimitOrderResponse, error) {
	params := url.Values{}
	if err := schema.NewEncoder().Encode(r, params); err != nil {
		return nil, err
	}

	resp, err := h.doRequest(ctx, http.MethodPost, fmt.Sprintf(sellLimitOrderURL, p), bytes.NewBuffer([]byte(params.Encode())), true)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	teeReader := io.TeeReader(resp.Body, &buf)

	var result CreateSellLimitOrderResponse
	if err := json.NewDecoder(teeReader).Decode(&result); err != nil {
		return nil, err
	}

	// handle status 200 with error
	if result.ID == "" {
		resp.Body = ioutil.NopCloser(bytes.NewReader(buf.Bytes()))
		resp.StatusCode = http.StatusTeapot
		return nil, newErrorFromResponse(resp)
	}

	return &result, nil
}

// GetWebsocketsToken retrieves a token that can be used for subscribing to private WebSocket channels.
// Docs https://www.bitstamp.net/api/#websockets-token
// For private Websocket access, you need to contact support.
func (h *HTTPAPI) GetWebsocketsToken(ctx context.Context) (*GetWebsocketTokenResponse, error) {
	resp, err := h.doRequest(ctx, http.MethodPost, websocketsTokenURL, nil, true)
	if err != nil {
		return nil, err
	}

	var result GetWebsocketTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// newRequest creates a http request that can be used to call public APIs
func (h *HTTPAPI) newRequest(ctx context.Context, method string, uri string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s%s", h.baseURL, uri), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	return req, nil
}

// newAuthenticatedRequest creates a http request that can be used to call private APIs
// Docs https://www.bitstamp.net/api/#api-authentication
func (h *HTTPAPI) newAuthenticatedRequest(ctx context.Context, method string, uri string, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(fmt.Sprintf("%s%s", h.baseURL, uri))
	if err != nil {
		return nil, fmt.Errorf("failed to parse url `%s`", u)
	}

	var bodyBytes []byte
	if body != nil {
		bodyBytes, err = ioutil.ReadAll(body)
		if err != nil {
			return nil, fmt.Errorf("failed to read request body, %w", err)
		}
		body = bytes.NewBuffer(bodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Auth", "BITSTAMP"+" "+h.key)
	req.Header.Set("X-Auth-Nonce", uuid.NewString())
	req.Header.Set("X-Auth-Timestamp", fmt.Sprintf("%d", time.Now().UTC().UnixMilli()))
	req.Header.Set("X-Auth-Version", "v2")
	req.Header.Set("Accept", "application/json")
	if bodyBytes != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	// Sign request
	rq := u.RawQuery
	if rq != "" {
		rq = "?" + rq
	}
	signatureString := fmt.Sprintf(
		"BITSTAMP %s%s%s%s%s%s%s%s%s%s",
		h.key, method, u.Host, u.Path, rq, req.Header.Get("Content-Type"), req.Header.Get("X-Auth-Nonce"), req.Header.Get("X-Auth-Timestamp"), "v2", string(bodyBytes))

	mac := hmac.New(sha256.New, []byte(h.secret))
	if _, err := mac.Write([]byte(signatureString)); err != nil {
		return nil, fmt.Errorf("failed to sign request. %w", err)
	}
	req.Header.Set("X-Auth-Signature", hex.EncodeToString(mac.Sum(nil)))

	return req, nil
}

func (h *HTTPAPI) doRequest(ctx context.Context, method string, uri string, payload io.Reader, private bool) (*http.Response, error) {
	var (
		req    *http.Request
		reqErr error
	)

	if private {
		req, reqErr = h.newAuthenticatedRequest(ctx, method, uri, payload)
		if reqErr != nil {
			return nil, reqErr
		}
	} else {
		req, reqErr = h.newRequest(ctx, method, uri, payload)
		if reqErr != nil {
			return nil, reqErr
		}
	}

	if h.debug {
		outgoing, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(outgoing))
	}

	resp, err := h.handle.Do(req)
	if err != nil {
		return nil, err
	}

	if h.debug {
		incoming, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(incoming))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, newErrorFromResponse(resp)
	}

	return resp, nil
}
