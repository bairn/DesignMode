package main

import (
	"DesignMode/demo/Abstract_Factory"
	"DesignMode/demo/Bridge"
	"DesignMode/demo/Builder"
	"DesignMode/demo/Composite"
	"DesignMode/demo/Factory"
	"fmt"
)

func main11() {
	var fac Factory.OperatorFactory
	//fac = Factory.PlusOperatorFactory{}
	fac = Factory.SubOperatorFactory{}
	op := fac.Create()
	op.SetLeft(20)
	op.SetRight(10)

	fmt.Println(op.Result())
}

func main22() {
	var factory Abstract_Factory.DAOFactory
	//factory = &Abstract_Factory.MySQLFactory{}
	factory = &Abstract_Factory.OracleFactory{}

	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()

}

func main33() {
	//builder := &Builder.StringBuilder{}
	builder := &Builder.IntBuilder{}

	dict := Builder.NewDirector(builder)
	dict.Makedata()
	fmt.Println(builder.GetResult())
}

func main44() {
	root := Composite.NewComponent(Composite.CompositeNode, "root")
	r1 := Composite.NewComponent(Composite.CompositeNode, "r1")
	r2 := Composite.NewComponent(Composite.CompositeNode, "r2")
	r3 := Composite.NewComponent(Composite.CompositeNode, "r3")

	l1 := Composite.NewComponent(Composite.LeafNode, "l1")
	l2 := Composite.NewComponent(Composite.LeafNode, "l2")
	l3 := Composite.NewComponent(Composite.LeafNode, "l3")

	root.AddChild(r1)
	root.AddChild(r2)
	r1.AddChild(r3)

	r1.AddChild(l1)

	r2.AddChild(l2)
	r2.AddChild(l3)

	root.Print("")
}

func main()  {
	//m:=Bridge.NewComonMessage(Bridge.ViaSMS())
	//m:=Bridge.NewComonMessage(Bridge.ViaEmail())
	//m.SendMessage("baBy 你好","nimei")

	//m := Bridge.NewUrencyMessage(Bridge.ViaEmail())
	m := Bridge.NewUrencyMessage(Bridge.ViaSMS())
	m.SendMessage("baby 你好", "nimei")
}
