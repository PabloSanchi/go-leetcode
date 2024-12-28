package client

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net"
	"pubsub/commons"
	"strings"
)

type Client struct {
	addr string
	conn net.Conn
	ctx  context.Context
}

func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
		conn: nil,
		ctx:  context.Background(),
	}
}

func (c *Client) Connect() bool {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return false
	}

	c.conn = conn
	return true
}

func (c *Client) Disconnect() {
	c.conn.Close()
}

func (c *Client) Subscribe(topic string) (<-chan commons.Message, error) {
	msgChan := make(chan commons.Message)

	_, err := c.conn.Write([]byte("SUBSCRIBE(" + topic + ")\n"))
	if err != nil {
		slog.Error("could not subscribe", "topic", topic)
		close(msgChan)
		return msgChan, err
	}

	go func() {
		defer close(msgChan)
		reader := bufio.NewReader(c.conn)
		for {
			select {
			case <-c.ctx.Done():
				slog.Info("client unsubscribed, closing data channel")
				return
			default:
				rawMsg, err := reader.ReadString('\n')
				if err != nil {
					slog.Error("could not read message from server")
					return
				}
				var msg commons.Message
				err = json.Unmarshal([]byte(strings.TrimSpace(rawMsg)), &msg)
				if err != nil {
					slog.Error("error deserializing topic message", err)
					continue
				}

				msgChan <- msg
			}
		}

	}()

	return msgChan, nil
}

func (c *Client) Unsubscribe(topic string) bool {
	_, err := c.conn.Write([]byte("UNSUBSCRIBE(" + topic + ")\n"))
	if err != nil {
		fmt.Println(err)
		return false
	}

	c.ctx.Done()
	return true
}

func (c *Client) Publish(topic string, msg commons.Message) bool {
	decodedMsg, err := json.Marshal(msg)
	if err != nil {
		slog.Error("could not serialize message")
		return false
	}

	message := fmt.Sprintf("PUBLISH(%s): %s\n", topic, string(decodedMsg))
	_, err = c.conn.Write([]byte(message))
	if err != nil {
		slog.Error("could not publish message")
		return false
	}

	return true
}
