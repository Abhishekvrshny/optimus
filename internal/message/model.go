package message

type Message struct {
	Body string `json:"body"`
	Headers map[string][]string `json:"headers"`
	Id string `json:"id"`
}
