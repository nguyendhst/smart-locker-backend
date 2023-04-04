package alert

import (
	"context"
	"fmt"
	"log"
	"smart-locker/backend/fcm"
	"strconv"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	swagger "github.com/nguyendhst/adafruit-go-client-v2"

	firebase "firebase.google.com/go/v4"
)

var (
	// temporary
	feedkeys      = []string{"locker1-temperature", "locker2-temperature"}
	ErrorAnalysis = fmt.Errorf("Analysis Error")
)

type (
	Alert struct {
		FirebaseApp *firebase.App
		tempChan    chan float32
	}
)

func NewAlert() (*Alert, error) {
	app, err := fcm.NewClient()
	if err != nil {
		return nil, err
	}
	return &Alert{
		FirebaseApp: app,
		tempChan:    make(chan float32),
	}, nil
}

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

	// check if no temp is above 65
	//for _, v := range values {
	//	if v > 65 {
	//		m.SendTempAlert(float32(v))
	//		return
	//	}
	//}

	//// detect sudden change in temperature
	//mean, err := stats.Mean(values)
	//if err != nil {
	//	log.Println(ErrorAnalysis, err)
	//}
	//fmt.Printf("Mean temperature: %f\n", mean)

	//stdDev, err := stats.StandardDeviationPopulation(values)
	//if err != nil {
	//	log.Println(ErrorAnalysis, err)
	//}

	//// find outliers (https://en.wikipedia.org/wiki/Chauvenet%27s_criterion)
	//{
	//	p := 1 - 1/float64(len(values)) // p is probability represented by one tail of the normal distribution
	//	dmax := 1.0 / (2 * float64(len(values))) * (1 - p)
	//}

}

func TemperatureHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", msg.Topic(), msg.Payload())
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
