package main

import "fmt"

type Singleton struct{}

func (st *Singleton) Say() {
	fmt.Println("Say Something.")
}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func GetInstance() *Singleton {
	return singleton
}

func main() {
	GetInstance().Say()
}
