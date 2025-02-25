package main

import "fmt"

type Phone interface {
	SayPhoneName()
}

type HuaWei struct{}

type Apple struct{}

func NewPhone(phone string) Phone {
	switch phone {
	case "apple":
		return new(Apple)
	case "huawei":
		return new(HuaWei)
	default:
		return nil
	}
}

func (h *HuaWei) SayPhoneName() {
	fmt.Println("HuaWei")
}

func (a *Apple) SayPhoneName() {
	fmt.Println("Apple")
}

func main() {
	NewPhone("apple").SayPhoneName()
	NewPhone("huawei").SayPhoneName()
}
