package subscriber

type Subscriber struct {
	name string `json:"name, omotempty"`
	topic string `json:"topic,omitempty"`
	Url string `json:"url"`
	TimeoutInMs int `json:"timeout_in_ms"`
}
