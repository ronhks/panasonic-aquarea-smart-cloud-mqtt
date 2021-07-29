package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strconv"
	"strings"
	"time"
	config "github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/config"
)

var mqttClient mqtt.Client
var token mqtt.Token

func initMqttConnection() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s://%s:%s", "tcp", config.MqttServer, config.MqttPort))
	opts.SetPassword(config.MqttPass)
	opts.SetUsername(config.MqttLogin)
	opts.SetClientID(config.MqttClientID)

	opts.SetKeepAlive(MqttKeepalive)
	opts.SetOnConnectHandler(startsub)
	opts.SetConnectionLostHandler(connLostHandler)

	// connect to broker
	mqttClient = mqtt.NewClient(opts)
	//defer client.Disconnect(uint(2))

	token = mqttClient.Connect()
	if token.Wait() && token.Error() != nil {
		fmt.Printf("Fail to connect broker, %v", token.Error())
	}
}

func connLostHandler(c mqtt.Client, err error) {
	fmt.Printf("Connection lost, reason: %v\n", err)

	//TODO Perform additional action...
}

func startsub(c mqtt.Client) {
	c.Subscribe("aquarea/+/+/set", 2, HandleMSGfromMQTT)

	//TODO Perform additional action...
}

func HandleMSGfromMQTT(c mqtt.Client, msg mqtt.Message) {
	s := strings.Split(msg.Topic(), "/")
	if len(s) > 3 {
		DeviceID := s[1]
		Operation := s[2]
		fmt.Printf("Device ID %s \n Operation %s", DeviceID, Operation)
		if Operation == "Zone1SetpointTemperature" {
			i, err := strconv.ParseFloat(string(msg.Payload()), 32)
			fmt.Printf("i=%d, type: %T\n err: %s", i, i, err)
			//TODO fix thisstr := MakeChangeHeatingTemperatureJSON(DeviceID, 1, int(i))
			//TODO fmt.Printf("\n %s \n ", str)
			//TODO set user action SetUserOption(client, DeviceID, str)

		}
	}
	fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	fmt.Printf(".")

}

func PublishLog(topic string, msg string) {

	topicWithRoot := config.MqttTopicRoot + topic
	fmt.Println("Published to topic: ", topicWithRoot, " with data: ", msg)
	msg = strings.TrimSpace(msg)
	msg = strings.ToUpper(msg)
	token = mqttClient.Publish(topicWithRoot, byte(0), false, msg)
	if token.Wait() && token.Error() != nil {
		fmt.Printf("Fail to publish, %v", token.Error())
	}

	updateLastUpdatedTimestamp()

}

func updateLastUpdatedTimestamp() {
	lastUpdateTopic := config.MqttTopicRoot + "/log/LastUpdated"
	nowEpoch := fmt.Sprintf("%d", time.Now().Unix())
	fmt.Println("Published to topic: ", lastUpdateTopic, " timestamp: ", nowEpoch)
	token = mqttClient.Publish(lastUpdateTopic, byte(0), false, nowEpoch)
	if token.Wait() && token.Error() != nil {
		fmt.Printf("Fail to publish, %v", token.Error())
	}
}