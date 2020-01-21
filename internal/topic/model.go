package topic

type CreateTopicRequest struct {
	TimeoutInMs int `json:"timeout_in_ms"`
	IsDurable bool `json:"is_durable"`
}
