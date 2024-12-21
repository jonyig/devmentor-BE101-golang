package person

import "be101/language"

type PersonInterface interface {
	language.LanguageInterface
	GetPersonName() string
}
