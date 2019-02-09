package mylib

import (
	"fmt"
	"time"
)

func longProcess(ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func ExecContextTest() {
	ch := make(chan string)
	go longProcess(ch)

	for {
		select {
		case <-ch:
			fmt.Println("success")
			return
		}
	}
}
