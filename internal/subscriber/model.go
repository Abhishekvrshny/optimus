package subscriber

type Subscriber struct {
	topic string `json:"topic,omitempty"`
	url string `json:"url"`
	timeoutInMs int `json:"timeout_in_ms"`
}
