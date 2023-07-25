package main

import (
	"fmt"
)

// 定义一个接口用于定义具体实现需要遵循的方法
type MessageSender interface {
	SendMessage(message string) error
}

// 实现具体的消息发送器，这里以邮件发送为例
type EmailSender struct{}

func (e *EmailSender) SendMessage(message string) error {
	fmt.Println("Sending email:", message)
	// 在此处编写邮件发送逻辑...
	return nil
}

// 业务逻辑组件，它将依赖于 MessageSender 接口而不是具体的实现类。
type NotificationService struct {
	sender MessageSender // 使用接口类型作为成员变量类型
}

func (n *NotificationService) SendNotification(message string) error {
	err := n.sender.SendMessage(message)
	if err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}
	fmt.Println("Notification sent successfully")
	return nil
}

func main() {
	emailSender := &EmailSender{} // 创建具体实现类对象

	notificationService := &NotificationService{
		sender: emailSender, // 将具体实现传入业务逻辑组件中
	}

	err := notificationService.SendNotification("Hello, World!")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
