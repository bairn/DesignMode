package main

import "fmt"

type Cloneable interface {
	Clone() Cloneable
}




type PrototypeManger struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManger() *PrototypeManger {
	return &PrototypeManger{prototypes:make(map[string]Cloneable)}
}

func (p *PrototypeManger) Get(name string) Cloneable {
	return p.prototypes[name]
}

func (p*PrototypeManger) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}


type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	//浅复制
	//return t

	//深复制
	tc := *t
	return &tc
}



type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t
	return &tc
}



func main() {

	mgr := NewPrototypeManger()
	t1 := &Type1{name:"type1"}

	mgr.Set("t1", t1)

	t11 := mgr.Get("t1")
	t22 := t11.Clone()
	if t11 == t22 {
		fmt.Println("浅复制")
	} else {
		fmt.Println("深复制")
	}



}