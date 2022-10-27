package mqtt

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	conf "github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/config"
	"github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/data"
	"github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/device"
	"github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/heat"
	"github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt/src/water"
	log "github.com/sirupsen/logrus"
)

var mqttClient mqtt.Client
var token mqtt.Token

var maxConnectionTries = 3

func InitMqttConnection() {

	config := conf.GetConfig()

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s://%s:%s", "tcp", config.MqttServer, config.MqttPort))
	opts.SetPassword(config.MqttPass)
	opts.SetUsername(config.MqttLogin)
	opts.SetClientID(config.MqttClientID)

	opts.SetKeepAlive(config.MqttKeepalive)
	opts.SetOnConnectHandler(subscribe)
	opts.SetConnectionLostHandler(connLostHandler)

	mqttClient = mqtt.NewClient(opts)

	token = mqttClient.Connect()
	if token.Wait() && token.Error() != nil {
		log.Error("Fail to connect broker, %v", token.Error())
	}
}

func connLostHandler(mqqtClient mqtt.Client, err error) {
	log.Errorf("Connection lost, reason: %v\n", err)

	for maxConnectionTries > 0 {
		mqqtClient.Connect()
		maxConnectionTries--
	}
}

func subscribe(mqttClient mqtt.Client) {
	mqttClient.Subscribe(conf.GetConfig().MqttTopicRoot+"/water/temp/set", 0, setWaterTempHandler)
	mqttClient.Subscribe(conf.GetConfig().MqttTopicRoot+"/water/operation/on", 0, setWaterOperationOnHandler)
	mqttClient.Subscribe(conf.GetConfig().MqttTopicRoot+"/water/operation/off", 0, setWaterOperationOffHandler)
	mqttClient.Subscribe(conf.GetConfig().MqttTopicRoot+"/water/operation/set", 0, setWaterOperationHandler)
	mqttClient.Subscribe(conf.GetConfig().MqttTopicRoot+"/heat/operation/on", 0, setHeatOperationOnHandler)
	mqttClient.Subscribe(conf.GetConfig().MqttTopicRoot+"/heat/operation/off", 0, setHeatOperationOffHandler)
	mqttClient.Subscribe(conf.GetConfig().MqttTopicRoot+"/heat/operation/set", 0, setHeatOperationHandler)
	mqttClient.Subscribe(conf.GetConfig().MqttTopicRoot+"/operation/on", 0, setOperationOnHandler)
	mqttClient.Subscribe(conf.GetConfig().MqttTopicRoot+"/operation/off", 0, setOperationOffHandler)
	mqttClient.Subscribe(conf.GetConfig().MqttTopicRoot+"/operation/set", 0, setOperationHandler)
}

func setWaterTempHandler(_ mqtt.Client, msg mqtt.Message) {
	var setTemp data.SetTemp
	err := json.Unmarshal(msg.Payload(), &setTemp)

	if err != nil {
		log.Error("Fail to parse JSON, %v", token.Error())
	}

	err = water.SetWaterTemp(setTemp.NewTemp)
	if err != nil {
		log.Error(err)
		return
	}

}

func setWaterOperationOnHandler(_ mqtt.Client, _ mqtt.Message) {
	err := water.SetOperationOn()
	if err != nil {
		log.Error(err)
		return
	}
}

func setWaterOperationOffHandler(_ mqtt.Client, _ mqtt.Message) {
	err := water.SetOperationOff()
	if err != nil {
		log.Error(err)
		return
	}
}
func setHeatOperationOnHandler(_ mqtt.Client, _ mqtt.Message) {
	err := heat.SetOperationOn()
	if err != nil {
		log.Error(err)
		return
	}
}

func setHeatOperationOffHandler(_ mqtt.Client, _ mqtt.Message) {
	err := heat.SetOperationOff()
	if err != nil {
		log.Error(err)
		return
	}
}
func setOperationOnHandler(_ mqtt.Client, _ mqtt.Message) {
	err := device.SetOperationOn()
	if err != nil {
		log.Error(err)
		return
	}
}
func setOperationOffHandler(_ mqtt.Client, _ mqtt.Message) {
	err := device.SetOperationOff()
	if err != nil {
		log.Error(err)
		return
	}
}

func setOperationHandler(_ mqtt.Client, msg mqtt.Message) {

	payload := string(msg.Payload())
	var err error

	if data.OFF_STR == payload {
		err = device.SetOperationOff()
	} else if data.ON_STR == payload {
		err = device.SetOperationOn()
	}
	if err != nil {
		log.Error(err)
		return
	}
}
func setWaterOperationHandler(_ mqtt.Client, msg mqtt.Message) {

	payload := string(msg.Payload())
	var err error

	if data.OFF_STR == payload {
		err = water.SetOperationOff()
	} else if data.ON_STR == payload {
		err = water.SetOperationOn()
	}
	if err != nil {
		log.Error(err)
		return
	}
}
func setHeatOperationHandler(_ mqtt.Client, msg mqtt.Message) {

	payload := string(msg.Payload())
	var err error

	if data.OFF_STR == payload {
		err = heat.SetOperationOff()
	} else if data.ON_STR == payload {
		err = heat.SetOperationOn()
	}
	if err != nil {
		log.Error(err)
		return
	}
}

func PublishStatus(statusData data.StatusData) {
	publishLog("/outdoor/temp/now", fmt.Sprintf("%d", statusData.Status[0].OutdoorNow))
	publishLog("/heat/temp/max", fmt.Sprintf("%d", statusData.Status[0].ZoneStatus[0].HeatMax))
	publishLog("/heat/temp/min", fmt.Sprintf("%d", statusData.Status[0].ZoneStatus[0].HeatMin))
	publishLog("/heat/operation", fmt.Sprintf("%d", statusData.Status[0].ZoneStatus[0].OperationStatus))
	publishLog("/water/temp/now", fmt.Sprintf("%d", statusData.Status[0].TankStatus[0].TemparatureNow))
	publishLog("/water/temp/max", fmt.Sprintf("%d", statusData.Status[0].TankStatus[0].HeatMax))
	publishLog("/water/temp/min", fmt.Sprintf("%d", statusData.Status[0].TankStatus[0].HeatMin))
	publishLog("/water/operation", fmt.Sprintf("%d", statusData.Status[0].TankStatus[0].OperationStatus))
	publishLog("/operation", fmt.Sprintf("%d", statusData.Status[0].OperationStatus))

}

func publishLog(topic string, msg string) {

	topicWithRoot := conf.GetConfig().MqttTopicRoot + topic
	log.Trace("Published to topic: ", topicWithRoot, " with data: ", msg)
	msg = strings.TrimSpace(msg)
	msg = strings.ToUpper(msg)
	token = mqttClient.Publish(topicWithRoot, byte(0), false, msg)
	if token.Wait() && token.Error() != nil {
		log.Errorf("Fail to publish, %v", token.Error())
	}

	updateLastUpdatedTimestamp()

}

func updateLastUpdatedTimestamp() {
	lastUpdateTopic := conf.GetConfig().MqttTopicRoot + "/log/LastUpdated"
	nowEpoch := fmt.Sprintf("%d", time.Now().Unix())
	log.Trace("Published to topic: ", lastUpdateTopic, " timestamp: ", nowEpoch)
	token = mqttClient.Publish(lastUpdateTopic, byte(0), false, nowEpoch)
	if token.Wait() && token.Error() != nil {
		log.Errorf("Fail to publish, %v", token.Error())
	}
}
