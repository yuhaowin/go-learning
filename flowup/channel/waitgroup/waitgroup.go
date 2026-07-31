package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %d\n", id, n)
		w.wg.Done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWork(id, w)
	return w
}

func main() {
	wg := sync.WaitGroup{}

	var workers [10]worker
	for i := range 10 {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i, w := range workers {
		w.in <- i * 10
	}
	for i, w := range workers {
		w.in <- i * 100
	}
	wg.Wait()
}
