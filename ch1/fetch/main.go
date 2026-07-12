package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		fmt.Printf("received url: %s\n", url)
		resp, err := http.Get(url)
		if err != nil {
			// Fprintf ]把格式化后的字符串写到指定的 io.Writer，这里是 os.Stderr（标准错误流），而不是像 fmt.Printf 那样写到标准输出 os.Stdout。
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", body)
	}
}
