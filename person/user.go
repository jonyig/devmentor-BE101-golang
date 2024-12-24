package person

import "be101/language"

type User struct {
	PersonName   string
	IdentityType string
	Language     language.LanguageInterface
}

func (u User) Speak(eventName string) string {
	return u.Language.Speak(eventName)
}

func (u User) GetPersonName() string {
	return u.PersonName
}
