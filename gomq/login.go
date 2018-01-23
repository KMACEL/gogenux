package gomq

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// Login is
type Login struct {
	SetClientID string
	SetUsername string
	SetPassword string
}

// Connect is
func Connect(l Login) MQTT.Client {
	opts := MQTT.NewClientOptions().AddBroker(broker)
	opts.SetClientID(l.SetClientID)
	opts.SetUsername(l.SetUsername)
	opts.SetPassword(l.SetPassword)
	opts.SetDefaultPublishHandler(f)

	client = MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	// TODO : Client i det cliente gönder. Her yer oradan alsın
	return client
}

func getClient() MQTT.Client {
	return client
}
