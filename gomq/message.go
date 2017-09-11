package gomq

import (
	"strconv"
	"time"
)

// SendMessage is
func SendMessage() {
	topic := "gopher@go/publish/DeviceProfile/goNode/goSensor"
	for i := 0; i < 5; i++ {
		payload := "{data:{sensorData:[{date:" + strconv.FormatInt(time.Now().Unix(), 10) + "000,values:[555]}],formatVersion:2}}"
		token := client.Publish(topic, 1, false, payload)
		token.Wait()
		time.Sleep(10 * time.Second)
	}
}
