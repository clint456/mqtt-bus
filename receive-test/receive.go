package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"mqtt-bus/mqttclient"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
)

func main() {
	// 定义消息处理回调函数
	handler := func(topic string, event dtos.Event) {
		// fmt.Printf("收到消息 - 主题: %s, 事件ID: %s\n", topic, event.Id)
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

	// 设置信号处理
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("MQTT 客户端运行中，按 Ctrl+C 退出...")
	// 等待信号
	<-sigChan
	fmt.Println("收到终止信号，正在停止...")

	// 停止客户端
	client.Stop()
	fmt.Println("程序已退出")
}
