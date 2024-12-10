package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	user1 := Person{
		PersonName:   "Alex",
		IdentityType: "User",
		Language:     "en-US",
		Events:       []Event{},
	}

	user1.Events = append(user1.Events, NewEventRegisterSuccess(&user1)) // 註冊成功: email & sms
	user1.Events = append(user1.Events, NewEventScheduleSuccess(user1))  // 學生 預約課程: email & telegram
	user1.Events = append(user1.Events, NewEventCancelClasses(user1))    // 學生 取消課程: email & telegram
	user1.Events = append(user1.Events, NewEventNewYearCelebrate(user1)) // 新年通知
}

/* Event */

func GenerateUniqueID() string {
	timestamp := time.Now().UnixNano()
	randomPart := rand.Int63n(1e6)
	return fmt.Sprintf("%d-%06d", timestamp, randomPart)
}

type Event interface {
	GetEventID() string
}

type EventRegisterSuccess struct {
	EventID string
	Person  interface{}
}

func (e EventRegisterSuccess) GetEventID() string {
	return e.EventID
}

func NewEventRegisterSuccess(p *Person) EventRegisterSuccess {
	p.IdentityType = "Student"
	newNotificationByMail := EmailNotificationChannel{}
	newNotificationByMail.Send("Register Success")
	newNotificationBySMS := SmsNotificationChannel{}
	newNotificationBySMS.Send("Register Success")
	return EventRegisterSuccess{
		EventID: GenerateUniqueID(),
	}
}

type EventScheduleSuccess struct {
	EventID string
	Person  interface{}
}

func (e EventScheduleSuccess) GetEventID() string {
	return e.EventID
}

func NewEventScheduleSuccess(person interface{}) EventScheduleSuccess {
	newNotificationByMail := EmailNotificationChannel{}
	newNotificationByMail.Send("Schedule Success")
	newNotificationByTelegram := TelegramNotificationChannel{}
	newNotificationByTelegram.Send("Schedule Success")
	return EventScheduleSuccess{
		EventID: GenerateUniqueID(),
		Person:  person,
	}
}

type EventNewyearCelebrate struct {
	EventID string
	Person  interface{}
}

func (e EventNewyearCelebrate) GetEventID() string {
	return e.EventID
}

func NewEventNewYearCelebrate(p Person) EventNewyearCelebrate {
	newNotificationByLine := LineNotificationChannel{}
	newNotificationByLine.Send("New Year Celebrate", p)
	return EventNewyearCelebrate{
		EventID: GenerateUniqueID(),
	}
}

type EventCancelClasses struct {
	EventID string
	Person  interface{}
}

func (e EventCancelClasses) GetEventID() string {
	return e.EventID
}

func NewEventCancelClasses(person interface{}) EventCancelClasses {
	newNotificationByMail := EmailNotificationChannel{}
	newNotificationByMail.Send("Cancel the Classes")
	newNotificationByTelegram := TelegramNotificationChannel{}
	newNotificationByTelegram.Send("Cancel the Classes")
	return EventCancelClasses{
		EventID: GenerateUniqueID(),
		Person:  person,
	}
}

/* Person */

type Person struct {
	PersonName   string
	IdentityType string
	Language     string
	Events       []Event
}

type User struct {
	Person
}

type Student struct {
	Person
}

func (s Student) Echo() {
	fmt.Println(s.PersonName + " now is a student")
}

/* Notification */

type NotificationChannel struct {
	PersonName string
	Text       string
}

func (nc *NotificationChannel) Send(eventTitle string) {
	fmt.Println(eventTitle + " - Sending notification")
}

type EmailNotificationChannel struct {
	NotificationChannel
}

func (emailnc *EmailNotificationChannel) Send(eventTitle string) {
	fmt.Println(eventTitle + " - Sending email notification")
}

type SmsNotificationChannel struct {
	NotificationChannel
}

func (smsnc *SmsNotificationChannel) Send(eventTitle string) {
	fmt.Println(eventTitle + " - Sending SMS notification")
}

type TelegramNotificationChannel struct {
	NotificationChannel
}

func (telegramnc *TelegramNotificationChannel) Send(eventTitle string) {
	fmt.Println(eventTitle + " - Sending Telegram notification")
}

type LineNotificationChannel struct {
	NotificationChannel
}

func (linenc *LineNotificationChannel) Send(eventTitle string, p Person) {
	fmt.Println(eventTitle + " - Sending Line notification using " + p.Language + " language")
}
