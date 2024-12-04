package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

var sum, count float64

func main() {
	// Connect to the NATS server
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		fmt.Printf("Error connecting to NATS: %v\n", err)
		return
	}
	defer nc.Close()

	// Create a JetStream context
	js, err := jetstream.New(nc)
	if err != nil {
		fmt.Printf("Error creating JetStream context: %v\n", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define the stream and subject
	stream := "SENSORS"
	subject := "sensors.pizzanini.temp"

	// Create an ephemeral consumer
	consumer, err := js.CreateConsumer(ctx, stream, jetstream.ConsumerConfig{
		AckPolicy:         jetstream.AckExplicitPolicy,
		FilterSubject:     subject,
		InactiveThreshold: 10 * time.Millisecond,
	})
	if err != nil {
		fmt.Printf("Error creating ephemeral consumer: %v\n", err)
		return
	}

	fmt.Println("Consuming messages with ephemeral consumer...")
	// Start consuming messages using the consumer
	subscription, err := consumer.Consume(func(msg jetstream.Msg) {
		value, err := strconv.ParseFloat(string(msg.Data()), 64)
		if err != nil {
			fmt.Printf("Invalid temperature value: %s\n", string(msg.Data()))
			msg.Nak() // Negative acknowledgment
		}

		// Update the sum and count
		sum += value
		count++

		// Acknowledge the message
		msg.Ack()
	})
	if err != nil {
		fmt.Printf("Error starting consumer subscription: %v\n", err)
		return
	}
	defer subscription.Stop()
	time.Sleep(time.Second * 1)
	if count == 0 {
		fmt.Println("No temperature data found.")
	} else {
		average := sum / count
		fmt.Printf("Average temperature: %.2f\n", average)
	}
}
