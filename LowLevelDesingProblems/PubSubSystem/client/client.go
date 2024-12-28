package client

import (
	"bufio"
	"context"
	"google.golang.org/protobuf/proto"
	"log/slog"
	"net"
	"pubsub/commons"
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
	command := &commons.Command{
		Type:  commons.Command_SUBSCRIBE,
		Topic: topic,
	}

	var encodedCmd []byte
	if err := commons.Marshal(&encodedCmd, command); err != nil {
		slog.Error("could not create command message")
		close(msgChan)
		return msgChan, err
	}

	if _, err := c.conn.Write(encodedCmd); err != nil {
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
				var rawMsg []byte
				if err := commons.ReadMsg(reader, &rawMsg); err != nil {
					slog.Error("could not read message from server")
					return
				}

				var msg commons.Message
				if err := proto.Unmarshal(rawMsg, &msg); err != nil {
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
	command := &commons.Command{
		Type:  commons.Command_UNSUBSCRIBE,
		Topic: topic,
	}

	var encodedCmd []byte
	if err := commons.Marshal(&encodedCmd, command); err != nil {
		slog.Error("could not create command message")
		return false
	}

	if _, err := c.conn.Write(encodedCmd); err != nil {
		slog.Error("could not unsubscribe", "topic", topic)
		return false
	}

	c.ctx.Done()
	return true
}

func (c *Client) Publish(topic string, msg *commons.Message) bool {
	command := &commons.Command{
		Type:  commons.Command_PUBLISH,
		Topic: topic,
		Msg:   msg,
	}

	var encodedCmd []byte
	if err := commons.Marshal(&encodedCmd, command); err != nil {
		slog.Error("could not create command message")
		return false
	}

	if _, err := c.conn.Write(encodedCmd); err != nil {
		slog.Error("could not publish message")
		return false
	}

	return true
}
