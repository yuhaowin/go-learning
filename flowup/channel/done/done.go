package main

import (
	"fmt"
)

type worker struct {
	in   chan int
	done chan bool
}

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %d\n", id, n)
		go func() {
			w.done <- true
		}()
	}
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w)
	return w
}

func main() {
	var workers [10]worker
	for i := range 10 {
		workers[i] = createWorker(i)
	}
	for i, w := range workers {
		w.in <- i * 10
	}
	for i, w := range workers {
		w.in <- i * 100
	}
	// wait for all of them
	for _, w := range workers {
		<-w.done
		<-w.done
	}
}
