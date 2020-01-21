package topic

type Topic struct {
	name string `json:"name, omitempty"`
	TimeoutInMs int `json:"timeout_in_ms"`
	IsDurable bool `json:"is_durable"`
}
