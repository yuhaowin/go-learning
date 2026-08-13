package main

import (
	"github.com/yuhaowin/go-learning/crawler/implement/simple/engine"
	"github.com/yuhaowin/go-learning/crawler/model"
	"github.com/yuhaowin/go-learning/crawler/parser"
	"github.com/yuhaowin/go-learning/crawler/saver"
)

func main() {
	engine.SimpleEngine{ItemChan: saver.ItemSaver()}.Run(model.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
