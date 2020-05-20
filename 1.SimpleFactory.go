package main

import (
	"fmt"
)

type API interface {
	Say(name string) string
}

func NewAPI(str string) API {
	if str == "en" {
		return &English{}
	} else if str == "cn" {
		return &Chinese{}
	}
	return nil
}

type Chinese struct {}
func (*Chinese) Say(name string) string {
	return "你好 " + name
}

type English struct {}
func (*English) Say(name string) string {
	return "hello " + name
}

type Japanese struct {}
func (*Japanese) Say(name string) string {
	return "鬼子你好" + name
}

func main11(){
	api:=NewAPI("cn")
	server:=api.Say("张海涛")
	fmt.Println(server)
}
func main12(){
	api:=NewAPI("en")
	server:=api.Say("Alex")
	fmt.Println(server)
}
