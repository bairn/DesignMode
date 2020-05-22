package Chain

import "fmt"

type RequestChain struct {
	Manger
	successor *RequestChain
}

func (rc *RequestChain) SetSuccessor(endrc *RequestChain) {
	rc.successor = endrc
}

func (rc *RequestChain) HaveRight(money int) bool {
	fmt.Println("test")
	return true
}

func (rc * RequestChain) HandleFeeRequest(name string, money int) bool {
	if rc.Manger.HaveRight(money) {
		return rc.Manger.HandleFeeRequest(name, money)
	}
	if rc.successor != nil {
		return rc.successor.HandleFeeRequest(name, money)
	}

	return false
}

