package main

import (
	"log/slog"
	"pubsub/client"
)

func main() {
	cl := client.NewClient("localhost:8080")
	ok := cl.Connect()
	if !ok {
		slog.Error("failed to connect to server")
	}
	defer cl.Disconnect()

	dataChan, err := cl.Subscribe("default")
	if err != nil {
		slog.Error("failed to subscribe to channel")
	}

	for data := range dataChan {
		slog.Info("message received", "data", data.String())
	}

}
