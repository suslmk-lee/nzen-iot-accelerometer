package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"nzen-iot-accelerometer/common"
	"os"
	"os/signal"
	"syscall"
)

var MQTTBroker string
var MQTTClientID string
var MQTTTopic string

func init() {
	MQTTBroker = common.ConfInfo["mqtt.broker"]
	MQTTClientID = common.ConfInfo["mqtt.client.id"]
	MQTTTopic = common.ConfInfo["mqtt.topic"]
}

// AccelerometerData 구조체 정의
type AccelerometerData struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// MQTT 메시지 핸들러
var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var data AccelerometerData
	err := json.Unmarshal(msg.Payload(), &data)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return
	}
	// 여기서 데이터를 처리합니다.
	processAccelerometerData(data)
}

// 가속도 데이터 처리 함수
func processAccelerometerData(data AccelerometerData) {
	// 예: 데이터를 콘솔에 출력
	fmt.Printf("Received accelerometer data: X=%.2f, Y=%.2f, Z=%.2f\n", data.X, data.Y, data.Z)
}

func main() {
	// MQTT 클라이언트 옵션 설정
	opts := mqtt.NewClientOptions().AddBroker(MQTTBroker).SetClientID(MQTTClientID)
	opts.SetDefaultPublishHandler(messageHandler)

	// MQTT 클라이언트 생성 및 연결
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to MQTT broker: %v", token.Error())
	}

	// MQTT 토픽 구독
	if token := client.Subscribe(MQTTTopic, 0, nil); token.Wait() && token.Error() != nil {
		log.Fatalf("Error subscribing to topic: %v", token.Error())
	}

	// 프로그램 종료를 위한 신호 처리 설정
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc

	// MQTT 클라이언트 종료
	client.Disconnect(250)
	log.Println("Program terminated")
}
