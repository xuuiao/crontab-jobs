package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var x int64 = 0x3333333333333333

func main() {
	//fmt.Println("start")
	//var mutex  = &sync.RWMutex{}
	//go rwMutexFunc1(mutex)
	//go rwMutexFunc2(mutex)
	//fmt.Println("main wait...")
	//<- time.After(time.Second * 15)
	//fmt.Println("main over")
	go storeFunc()
	for {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("%x\n", atomic.LoadInt64(&x))
	}
}

func storeFunc() {
	for i := 0; ; i++ {
		if i%2 == 0 {
			x = 0x1111111111111111
			//atomic.StoreInt64(&x, 0x1111111111111111)
		} else {
			//atomic.StoreInt64(&x, 0x2222222222222222)
			x = 0x2222222222222222
		}
	}
}

func rwMutexFunc1(m *sync.RWMutex) {
	m.RLock()
	fmt.Println("func1 lock")
	<-time.After(time.Second * 5)
	fmt.Println("func1 time out")
	m.RUnlock()
	fmt.Println("func1 unlock")
}

func rwMutexFunc2(m *sync.RWMutex) {
	m.Lock()
	fmt.Println("func2 lock")
	<-time.After(time.Second * 5)
	fmt.Println("func2 time out")
	m.Unlock()
	fmt.Println("func2 unlock")
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
