package bitstamp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	ErrAlreadySubscribed        = errors.New("you can subscribe once per client instance")
	ErrUnableToParseMessage     = errors.New("failed to parse websocket message")
	ErrReceivedReconnectMessage = errors.New("Bitstamp requested to reconnect")
	ErrReadMessage              = errors.New("failed to read message")
	ErrWriteMessage             = errors.New("failed to write message")
)

type WebsocketMessage struct {
	Message    interface{}
	RawMessage []byte
	Error      error
}

type WebsocketAPI struct {
	conn        *websocket.Conn
	address     string
	channelSubs sync.Map
}

func NewWebsocketAPI(opts ...wsOption) (*WebsocketAPI, error) {
	w := WebsocketAPI{
		address:     "wss://ws.bitstamp.net",
		channelSubs: sync.Map{},
	}

	// override defaults via available functional options
	for i := range opts {
		opts[i](&w)
	}

	c, _, err := websocket.DefaultDialer.Dial(w.address, nil)
	if err != nil {
		return nil, err
	}
	w.conn = c

	return &w, nil
}

// Consume subscribe to channel(s) and start consuming messages
// do not call this on same instance twice
func (w *WebsocketAPI) Consume(ctx context.Context, channels ...Channel) (<-chan WebsocketMessage, error) {
	messages := make(chan WebsocketMessage)
	wsMessages := make(chan []byte)

	go func() {
		for {
			msg, err := w.readMessage(ctx)
			if err != nil {
				messages <- WebsocketMessage{Error: err}
				close(wsMessages)
				return
			}

			wsMessages <- msg
		}
	}()

	go func() {
		defer close(messages)

		for {
			select {
			case <-ctx.Done():
				return

			case m, ok := <-wsMessages:
				if !ok {
					return
				}

				msg, err := parseMessage(m)
				if err != nil {
					messages <- WebsocketMessage{
						Error:      fmt.Errorf("unable to parse message, %w", err),
						RawMessage: m,
					}

					continue
				}

				messages <- WebsocketMessage{
					Message:    msg,
					RawMessage: m,
				}
			}
		}
	}()

	if err := w.SubscribeToChannels(ctx, channels...); err != nil {
		return nil, err
	}

	return messages, nil
}

// SubscribeToChannels use this method to subscribe to channel(s)
func (w *WebsocketAPI) SubscribeToChannels(ctx context.Context, channels ...Channel) error {
	for i := range channels {
		m := fmt.Sprintf(`{"event":"bts:subscribe","data":{"channel": "%s"}}`, channels[i].String())

		if err := w.writeMessage(ctx, []byte(m)); err != nil {
			return fmt.Errorf("failed to subscribe to channel %s, %w", channels[i].String(), err)
		}

		w.channelSubs.Store(channels[i], struct{}{})
	}

	return nil
}

// UnSubscribe use this method to unsubscribe from channel(s)
func (w *WebsocketAPI) UnSubscribeFromChannels(ctx context.Context, channels ...Channel) error {
	for i := range channels {
		m := fmt.Sprintf(`{"event":"bts:unsubscribe","data":{"channel": "%s"}}`, channels[i].String())

		if err := w.conn.WriteMessage(websocket.TextMessage, []byte(m)); err != nil {
			return fmt.Errorf("failed to unsubscribe from channel %s, %w", channels[i].String(), err)
		}

		w.channelSubs.Delete(channels[i])
	}

	return nil
}

// GetSubscriptions returns a list of tracked subscriptions
func (w *WebsocketAPI) GetSubscriptions() []Channel {
	var channels []Channel

	w.channelSubs.Range(func(key, _ interface{}) bool {
		k, _ := key.(Channel)
		channels = append(channels, k)

		return true
	})

	return channels
}

func (w *WebsocketAPI) readMessage(ctx context.Context) ([]byte, error) {
	_, msg, err := w.conn.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("%w, %s", ErrReadMessage, err)
	}

	return msg, nil
}

func (w *WebsocketAPI) writeMessage(ctx context.Context, m []byte) error {
	if err := w.conn.WriteMessage(websocket.TextMessage, []byte(m)); err != nil {
		return fmt.Errorf("%w, %s", ErrWriteMessage, err)
	}

	return nil
}

// Close connection
func (w *WebsocketAPI) Close() error {
	return w.conn.Close()
}

func parseMessage(m []byte) (interface{}, error) {
	if m == nil {
		return nil, nil
	}

	var wsMsg WebSocketMessage
	if err := json.Unmarshal(m, &wsMsg); err != nil {
		return nil, fmt.Errorf("%w, %s", ErrUnableToParseMessage, err)
	}
	if wsMsg.Event == "bts:request_reconnect" {
		return nil, ErrReceivedReconnectMessage
	}

	var msg interface{}

	switch {
	case strings.HasPrefix(wsMsg.Channel, "live_trades_") && wsMsg.Event == "trade":
		var channelMSG LiveTickerChannel
		if err := json.Unmarshal(m, &channelMSG); err != nil {
			return nil, fmt.Errorf("%w, %s", ErrUnableToParseMessage, err)
		}
		msg = channelMSG

	case strings.HasPrefix(wsMsg.Channel, "live_orders_") &&
		(wsMsg.Event == "order_created" || wsMsg.Event == "order_changed " || wsMsg.Event == "order_deleted"):
		var channelMSG LiveOrdersChannel
		if err := json.Unmarshal(m, &channelMSG); err != nil {
			return nil, fmt.Errorf("%w, %s", ErrUnableToParseMessage, err)
		}
		msg = channelMSG

	case strings.HasPrefix(wsMsg.Channel, "order_book_") && wsMsg.Event == "data":
		var lobc LiveOrderBookChannel
		if err := json.Unmarshal(m, &lobc); err != nil {
			return nil, fmt.Errorf("%w, %s", ErrUnableToParseMessage, err)
		}
		msg = lobc

	case strings.HasPrefix(wsMsg.Channel, "detail_order_book_") && wsMsg.Event == "data":
		var ldobc LiveDetailOrderBookChannel
		if err := json.Unmarshal(m, &ldobc); err != nil {
			return nil, fmt.Errorf("%w, %s", ErrUnableToParseMessage, err)
		}
		msg = ldobc

	case strings.HasPrefix(wsMsg.Channel, "diff_order_book_") && wsMsg.Event == "data":
		var lfobc LiveFullOrderBook
		if err := json.Unmarshal(m, &lfobc); err != nil {
			return nil, fmt.Errorf("%w, %s", ErrUnableToParseMessage, err)
		}
		msg = lfobc

	default:
		msg = wsMsg
	}

	return msg, nil
}
