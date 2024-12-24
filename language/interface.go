package language

type LanguageInterface interface {
	Speak(eventName string) string
}
