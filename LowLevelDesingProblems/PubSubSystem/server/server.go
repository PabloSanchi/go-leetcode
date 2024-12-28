package server

import (
	"bufio"
	"errors"
	"google.golang.org/protobuf/proto"
	"io"
	"log/slog"
	"net"
	"pubsub/commons"
)

const (
	ADDR string = ":8080"
)

type ConfigParam func(config *Config)

type Config struct {
	addr   string
	topics map[string]*Topic
}

type PubSub struct {
	addr   string
	topics map[string]*Topic
}

func WithTopic(topic *Topic) ConfigParam {
	return func(config *Config) {
		if len(config.topics) == 0 {
			config.topics = make(map[string]*Topic)
		}

		config.topics[topic.id] = topic
	}
}

func WithTopics(topics []*Topic) ConfigParam {
	return func(config *Config) {
		topicSlice := make(map[string]*Topic)
		for _, topic := range topics {
			topicSlice[topic.id] = topic
		}
	}
}

func WithListenAddr(addr string) ConfigParam {
	return func(config *Config) {
		config.addr = addr
	}
}
func NewPubSub(configs ...ConfigParam) *PubSub {
	cfg := defaultConfig()
	for _, config := range configs {
		if config != nil {
			config(cfg)
		}
	}

	return &PubSub{
		addr:   cfg.addr,
		topics: cfg.topics,
	}
}

func (ps *PubSub) Start() {
	server, err := net.Listen("tcp", ps.addr)
	if err != nil {
		slog.Error("could not start pubsub server")
		panic(err)
	}

	slog.Info("listening on", "port", ADDR)

	for {
		conn, err := server.Accept()
		if err != nil {
			slog.Error("incoming connection failed")
		}

		go ps.handleConnection(conn)
	}
}

func (ps *PubSub) handleConnection(conn net.Conn) {
	addr := conn.RemoteAddr()
	reader := bufio.NewReader(conn)

	for {
		var command commons.Command
		if err := readAndUnmarshal(reader, &command); err != nil {
			if errors.Is(err, io.EOF) {
				slog.Info("client disconnected", "address", addr)
			} else {
				slog.Error("could not unmarshal client message", "error", err)
			}
			conn.Close()
			return
		}

		switch command.Type {
		case commons.Command_SUBSCRIBE:
			ps.handleSubscribe(conn, command.Topic)

		case commons.Command_UNSUBSCRIBE:
			ps.handleUnsubscribe(conn, command.Topic)

		case commons.Command_PUBLISH:
			var clientMsg []byte
			if err := commons.Marshal(&clientMsg, command.Msg); err != nil {
				slog.Error("could not create message", "error", err)
				continue
			}

			ps.handlePublish(command.Topic, clientMsg)

		default:
			slog.Warn("unknown command type or client disconnected", "type", command.Type, "address", addr)
		}

	}

}

func readAndUnmarshal(reader *bufio.Reader, command *commons.Command) error {
	var msg []byte
	if err := commons.ReadMsg(reader, &msg); err != nil {
		return err
	}

	if err := proto.Unmarshal(msg, command); err != nil {
		return err
	}

	return nil
}

func (ps *PubSub) handleSubscribe(conn net.Conn, topic string) bool {
	slog.Info("client subscribed", "topic", topic)
	t, ok := ps.topics[topic]
	if !ok {
		slog.Error("topic does not exist", "topic", topic)
		return false
	}

	t.AddSubscriber(conn)
	return true
}

func (ps *PubSub) handleUnsubscribe(conn net.Conn, topic string) bool {
	slog.Info("client unsubscribed", "topic", topic)
	t, ok := ps.topics[topic]
	if !ok {
		slog.Error("topic does not exist", "topic", topic)
		return false
	}

	t.RemoveSubscriber(conn)
	return true

}

func (ps *PubSub) handlePublish(topic string, msg []byte) {
	slog.Info("publishing message", "topic", topic)
	t, ok := ps.topics[topic]
	if !ok {
		slog.Error("topic does not exist", "topic", topic)
		return
	}

	t.Broadcast(msg)
}

func defaultConfig() *Config {
	topics := make(map[string]*Topic)
	defaultTopic := NewTopic("default")
	topics[defaultTopic.id] = defaultTopic

	return &Config{
		addr:   ADDR,
		topics: topics,
	}
}
