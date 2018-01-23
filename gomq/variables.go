package gomq

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var client MQTT.Client

const (
	broker = "ssl://mqtt.ardich.com:8883"
)

// QueueJSON is
type QueueJSON struct {
	Body []struct {
		Params []interface{} `json:"params"`
		Type   string        `json:"type"`
	} `json:"body"`
	Header struct {
		MsgID          string `json:"msgId"`
		MaxMessageSize int    `json:"maxMessageSize"`
	} `json:"header"`
}
