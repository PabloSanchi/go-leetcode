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

	msg := commons.NewMessage("Hello from producer!")

	sent := cl.Publish("default", msg)
	if !sent {
		slog.Error("failed to publish message")
	}

	slog.Info("message published successfully", "topic", "default", "message", msg)
}
