package event

import (
	"be101_golang/models/notifier"
	"be101_golang/models/user"
)

type Event interface {
	AddNotifier(notifier notifier.Notifier)
	Trigger(user user.User)
}

type Signup struct {
	studentID string
}

type BookClass struct {
	classID string
}

type CancelClass struct {
	classID string
}

func (s Signup) AddNotifier(notifier notifier.Notifier) {}
func (s Signup) Trigger(user user.User)                 {}

func (b BookClass) AddNotifier(notifier notifier.Notifier) {}
func (b BookClass) Trigger(user user.User)                 {}

func (c CancelClass) AddNotifier(notifier notifier.Notifier) {}
func (c CancelClass) Trigger(user user.User)                 {}
