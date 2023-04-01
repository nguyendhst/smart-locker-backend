package fcm

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
)

func NewClient() (*firebase.App, error) {
	//opt := option.WithCredentialsFile("/Users/nguyen/study/da/backend/go-app-357215-firebase-adminsdk-oj7ue-9b6150c91b.json")
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}

func SendAlert(ctx context.Context, client *messaging.Client, topic, temp string) {
	// [START send_to_topic_golang]
	// The topic name can be optionally prefixed with "/topics/".

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Temperature Alert",
			Body:  "Temperature is " + temp,
		},
		Data: map[string]string{
			topic: temp,
		},
		Topic: topic,
	}

	// Send a message to the devices subscribed to the provided topic.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
	// [END send_to_topic_golang]
}
