package events

type Message struct {
	EventId         string      `json:"event_id"`
	EventName       string      `json:"event_name"`
	EventDataFormat string      `json:"event_data_format"`
	Type            string      `json:"type"`
	Timestamp       string      `json:"timestamp"`
	Version         string      `json:"version"`
	Country         string      `json:"country"`
	Origin          string      `json:"origin"`
	Payload         interface{} `json:"payload"`
}
