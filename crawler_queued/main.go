package main

import (
	"github.com/yuhaowin/go-learning/crawler_queued/engine"
	"github.com/yuhaowin/go-learning/crawler_queued/parser"
	"github.com/yuhaowin/go-learning/crawler_queued/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
