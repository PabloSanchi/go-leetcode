package main

import "pubsub/server"

func main() {
	srv := server.NewPubSub()
	srv.Start()
}
