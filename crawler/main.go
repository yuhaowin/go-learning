package main

import (
	"github.com/yuhaowin/go-learning/crawler/engine"
	"github.com/yuhaowin/go-learning/crawler/parser"
)

func main() {
	engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
