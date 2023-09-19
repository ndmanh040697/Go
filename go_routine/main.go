package main

import (
	"fmt"
	"time"
)

func main() {
	go doSomething1()
	go doSomething2()
	go doSomething3()
	doSomething4()
	time.Sleep(time.Millisecond)
}

func doSomething1() {
	time.Sleep(time.Second * 2)
	fmt.Println("success1")
}

func doSomething2() {
	time.Sleep(time.Second * 2)
	fmt.Println("success2")
}

func doSomething3() {
	time.Sleep(time.Second * 2)
	fmt.Println("success3")
}

func doSomething4() {
	time.Sleep(time.Second * 2)
	fmt.Println("success4")
}
