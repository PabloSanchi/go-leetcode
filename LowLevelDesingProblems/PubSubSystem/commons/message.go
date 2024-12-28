package commons

type Message struct {
	Content string `json:"content"`
}

func NewMessage(content string) Message {
	return Message{Content: content}
}

func (m *Message) String() string {
	return m.Content
}
