package main

func main() {
	println("git test")
}

type Notifier interface {
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
