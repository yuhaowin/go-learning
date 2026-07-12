package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// *http.Request 是指向这个结构体的指针。之所以用指针而不是值传递，是因为 Request 结构体较大，用指针可以避免每次调用 handler 都拷贝整个结构体；同时 Go 标准库内部（比如读取 body）也需要能修改/共享同一个 Request 实例。
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
