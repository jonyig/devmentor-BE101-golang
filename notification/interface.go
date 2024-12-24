package notification

import "be101/person"

type NotificationInterface interface {
	Send(person.PersonInterface, string)
}
