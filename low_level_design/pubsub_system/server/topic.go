package server

import (
	"log/slog"
	"net"
	"sync"
)

type Topic struct {
	id          string
	subscribers map[net.Conn]struct{} // empty structs occupies 0 bytes

	l sync.RWMutex
}

func NewTopic(id string) *Topic {
	return &Topic{
		id:          id,
		subscribers: make(map[net.Conn]struct{}),
	}
}

func (t *Topic) AddSubscriber(conn net.Conn) {
	t.l.Lock()
	defer t.l.Unlock()
	t.subscribers[conn] = struct{}{}
}

func (t *Topic) RemoveSubscriber(conn net.Conn) {
	t.l.Lock()
	defer t.l.Unlock()
	delete(t.subscribers, conn)
}

func (t *Topic) Broadcast(msg []byte) {
	t.l.RLock()
	defer t.l.RUnlock()

	msg = append(msg, byte('\n'))

	for conn := range t.subscribers {
		_, err := conn.Write(msg)
		if err != nil {
			slog.Error("could not send message to subscriber")
		}
	}
}
