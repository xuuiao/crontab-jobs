package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("start")
	var mutex = &sync.Mutex{}
	go mutexFunc1(mutex)
	go mutexFunc2(mutex)
	fmt.Println("main wait...")
	<-time.After(time.Second * 15)
	fmt.Println("main over")
}

func mutexFunc1(m *sync.Mutex) {
	m.Lock()
	fmt.Println("func1 lock")
	<-time.After(time.Second * 5)
	fmt.Println("func1 time out")
	m.Unlock()
	fmt.Println("func1 unlock")
}

func mutexFunc2(m *sync.Mutex) {
	m.Lock()
	fmt.Println("func2 lock")
	<-time.After(time.Second * 5)
	fmt.Println("func2 time out")
	m.Unlock()
	fmt.Println("func2 unlock")
}
