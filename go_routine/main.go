package main

import (
	"fmt"
	// "time"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(4)
	go doSomething1(wg)
	go doSomething2(wg)
	go doSomething3(wg)
	go doSomething4(wg)
	wg.Wait()
	// time.Sleep(time.Millisecond)
}

func doSomething1(wg *sync.WaitGroup) {
	// time.Sleep(time.Second * 2)
	fmt.Println("success1")
	wg.Done()
}

func doSomething2(wg *sync.WaitGroup) {
	// time.Sleep(time.Second * 2)
	fmt.Println("success2")
	wg.Done()
}

func doSomething3(wg *sync.WaitGroup) {
	// time.Sleep(time.Second * 2)
	fmt.Println("success3")
	wg.Done()
}

func doSomething4(wg *sync.WaitGroup) {
	// time.Sleep(time.Second * 2)
	fmt.Println("success4")
	wg.Done()
}
