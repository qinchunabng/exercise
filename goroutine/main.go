package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func putNum(intChan chan int) {
	for i := 2; i < 1000; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for num := range intChan {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	wg.Done()
	exitChan <- true
}

func printPrime(primeChan chan int) {
	for v := range primeChan {
		fmt.Println(v)
	}
	wg.Done()
}

func main() {
	start := time.Now().UnixMilli()
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 16)

	//存放数字的协程
	wg.Add(1)
	go putNum(intChan)

	//统计素数的协程
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}

	//打印素数的协程
	wg.Add(1)
	go printPrime(primeChan)

	wg.Add(1)
	go func() {
		for i := 0; i < 20; i++ {
			<-exitChan
		}
		wg.Done()
		close(primeChan)
	}()

	wg.Wait()
	cost := time.Now().UnixMilli() - start
	fmt.Printf("统计完成，耗时%d毫秒\n", cost)
}
