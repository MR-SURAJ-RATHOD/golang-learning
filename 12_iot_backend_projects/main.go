package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// round2 keeps only 2 digits after decimal (e.g. 230.2222 → 230.22)
func round2(v float64) float64 {
	return math.Round(v*100) / 100
}

const (
	broker = "tcp://broker.hivemq.com:1883"
	topic  = "golang-learning/sensor/temp"
)

type SensorReading struct {
	Temp      float64
	Timestamp int64
}

func main() {
	fmt.Println("=== 12 IoT Backend ===")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	readings := make(chan SensorReading, 100)
	var wg sync.WaitGroup

	// Lesson 1: Connect to MQTT broker
	client := connectMQTT()
	defer client.Disconnect(250)

	// Lesson 2: Worker pool processes incoming sensor data
	numWorkers := 3
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go func(id int) {
			defer wg.Done()
			processor(ctx, id, readings)
		}(i)
	}

	// Lesson 3: Subscribe and push to channel
	client.Subscribe(topic, 0, func(_ mqtt.Client, msg mqtt.Message) {
		var temp float64
		fmt.Sscanf(string(msg.Payload()), `{"temp": %f`, &temp)
		temp = round2(temp)
		select {
		case readings <- SensorReading{Temp: temp, Timestamp: time.Now().Unix()}:
		case <-ctx.Done():
		}
	})

	// Lesson 4: Simulate sensor publisher
	go publishSensorData(client, ctx)

	fmt.Println("Running for 15 seconds... (Ctrl+C to stop early)")
	wg.Wait()
	fmt.Println("Done.")
}

func connectMQTT() mqtt.Client {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID("go_iot_backend")
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	fmt.Println("✅ Connected to MQTT broker")
	return client
}

func publishSensorData(client mqtt.Client, ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			temp := round2(20.0 + rand.Float64()*10.0)
			payload := fmt.Sprintf(`{"temp": %.2f, "ts": %d}`, temp, time.Now().Unix())
			client.Publish(topic, 0, false, payload)
			fmt.Printf("📡 Sensor published: %.2f°C\n", temp)
		}
	}
}

func processor(ctx context.Context, id int, readings <-chan SensorReading) {
	for {
		select {
		case <-ctx.Done():
			return
		case r, ok := <-readings:
			if !ok {
				return
			}
			status := "OK"
			if r.Temp > 28.0 {
				status = "⚠️ HIGH TEMP ALERT"
			}
			fmt.Printf("  Worker %d: %.2f°C — %s\n", id, r.Temp, status)
		}
	}
}
