package main

import (
	"fmt"
	"sync"
	"time"
)

func Task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running \n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

func main() {

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)

	go Task("A", &waitGroup)
	go Task("B", &waitGroup)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task annonymous is running \n", i)
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()
}
