package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
	//import the Paho Go MQTT library

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

//define a function for the default message handler
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	log.Printf("TOPIC: %s\n", msg.Topic())
	log.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	connect()
	go sendMessage()

	var pressKey string
	fmt.Scan(&pressKey)

}

var client MQTT.Client

func connect() {
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker("ssl://mqtt.ardich.com:8883")
	opts.SetClientID("gopher@go")
	opts.SetUsername("gophergo")
	opts.SetPassword("12345678")
	opts.SetDefaultPublishHandler(f)

	//create and start a client using the above ClientOptions
	client = MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func sendMessage() {
	topic := "gopher@go/publish/DeviceProfile/goNode/goSensor"
	for i := 0; i < 5; i++ {
		payload := "{data:{sensorData:[{date:" + strconv.FormatInt(time.Now().Unix(), 10) + "000,values:[555]}],formatVersion:2}}"
		token := client.Publish(topic, 1, false, payload)
		token.Wait()
		time.Sleep(10 * time.Second)
	}
}
