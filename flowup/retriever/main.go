package main

import (
	"fmt"
	"golearning/flowup/retriever/mock"
	"golearning/flowup/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string) string
}

type Connector interface {
	Poster
	Retriever
}

func download(r Retriever) string {
	return r.Get("https://www.baidu.com")
}

func upload(p Poster) string {
	return p.Post("https://www.baidu.com")
}

func session(c Connector) {
	c.Post("")
	c.Get("")
}

func main() {
	var r Retriever

	r = mock.Retriever{Contents: "this is a fake Retriever"}
	inspect(r)
	fmt.Println(&r)

	r = real.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   10,
	}
	inspect(r)

	fmt.Println(&r)

	// Type assertion
	if retriever, ok := r.(mock.Retriever); ok {
		fmt.Println(retriever.Contents)
	} else {
		fmt.Println("is not a mock retriever")
	}

	fmt.Println(download(r))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)

	// Type switch
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
