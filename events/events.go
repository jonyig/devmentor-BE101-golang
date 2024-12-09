package events

import (
	"be101/notifications"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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

func NewEventRegisterSuccess(person interface{}) EventRegisterSuccess {
	newNotificationByMail := notifications.EmailNotificationChannel{}
	newNotificationByMail.Send()
	newNotificationBySMS := notifications.SmsNotificationChannel{}
	newNotificationBySMS.Send()
	return EventRegisterSuccess{
		EventID: GenerateUniqueID(),
		Person:  person,
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
	newNotificationByMail := notifications.EmailNotificationChannel{}
	newNotificationByMail.Send()
	newNotificationByTelegram := notifications.TelegramNotificationChannel{}
	newNotificationByTelegram.Send()
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

func NewEventNewyearCelebrate(person interface{}) EventNewyearCelebrate {
	newNotificationByLine := notifications.LineNotificationChannel{}
	newNotificationByLine.Send(person.language)
	return EventNewyearCelebrate{
		EventID: GenerateUniqueID(),
		Person:  person,
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
	newNotificationByMail := notifications.EmailNotificationChannel{}
	newNotificationByMail.Send()
	newNotificationByTelegram := notifications.TelegramNotificationChannel{}
	newNotificationByTelegram.Send()
	return EventCancelClasses{
		EventID: GenerateUniqueID(),
		Person:  person,
	}
}
