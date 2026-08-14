package engine

import (
	"github.com/yuhaowin/go-learning/crawler/internal/model"
)

// Scheduler 由 internal/scheduler 下的具体实现（SimpleScheduler、QueuedScheduler）满足，
// ConcurrentEngine 通过替换 Scheduler 实现来切换调度策略。
type Scheduler interface {
	Notifier
	Submit(model.Request)
	WorkerChan() chan model.Request
	Run()
}

type Notifier interface {
	WorkerReady(chan model.Request)
}

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan any
}

func (e *ConcurrentEngine) Run(seeds ...model.Request) {
	out := make(chan model.ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out

		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan model.Request, out chan model.ParseResult, notifier Notifier) {
	go func() {
		for {
			notifier.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
