package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mqtt-bus/mqttclient"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
	"github.com/google/uuid"
)

func main() {
	// 定义消息处理回调函数
	handler := func(topic string, event dtos.Event) {
		fmt.Printf("收到消息 - 主题: %s, 事件ID: %s\n", topic, event.Id)
		for _, reading := range event.Readings {
			fmt.Printf("  Reading: Device=%s, Resource=%s, ValueType=%s, Value=%s\n",
				reading.DeviceName, reading.ResourceName, reading.ValueType, reading.SimpleReading.Value)
		}
	}

	// 创建 MQTT 客户端
	configPath := "config.toml"
	client, err := mqttclient.NewMQTTClient(configPath, handler)
	if err != nil {
		fmt.Printf("创建 MQTT 客户端失败: %v\n", err)
		return
	}

	// 启动客户端
	if err := client.Start(); err != nil {
		fmt.Printf("启动 MQTT 客户端失败: %v\n", err)
		return
	}

	// 测试发布消息
	topic := "edgex/events/test"
	fmt.Println("发布测试消息...")
	customEvent := dtos.Event{
		Id:          uuid.New().String(),
		DeviceName:  "TestDevice",
		ProfileName: "TestProfile",
		SourceName:  "TestSource",
		Origin:      time.Now().UnixNano(),
		Readings: []dtos.BaseReading{
			{
				Id:           uuid.New().String(),
				DeviceName:   "TestDevice",
				ResourceName: "TestResource",
				ProfileName:  "TestProfile",
				Origin:       time.Now().UnixNano(),
				ValueType:    "String",
				SimpleReading: dtos.SimpleReading{
					Value: fmt.Sprintf("Test value at %s", time.Now().Format(time.RFC3339)),
				},
			},
		},
	}
	if err := client.PublishEvent(customEvent, topic); err != nil {
		fmt.Printf("发布测试消息失败: %v\n", err)
	} else {
		fmt.Printf("成功发布测试消息到主题 %s, 事件ID: %s\n", topic, customEvent.Id)
	}

	// 设置信号处理
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("MQTT 客户端运行中，按 Ctrl+C 退出...")
	<-sigChan
	fmt.Println("收到终止信号，正在停止...")

	// 停止客户端
	client.Stop()
	fmt.Println("程序已退出")
}
