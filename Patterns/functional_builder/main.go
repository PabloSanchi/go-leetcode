package main

import (
	"log/slog"
	"net"
	"os"
	"strconv"
)

type Param func(*Config)

type Config struct {
	Protocol string
	Host     string
	Port     int
	Origins  []string
}

func (c *Config) Address() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}

func WithProtocol(p string) Param {
	return func(c *Config) {
		c.Protocol = p
	}
}

func WithHost(h string) Param {
	return func(c *Config) {
		c.Host = h
	}
}

func WithPort(p int) Param {
	return func(c *Config) {
		c.Port = p
	}
}

func WithOrigins(o []string) Param {
	return func(c *Config) {
		c.Origins = o
	}
}

// WithOrigin
// note: no need to check if the origin is nil as initially will take the defualt value
func WithOrigin(o string) Param {
	return func(c *Config) {
		c.Origins = append(c.Origins, o)
	}
}

func defaultConfig() *Config {
	return &Config{
		Protocol: "tcp",
		Host:     "localhost",
		Port:     8080,
		Origins:  []string{},
	}
}

type Server struct {
	config Config
}

func NewServer(params ...Param) *Server {
	opts := defaultConfig()
	for _, opt := range params {
		if opt != nil {
			opt(opts)
		}
	}

	return &Server{
		config: *opts,
	}
}

func (s *Server) Bootstrap() {
	listener, err := net.ListenPacket(s.config.Protocol, s.config.Address())
	if err != nil {
		slog.Error("could not start server", "error", err)
		os.Exit(1)
	}

	slog.Info("server started", "address", s.config.Address())

	for {
		buffer := make([]byte, 1024)
		n, addr, err := listener.ReadFrom(buffer)
		if err != nil {
			slog.Error("error reading packet", "error", err)
			continue
		}

		slog.Info("received packet", "client", addr, "data", string(buffer[:n]))
	}
}

func main() {
	server := NewServer(
		WithHost("127.0.0.1"),
		WithPort(3000),
		WithProtocol("udp"),
	)

	server.Bootstrap()
}
