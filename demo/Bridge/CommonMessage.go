package Bridge

type CommonMessage struct {
	method MessageImplementer
}

func NewComonMessage(method MessageImplementer) *CommonMessage {
	return &CommonMessage{method:method}
}

func (com *CommonMessage) SendMessage(text, to string) {
	com.method.Send(text, to)
}