package main

func main() {
	println("git test")
}

type MessageSender interface {
	Send(message, lang Language) Event
}

type EmailNotifier struct{}
type SMSNotifier struct{}
type TelegramNotifier struct{}

func (e *EmailNotifier) Send(message, lang Language) Event {
	return Event{}
}
func (e *SMSNotifier) Send(message, lang Language) Event {
	return Event{}
}
func (e *TelegramNotifier) Send(message, lang Language) Event {
	return Event{}
}

// 依賴注入
type NotificationService struct {
	Sender MessageSender
}

func NewNotificationService(message MessageSender) *NotificationService {
	return &NotificationService{
		Sender: message,
	}
}

func (n *NotificationService) Notify(message, lang Language) Event {
	return n.Sender.Send(message, lang)
}

type User struct {
	Lang Language
}

type Event struct {
	massage string
}

type Language struct {
	zh_TW string
	en_US string
}
