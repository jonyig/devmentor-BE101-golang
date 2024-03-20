package language

type Language interface {
	GetMessage(message map[string]map[string]string) string
}

type Zn_TW struct{}
type En_US struct{}

func (z Zn_TW) GetMessage(message map[string]map[string]string) string { return "" }
func (e En_US) GetMessage(message map[string]map[string]string) string { return "" }
