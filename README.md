![Tests](https://github.com/georlav/bitstamp/workflows/Tests/badge.svg)
![Golang-CI](https://github.com/georlav/bitstamp/workflows/Golang-CI/badge.svg)

# Bitstamp
An extensive client implementation of the [Bitstamp API](https://www.bitstamp.net/api/) using Go.

## Examples
Please check the [examples_test.go](examples_test.go) file for some basic usage examples.

## Endpoint Support

  ### HTTP API

  * **Public data functions**
    * [x] Ticker
    * [x] Hourly ticker
    * [x] Order book
    * [x] Transactions
    * [x] Trading pairs info
    * [x] OHLC data
    * [x] EUR/USD conversion rate
  * **Private functions**
    * [x] Account balance
    * [x] User transactions
    * [x] Crypto transactions
    * [x] Orders
      * [x] Open orders
      * [x] Order status
      * [x] Cancel order
      * [x] Cancel all orders
      * [ ] Buy limit order
      * [ ] Buy market order
      * [ ] Buy instant order
      * [x] Sell limit order
      * [ ] Sell market order
      * [ ] Sell instant order
    * [ ] Withdrawal requests
    * [ ] Crypto withdrawals
    * [ ] Crypto deposits
    * [ ] Transfer balance from Sub to Main Account
    * [ ] Transfer balance from Main to Sub Account
    * [ ] Open bank withdrawal
    * [ ] Bank withdrawal status
    * [ ] Cancel bank withdrawal
    * [ ] New liquidation address
    * [ ] Liquidation address info
    * [x] WebSockets token

   ### Websocket API v2

  * **Public channels**
    * [x] Live ticker
    * [x] Live orders
    * [x] Live order book
    * [x] Live detail order book
    * [x] Live full order book
  * **Private Channels**
    * [ ] Live ticker
    * [ ] Live orders

New pairs are constantly added so if you notice that a pair is missing you can run the following command and open pr

```bash
go generate ./...
```

This fetches all supported pairs from bitstamp and generates new pair and channel enums.

## Private functions and configuration
To be able to use private functions you need to generate an API key and a secret using your bitstamp account. To do that you need to visit.

Profile settings -> API access -> New API key

You can pass those to the client using the following functional options
```go
c := bitstamp.NewHTTPAPI(
	bitstamp.APIKeyOption("yourkey"),
	bitstamp.APISecretOption("yoursecret"),
)
```

Or you can set them as environmental variables and client will automatically read them
```bash
export BITSTAMP_KEY="yourkey"
export BITSTAMP_SECRET="yoursecret"
```

> **IMPORTANT:** Environmental variables override functional options.

## Running tests
To run the integration tests for public functions use
```go
go test ./... -v -race -short
```

To run all integration tests and cover also private functions you need to set your secret and key at your env and run test without -short
```go
go test ./... -v -race
```

## License
Distributed under the MIT License. See `LICENSE` for more information.

## Contact
George Lavdanis - georlav@gmail.com

Project Link: [https://github.com/georlav/bitstamp](https://github.com/georlav/bitstamp)



