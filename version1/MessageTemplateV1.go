package version1

type MessageTemplateV1 struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	From string `json:"from"`

	Subject map[string]string `json:"subject"`
	Text    map[string]string `json:"text"`
	Html    map[string]string `json:"html"`
	Status  string            `json:"status"`
}

func EmptyMessageTemplateV1() *MessageTemplateV1 {
	return &MessageTemplateV1{}
}

func NewMessageTemplateV1(id string, name string, from string, subject string, text string) *MessageTemplateV1 {
	return &MessageTemplateV1{
		Id:      id,
		Name:    name,
		From:    from,
		Subject: map[string]string{"en": subject},
		Text:    map[string]string{"en": text},
	}
}
