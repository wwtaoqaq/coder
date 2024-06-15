package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mutex := sync.Mutex{}

	mutex.Lock()
	fmt.Println("第一次加锁")
	go func() {
		mutex.Unlock()
		fmt.Println("其他协程解锁")
	}()
	time.Sleep(1000)
}
