package main

import (
	"fmt"
	//"time"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
	"mqtt-bus/mqttclient" // 替换为您的模块路径
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
	configPath := "config.toml" // 配置文件路径（可选）
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

	// 等待一段时间以观察消息（或保持运行）
	fmt.Println("MQTT 客户端运行中，按 Ctrl+C 退出...")
	select {} // 阻塞主程序
}
