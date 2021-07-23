package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("/Users/q.s.wang/OneDrive/projects/qidi/300 Specifications - 各项规范/360 Vendors/Nader/nader.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	payload := msg.Payload()
	payload = append(payload, ","...)

	payload = bytes.Replace(payload, []byte("'"), []byte("\""), -1)

	if _, err := f.Write(payload); err != nil {
		log.Fatal(err)
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {
	var broker = "test.mosquitto.org"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client-" + uuid.NewString())
	// opts.SetUsername("emqx")
	// opts.SetPassword("public")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)
	// publish(client)

	// instead of busy loop
	for {
		time.Sleep(time.Second)
		// do something
	}
}

// func publish(client mqtt.Client) {
// 	num := 10
// 	for i := 0; i < num; i++ {
// 		text := fmt.Sprintf("Message %d", i)
// 		token := client.Publish("topic/test", 0, false, text)
// 		token.Wait()
// 		time.Sleep(time.Second)
// 	}
// }

func sub(client mqtt.Client) {
	topic := "matrinos/60e51b59f13f88fe371ac29b/1"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
