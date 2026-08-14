package main

import (
	"github.com/yuhaowin/go-learning/crawler/internal/engine"
	"github.com/yuhaowin/go-learning/crawler/internal/parser"
	"github.com/yuhaowin/go-learning/crawler/internal/saver"
	"github.com/yuhaowin/go-learning/crawler/internal/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    saver.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
