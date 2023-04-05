package alert

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"smart-locker/backend/fcm"
	"strconv"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	swagger "github.com/nguyendhst/adafruit-go-client-v2"

	firebase "firebase.google.com/go/v4"
)

const (
	TEMP_THRESHOLD = 65.0
	MOIS_THRESHOLD = 0.0
)

var (
	ErrorAnalysis = fmt.Errorf("Analysis Error")

	// global alerter
	Alerter *Alert
)

type (
	Alert struct {
		FirebaseApp *firebase.App
		tempChan    chan float32
	}

	Payload struct {
		Id        uint64 `json:"id"`
		LastValue string `json:"last_value"`
		UpdatedAt string `json:"updated_at"`
		Key       string `json:"key"`
		Data      Data   `json:"data"`
	}

	Data struct {
		CreatedAt string `json:"created_at"`
		Value     string `json:"value"`
		Location  string `json:"location"`
		Id        string `json:"id"`
	}
)

func NewAlert() error {
	app, err := fcm.NewClient()
	if err != nil {
		return err
	}
	Alerter = &Alert{
		FirebaseApp: app,
		tempChan:    make(chan float32),
	}

	return nil
}

//func (p *Payload) UnmarshalJSON(b []byte) error {
//	type Alias Payload
//	aux := &struct {
//		*Alias
//	}{
//		Alias: (*Alias)(p),
//	}
//	if err := json.Unmarshal(b, &aux); err != nil {
//		return err
//	}
//	return nil
//}

func (m *Alert) Start(ctx context.Context, cfg *swagger.Configuration) {
	go m.startTempAlert(ctx)
}

func (m *Alert) AnalyzeTemp(temp [][]string) {
	// format: [[time:val], [time:val], ...]
	values := make([]float64, 0)
	for _, v := range temp {
		val, err := strconv.ParseFloat(v[1], 64)
		if err != nil {
			log.Println("Parse error", err)
		} else {
			values = append(values, (val))
		}
	}

	fmt.Printf("Recorded temperature: %f\n", values)
}

func ParsePayload(data []byte) (*Payload, error) {
	var payload = Payload{}
	err := json.Unmarshal(data, &payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

func TemperatureFeedCallback(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", msg.Topic(), msg.Payload())

	// if temp is above threshold, send alert
	// parse payload

	//payload, err := json.Marshal(msg.Payload())
	//if err != nil {
	//	log.Println("Error parsing payload", err)
	//}

	parsed, err := ParsePayload(msg.Payload())
	if err != nil {
		log.Println("Error parsing payload", err)
	}
	// field is last_value
	lastValue := parsed.LastValue
	fmt.Println("Last value:", lastValue)

	// try to convert to float64
	val, err := strconv.ParseFloat(lastValue, 64)
	if err != nil {
		log.Println("Error parsing payload", err)
	}

	// send alert
	if val > TEMP_THRESHOLD {
		Alerter.SendTempAlert(float32(val))
	}

}

func (m *Alert) startTempAlert(ctx context.Context) {
	for temp := range m.tempChan {
		client, err := m.FirebaseApp.Messaging(ctx)
		if err != nil {
			log.Fatalln(err)
		}
		fcm.SendAlert(ctx, client, "alert", fmt.Sprintf("%f", temp))
	}
}

func (m *Alert) SendTempAlert(temp float32) {
	m.tempChan <- temp
}

func (m *Alert) Close() {
	close(m.tempChan)
}
