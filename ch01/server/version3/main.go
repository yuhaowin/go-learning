package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	lissajousHandler := func(w http.ResponseWriter, r *http.Request) {
		//lissajous(w)
	}

	// 用 var 声明，写法如下：
	var lissajousHandler1 func(http.ResponseWriter, *http.Request) = func(w http.ResponseWriter, r *http.Request) {
		//lissajous(w)
	}

	// 或者让编译器自动推断类型（更常见）：
	var lissajousHandler2 = func(w http.ResponseWriter, r *http.Request) {
		//lissajous(w)
	}

	http.HandleFunc("/lissajous", lissajousHandler)
	http.HandleFunc("/lissajous1", lissajousHandler1)
	http.HandleFunc("/lissajous2", lissajousHandler2)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
