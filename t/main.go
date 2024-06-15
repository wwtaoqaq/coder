package main

import (
	"fmt"
	"time"
)

func startWorker(id int) {
	for i := 0; i < 100; i++ {
		go func() {
			id++
			fmt.Printf("id value: %d \n", id)
		}()
	}
}

func main() {
	startWorker(0)
	time.Sleep(3 * time.Second)

	// 主goroutine 继续执行其他操作
}
