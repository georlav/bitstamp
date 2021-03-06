package bitstamp_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/georlav/bitstamp"
)

func ExampleHTTPAPI() {
	c := bitstamp.NewHTTPAPI(
		bitstamp.EnableDebugOption(),                       // Show requests and responses for debugging purposes
		bitstamp.BaseURLOption("https://www.bitstamp.net"), // Change default API URL
		bitstamp.APIKeyOption("xxx"),                       // Can also be set via env variable BITSTAMP_KEY
		bitstamp.APISecretOption("xxx"),                    // Can also be set via env variable BITSTAMP_SECRET
		bitstamp.ClientTimeoutOption(15),                   // Change default client timeout
	)

	resp, err := c.GetTicker(context.Background(), bitstamp.BTCEUR)
	if err != nil {
		log.Fatalf("Failed to retrieve data, %s", err)
	}

	log.Println(resp.Ask, resp.Bid)
}

func ExampleWebsocketAPI() {
	// initialize client
	ws, err := bitstamp.NewWebsocketAPI()
	if err != nil {
		log.Fatalf("failed to initialize websocket client, %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Hour)
	defer cancel()

	// subscribe to channels and start consuming messages
	// consume should be called once per instance
	msgCH, err := ws.Consume(ctx,
		// You can add your needed channels
		bitstamp.LiveTradesBTCEURChannel,
		bitstamp.LiveTradesETHEURChannel,
		bitstamp.LiveTradesZRXEURChannel,
		// There are also some helpers that can be used to retrieve group of channels to batch subscribe to them
		//
		// bitstamp.GetAllChannels()...
		// bitstamp.GetAllChannels()...
		// bitstamp.GetUSDChannels()...,
		// bitstamp.GetBTCChannels()...,
		// bitstamp.GetGBPChannels()...,
		// bitstamp.GetLiveTradeChannels()...,
		// bitstamp.GetLiveOrderChannels()...,
		// bitstamp.GetDetailOrderBookChannels()...,
		// bitstamp.GetDiffOrderBookChannels()...,
	)
	if err != nil {
		log.Fatalf("failed to subscribe to channels, %s", err)
	}

	// Consume messages
	for msg := range msgCH {
		if msg.Error != nil {
			if errors.Is(err, bitstamp.ErrReceivedReconnectMessage) {
				log.Fatal("websocket server is to go under maintenance, you need to reconnect.")
			}

			log.Fatal(msg.Error)
		}

		// you can work with raw messages
		fmt.Println("Raw Message: ", string(msg.RawMessage))

		// you can work with objects
		switch v := msg.Message.(type) {
		case bitstamp.LiveTickerChannel:
			fmt.Println("Message: ", v.Channel, v.Event, v.Data, v.Data.Amount)

		case bitstamp.LiveOrdersChannel:
			fmt.Println("Message: ", v.Channel, v.Event, v.Data, v.Data.ID)

		case bitstamp.LiveOrderBookChannel:
			fmt.Println("Message: ", v.Channel, v.Event, v.Data)

		case bitstamp.LiveDetailOrderBookChannel:
			fmt.Println("Message: ", v.Channel, v.Event, v.Data)

		case bitstamp.LiveFullOrderBook:
			fmt.Println("Message: ", v.Channel, v.Event, v.Data)

		case bitstamp.WebSocketMessage:
			fmt.Println("Event Message: ", v.Channel, v.Event, v.Data)

		default:
			fmt.Println("Unknown message: ", msg.RawMessage)
		}

	}

	// subscribe to more channels
	if err := ws.SubscribeToChannels(context.Background(), bitstamp.LiveTradesXLMEURChannel); err != nil {
		log.Fatalf("Failed to unsubscribe, %s", err)
	}

	// unsubscribe from channels
	if err := ws.UnSubscribeFromChannels(context.Background(), bitstamp.LiveTradesXLMEURChannel); err != nil {
		log.Fatalf("Failed to unsubscribe, %s", err)
	}

	// Get active subscriptions
	fmt.Println("Active subscriptions", ws.GetSubscriptions())
}
