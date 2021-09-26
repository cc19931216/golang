package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go producer1(ch, &wg)
	wg.Add(1)
	go consumer1(ch, &wg)
	wg.Wait()
}

func producer1(ch chan<- int, wg *sync.WaitGroup) {
	defer close(ch)
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("producer: ", i)
		time.Sleep(time.Second)
	}
}

func consumer1(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for k := range ch {
		fmt.Println("consumer: ", k)
		time.Sleep(time.Second)
	}
}
