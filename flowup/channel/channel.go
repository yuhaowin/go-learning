package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		fmt.Printf("Worker %d received %d\n", id, <-c) // 这里「收」别人发送的数据
	}
}

func workerHandClose1(id int, c chan int) {
	for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func workerHandClose2(id int, c chan int) {
	for n := range c { // channel range
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func chanDemo1() {
	var channels [10]chan int
	for i := range 10 {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}
	for i := range 10 {
		c := channels[i]
		c <- i * 10
	}
	for i := range 10 {
		c := channels[i]
		c <- i * 100
	}
	time.Sleep(time.Millisecond)
}

// ------------------------------------------------------------------------------------------------

func createWorker(id int) chan<- int { // 也可以指明 channel 的方向，chan <- int, 这里 channel 是用来给别人「发」数据的
	c := make(chan int)
	go worker(id, c)
	return c
}

//func createWorker1(id int) (c chan<- int ){
//	c = make(chan int)
//	go func() {
//		for {
//			fmt.Printf("Worker %d received %d\n", id, <-c)
//		}
//	}()
//	return c
//}

func chanDemo2() {
	var channels [10]chan<- int
	for i := range 10 {
		channels[i] = createWorker(i)
	}
	for i := range 10 {
		c := channels[i]
		c <- i * 10
	}
	for i := range 10 {
		c := channels[i]
		c <- i * 100
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Millisecond)
}

func channelClose() { // close 永远是发送方发起的。
	c := make(chan int, 3)
	go workerHandClose2(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	close(c)
	time.Sleep(time.Millisecond)
}

// 理论基础 : Communication Sequential Process (CSP)
// Don't communicate by sharing memory; share memory by communicating.
// 不要通过共享内存来通信；通过通信来共享内存
func main() {
	// Channel as first-class citizen
	//chanDemo1()
	//chanDemo2()

	// Buffered Channel
	//bufferedChannel()

	// Channel Close and Range
	channelClose()
}
