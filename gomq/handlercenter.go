package gomq

import (
	"encoding/json"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	log.Printf("TOPIC: %s\n", msg.Topic())
	log.Printf("MSG: %s\n", msg.Payload())

	var queJSONVariable QueueJSON
	json.Unmarshal(msg.Payload(), &queJSONVariable)
	parseData(queJSONVariable.Body[0].Type)
}

func parseData(data string) {
	log.Println("Get Operation : ", data)
}
