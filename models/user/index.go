package user

import language "be101_golang/models/Language"

// User is the method set
type User interface {
	GetName() string
	GetPreferredLanguage() string
}

type Student struct {
	Name              string
	PreferredLanguage language.Language
}
type Guest struct {
	Name string
}

func (s Student) GetName() string                         { return s.Name }
func (s Student) GetPreferredLanguage() language.Language { return s.PreferredLanguage }
