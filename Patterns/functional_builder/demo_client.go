package main

import "net"

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:3000")
	if err != nil {
		panic(err)
	}

	if _, err := conn.Write([]byte("hello")); err != nil {
		panic(err)
	}

	conn.Close()
}
