package notifier

import "be101_golang/models/user"

type Notifier interface {
	Notify(user user.User, message string)
}

type EmailNotifier struct{}
type SMSNotifier struct{}
type TelegramNotifier struct{}

func (e EmailNotifier) Notify(user user.User, message string)    {}
func (s SMSNotifier) Notify(user user.User, message string)      {}
func (t TelegramNotifier) Notify(user user.User, message string) {}
