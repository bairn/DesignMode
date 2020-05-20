package main

import "fmt"

type Operator interface {
	SetLeft(int)
	SetRight(int)
	Result()int
}

type OperatorFactory interface {
	Create() Operator
}


type OperatorBase struct {
	left, right int
}

func (op *OperatorBase) SetLeft(left int) {
	op.left = left
}

func (op *OperatorBase) SetRight(right int) {
	op.right = right
}


type PlusOperator struct {
	*OperatorBase
}

func (o PlusOperator) Result() int {
	return o.left + o.right
}


type PlusOperatorFactory struct {}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{&OperatorBase{}}
}



//操作的抽象
type SubOperatorFactory struct {}
//操作类中包含操作数
type SubOperator struct {
	*OperatorBase
}
//实际运行
func (o SubOperator)Result()int{
	return o.left-o.right
}
func (SubOperatorFactory)Create()Operator{
	return  &SubOperator{&OperatorBase{}}
}


func main() {
	var fac OperatorFactory
	//fac = PlusOperatorFactory{}
	fac = SubOperatorFactory{}
	op := fac.Create()
	op.SetLeft(10)
	op.SetRight(5)

	fmt.Println(op.Result())
}