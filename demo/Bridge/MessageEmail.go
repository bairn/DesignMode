package Bridge

import "fmt"

type MessageEmail struct {

}

func ViaEmail() MessageImplementer {
	return &MessageEmail{}
}

func (msms *MessageEmail) Send(text, to string) {
	fmt.Printf("send %s to %s via Email", text, to)
}
