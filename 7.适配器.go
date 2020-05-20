package main

import "fmt"

type Target interface {
	Request() string
}

type adapter struct {
	Adaptee
}

func (adap *adapter) Request() string {
	return adap.SpecficRequest()
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{adaptee}
}





type Adaptee interface {
	SpecficRequest() string
}

func NewAdaptee () Adaptee {
	return &adapeeImpl{}
}

type adapeeImpl struct {}
func (ada *adapeeImpl) SpecficRequest() string {
	return "SpecficRequest你妹"
}



func main() {
	adapee := NewAdaptee()
	target := NewAdapter(adapee)
	res := target.Request()
	fmt.Println(res)

}
