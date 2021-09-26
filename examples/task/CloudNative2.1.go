package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var produce bool = true

func produceStop() {
	produce = false
}

func producer(num int, wg *sync.WaitGroup, ch chan<- string) {
	count := 1
	for produce {
		time.Sleep(time.Second)
		data := "producer: " + strconv.Itoa(num) + " make " + strconv.Itoa(count)
		ch <- data
		fmt.Println(data)
		count++
	}

	wg.Done()
}

func consumer(num int, wg *sync.WaitGroup, ch <-chan string) {
	for data := range ch {
		time.Sleep(time.Second * 2)
		fmt.Printf("consumer: %d recv: %s\n", num, data)
	}

	wg.Done()
}

func main() {
	ch := make(chan string, 5)
	wgp := &sync.WaitGroup{}
	wgv := &sync.WaitGroup{}

	for i := 1; i < 6; i++ {
		wgp.Add(1)
		go producer(i, wgp, ch)
	}

	for i := 1; i < 6; i++ {
		wgv.Add(1)
		go consumer(i, wgv, ch)
	}

	time.Sleep(time.Second * 10)
	go produceStop()

	wgp.Wait()
	close(ch)
	wgv.Wait()
}
