package server

import (
	"bufio"
	"log/slog"
	"net"
	"regexp"
	"strings"
)

const (
	ADDR               string = ":8080"
	SUBSCRIBE_PREFIX   string = "SUBSCRIBE"
	UNSUBSCRIBE_PREFIX string = "UNSUBSCRIBE"
	PUBLISH_PREFIX     string = "PUBLISH"
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
		msg, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				slog.Info("client disconnected", "addr", addr.String())
			} else {
				slog.Error("could not read message from client", "addr", addr.String(), "error", err)
			}
			conn.Close()
			return
		}

		if strings.HasPrefix(msg, SUBSCRIBE_PREFIX) {
			slog.Info("received subscribe message", "msg", msg)
			topic := regexp.MustCompile(`\w+\((.*?)\)`).FindStringSubmatch(msg)[1]
			ps.handleSubscribe(conn, topic)
		} else if strings.HasPrefix(msg, UNSUBSCRIBE_PREFIX) {
			slog.Info("received unsubscribe message", "msg", msg)
			topic := regexp.MustCompile(`\w+\((.*?)\)`).FindStringSubmatch(msg)[1]
			ps.handleUnsubscribe(conn, topic)
		} else if strings.HasPrefix(msg, PUBLISH_PREFIX) {
			slog.Info("received publish message", "msg", msg)
			matches := regexp.MustCompile(`\w+\((.*?)\): (.*)`).FindStringSubmatch(msg)
			topic := matches[1]
			clientMsg := matches[2]

			ps.handlePublish(topic, []byte(clientMsg))
		} else {
			slog.Info("unknown message", "msg", msg)
		}

	}

}

func (ps *PubSub) handleSubscribe(conn net.Conn, topic string) bool {
	t, ok := ps.topics[topic]
	if !ok {
		slog.Error("topic does not exist", "topic", topic)
		return false
	}

	t.AddSubscriber(conn)
	return true
}

func (ps *PubSub) handleUnsubscribe(conn net.Conn, topic string) bool {
	t, ok := ps.topics[topic]
	if !ok {
		slog.Error("topic does not exist", "topic", topic)
		return false
	}

	t.RemoveSubscriber(conn)
	return true

}

func (ps *PubSub) handlePublish(topic string, msg []byte) {
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
