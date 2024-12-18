package person

import "be101/language"

type User struct {
	PersonName   string
	IdentityType string
	Language     language.LanguageInterface
}
