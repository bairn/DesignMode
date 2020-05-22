package main

import (
	"DesignMode/行为模式/Chain"
	"DesignMode/行为模式/Observer"
	"DesignMode/行为模式/Strategy"
)


func main() {
	//ctx := Strategy.NewMMContext("marry", 18, &Strategy.Girl{})
	ctx := Strategy.NewMMContext("alis", 28, &Strategy.Women{})
	ctx.Pao()
}

func main2() {
	subject := Observer.NewSubject()
	r1 := Observer.NewReader("lixiang")
	r2 := Observer.NewReader("hupeng")
	r3 := Observer.NewReader("zhangbo")

	subject.Attch(r1)
	subject.Attch(r2)
	subject.Attch(r3)

	subject.UpdateContext("妹子来了")
	r4 := Observer.NewReader("yangjie")
	subject.Attch(r4)
	subject.UpdateContext("漂亮妹子来了")

}

func main1() {
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
