package main

import (
	"fmt"
	"sync"
)

type Single struct {

}

var singleton *Single
var once sync.Once

func GetInterface()*Single {
	once.Do(func() {
		singleton = &Single{}
	})
	return singleton
}

func main() {
	i1 := GetInterface()
	i2 := GetInterface()
	if i1 == i2 {
		fmt.Println("相等")
	} else {
		fmt.Println("不等")
	}
}
