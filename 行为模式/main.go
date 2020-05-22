package main

import "DesignMode/行为模式/Chain"

func main() {
	c1:=Chain.NewProjectManagerChain()
	c2:=Chain.NewDepManagerChain()
	c3:=Chain.NewGenaralManagerChain()

	c1.SetSuccessor(c2)
	c2.SetSuccessor(c3)

	var c Chain.Manger = c1
	c.HandleFeeRequest("hupeng", 1500)
	c.HandleFeeRequest("hupeng",500)
	c.HandleFeeRequest("lixiang",4500)
	c.HandleFeeRequest("lixiang",11500)
	c.HandleFeeRequest("weishangyin",4500)
	c.HandleFeeRequest("weishangyin",11500)
}
