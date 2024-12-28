package main

import (
	"log/slog"
	"pubsub/client"
	"pubsub/commons"
)

func main() {
	cl := client.NewClient("localhost:8080")
	ok := cl.Connect()
	if !ok {
		slog.Error("failed to connect to server")
	}
	defer cl.Disconnect()

	topic := "default"
	msg := &commons.Message{
		Content: "hello from producer!",
	}

	sent := cl.Publish(topic, msg)
	if !sent {
		slog.Error("failed to publish message")
	}

	slog.Info("message published successfully", "topic", topic, "message", msg)
}
