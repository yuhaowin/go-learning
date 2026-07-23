package main

import (
	"fmt"
	"golearning/flowup/retriever/mock"
	"golearning/flowup/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{Contents: "this is a fake Retriever"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
