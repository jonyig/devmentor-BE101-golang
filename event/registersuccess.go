package event

import (
	"be101/notification"
	"be101/person"
	"fmt"
)

type RegisterSuccess struct {
	notifies []notification.NotificationInterface
}

func (e *RegisterSuccess) SetNotify(n notification.NotificationInterface) {
	e.notifies = append(e.notifies, n)
}
func (e *RegisterSuccess) RoleTransition(p *person.User) *person.Student {
	newStudent := &person.Student{
		StudentId: "1000",
	}
	return newStudent
}

func (e *RegisterSuccess) Trigger(p *person.User) {
	e.SetNotify(notification.Email{})
	e.SetNotify(notification.Telegram{})
	e.SetNotify(notification.Sms{})
	newStudent := e.RoleTransition(p)
	for _, notify := range e.notifies {
		notify.Send()
		fmt.Printf("Register Success! ")
	}
	fmt.Println("Student ID is " + newStudent.StudentId)
}

// feat: event module initial
// chore:
// refactor:
// test:
// fix:
// remote: github, gitlab
// parent: parent project

// fork kubernetes

// remote: alexdev/kubernetes
// parent: google/kubernetes

// gitlab -> remote
// git push --set-upstream origin feature-01
// azure -> empty
// git push --set-upstream azure feature-01
