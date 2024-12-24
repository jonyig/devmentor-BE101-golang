package person

import "be101/language"

type Student struct {
	PersonName   string
	IdentityType string
	Language     language.LanguageInterface
}

func (s Student) Speak(eventName string) string {
	return s.Language.Speak(eventName)
}

func (s Student) GetPersonName() string {
	return s.PersonName
}
