package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()

	smokeSignal := make(chan bool)
	evilNinja := "Tommy"
	go attack(evilNinja, smokeSignal)
	fmt.Println(<-smokeSignal)

}

func attack(target string, attacked chan bool){
	time.Sleep(time.Second)
	fmt.Println("Throwing Ninga starts at", target)
	attacked<-true 

}