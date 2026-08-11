package main

import (
	"github.com/yuhaowin/go-learning/crawler_concurrent/engine"
	"github.com/yuhaowin/go-learning/crawler_concurrent/parser"
)

func main() {
	engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
