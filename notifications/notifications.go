package notifications

import "fmt"

type NotificationChannel struct {
	PersonName string
	Text       string
}

func (nc *NotificationChannel) Send() {
	fmt.Println("Sending notification")
}

type EmailNotificationChannel struct {
	NotificationChannel
}

func (emailnc *EmailNotificationChannel) Send() {
	fmt.Println("Sending email notification")
}

type SmsNotificationChannel struct {
	NotificationChannel
}

func (smsnc *SmsNotificationChannel) Send() {
	fmt.Println("Sending SMS notification")
}

type TelegramNotificationChannel struct {
	NotificationChannel
}

func (telegramnc *TelegramNotificationChannel) Send() {
	fmt.Println("Sending Telegram notification")
}

type LineNotificationChannel struct {
	NotificationChannel
}

func (linenc *LineNotificationChannel) Send(language string) {
	fmt.Println("Sending Line notification using " + language + " language")
}
