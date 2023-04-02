package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	myString := string(msg.Payload())
	// fmt.Printf("Received message: %s ", msg.Payload())
	fmt.Println("mssage from hive => ", myString)
}

func sub(client mqtt.Client) {
	topic := "topic/test/1"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	// fmt.Printf("Subscribed to topic: %s", topic)
}

// var knt int

func main() {
	fmt.Println("Connecting...")
	// knt = 0
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	var broker = "broker.mqttdashboard.com"
	var port = 1883
	// create a new MQTT client
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetUsername("admin")
	opts.SetPassword("1234")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(1 * time.Minute)
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
	fmt.Println("Connected to MQTT broker")

	sub(client)
	<-c
	// for {
	// if !client.IsConnected() {
	// 	if token := client.Connect(); token.Wait() && token.Error() != nil {
	// 		fmt.Println(token.Error())
	// 		time.Sleep(1 * time.Second)
	// 		continue
	// 	}
	// 	fmt.Println("Connected to MQTT broker")
	// }

	// //  text := fmt.Sprintf("Hello, MQTT! %v", time.Now())
	// // token := client.Publish("my/topic", 0, false, text)
	// sub(client)
	// //  token.Wait()
	// //  fmt.Println("Published message:", text)

	// // time.Sleep(5 * time.Second)
	// }
}
