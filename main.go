package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	user1 := User{
		Person: Person{
			PersonName:   "Alex",
			IdentityType: "User",
			Language:     "en-US",
			Events:       []Event{},
		},
	}
	fmt.Println("My name is " + user1.PersonName + ", I am a " + user1.IdentityType)
	user1.Events = append(user1.Events, NewEventRegisterSuccess(&user1)) // 註冊成功: email & sms
	fmt.Println("My name is " + user1.PersonName + ", I am a " + user1.IdentityType)
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

func NewEventRegisterSuccess(u *User) EventRegisterSuccess {
	u.IdentityType = "Student"
	newStudent := Student{
		Person:    u.Person,
		StudentId: fmt.Sprintf("S%06d", rand.Intn(1000000)),
	}
	u.Person = newStudent.Person
	newNotificationByMail := EmailNotificationChannel{}
	newNotificationByMail.Send("Register Success")
	newNotificationBySMS := SmsNotificationChannel{}
	newNotificationBySMS.Send("Register Success")
	return EventRegisterSuccess{
		EventID: GenerateUniqueID(),
		Person:  newStudent,
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

func NewEventNewYearCelebrate(u User) EventNewyearCelebrate {
	newNotificationByLine := LineNotificationChannel{}
	newNotificationByLine.Send("New Year Celebrate", u)
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

func (u User) Echo() {
	fmt.Println(u.PersonName + " now is a User")
}

type Student struct {
	Person
	StudentId string
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

func (linenc *LineNotificationChannel) Send(eventTitle string, u User) {
	fmt.Println(eventTitle + " - Sending Line notification using " + u.Language + " language")
}
