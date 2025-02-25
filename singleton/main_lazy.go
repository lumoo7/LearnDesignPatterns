package main

import (
	"fmt"
	"sync"
)

type Singleton struct{}

func (st *Singleton) Say() {
	fmt.Println("Say Something.")
}

var once sync.Once
var singleton *Singleton

func GetInstance() *Singleton {
	if singleton == nil {
		once.Do(func() {
			singleton = new(Singleton)
		})
	}
	return singleton
}

func main() {
	GetInstance().Say()
}
