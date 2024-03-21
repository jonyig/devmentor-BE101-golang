package user

import "be101_golang/models/language"

// User is the method set
type User interface {
	GetName() string
	GetPreferredLanguage() language.Language
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

func (g Guest) GetName() string                         { return g.Name }
func (g Guest) GetPreferredLanguage() language.Language { return nil }
