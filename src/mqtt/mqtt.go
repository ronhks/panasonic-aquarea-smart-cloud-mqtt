package mqtt

import (
	conf "config"
	"data"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"time"
)

var mqttClient mqtt.Client
var token mqtt.Token

func InitMqttConnection() {

	config := conf.GetConfig()

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s://%s:%s", "tcp", config.MqttServer, config.MqttPort))
	opts.SetPassword(config.MqttPass)
	opts.SetUsername(config.MqttLogin)
	opts.SetClientID(config.MqttClientID)

	opts.SetKeepAlive(time.Duration(config.MqttKeepalive))
	opts.SetOnConnectHandler(subscribe)
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

func subscribe(c mqtt.Client) {
	c.Subscribe(conf.GetConfig().MqttTopicRoot + "/water/temp/set", 0, handleMSGfromMQTT)
}

func handleMSGfromMQTT(c mqtt.Client, msg mqtt.Message) {
	topic := strings.Split(msg.Topic(), "/")
	for _, path := range topic {
		fmt.Printf("%s\n", path)
	}



}

func PublishStatus (statusData data.StatusData) {
	publishLog("/outdoor/temp/now", fmt.Sprintf("%d", statusData.Status[0].OutdoorNow))
	publishLog("/heat/temp/max", fmt.Sprintf("%d", statusData.Status[0].ZoneStatus[0].HeatMax))
	publishLog("/heat/temp/min", fmt.Sprintf("%d", statusData.Status[0].ZoneStatus[0].HeatMin))
	publishLog("/heat/operation", fmt.Sprintf("%d", statusData.Status[0].ZoneStatus[0].OperationStatus))
	publishLog("/water/temp/now", fmt.Sprintf("%d", statusData.Status[0].TankStatus[0].TemparatureNow))
	publishLog("/water/temp/max", fmt.Sprintf("%d", statusData.Status[0].TankStatus[0].HeatMax))
	publishLog("/water/temp/min", fmt.Sprintf("%d", statusData.Status[0].TankStatus[0].HeatMin))
	publishLog("/water/operation", fmt.Sprintf("%d", statusData.Status[0].TankStatus[0].OperationStatus))

}

func publishLog(topic string, msg string) {

	topicWithRoot := conf.GetConfig().MqttTopicRoot + topic
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
	lastUpdateTopic := conf.GetConfig().MqttTopicRoot + "/log/LastUpdated"
	nowEpoch := fmt.Sprintf("%d", time.Now().Unix())
	fmt.Println("Published to topic: ", lastUpdateTopic, " timestamp: ", nowEpoch)
	token = mqttClient.Publish(lastUpdateTopic, byte(0), false, nowEpoch)
	if token.Wait() && token.Error() != nil {
		fmt.Printf("Fail to publish, %v", token.Error())
	}
}