package mqttclient

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//func AllFeedsCallback(client mqtt.Client, msg mqtt.Message) {
//	fmt.Printf("Received message on topic: %s\nMessage: %s\n", msg.Topic(), msg.Payload())
//}

func SensorFeedCallback(client mqtt.Client, msg mqtt.Message) {
	// mux topic to more specific handler
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", msg.Topic(), msg.Payload())

}

//func TemperatureFeedCallback(client mqtt.Client, msg mqtt.Message) {
//	fmt.Printf("Received message on topic: %s\nMessage: %s\n", msg.Topic(), msg.Payload())
//}

func MoistureFeedCallback(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", msg.Topic(), msg.Payload())
}

func LockFeedCallback(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", msg.Topic(), msg.Payload())
}
