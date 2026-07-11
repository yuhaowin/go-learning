package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// Sprint：不打印到任何地方，而是返回一个字符串（S 开头就是 "String"，返回值是 string）
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	written, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	duration := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", duration, written, url)
}
