package main

import "github.com/sawickicezary/golang-design-patterns/singleton"

func main() {
	for i := 0; i < 24; i++ {
		singleton.GetInstance()
	}
}
