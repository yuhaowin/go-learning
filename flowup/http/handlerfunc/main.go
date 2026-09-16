package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

// ---------- 1. 标准库的 HandlerFunc ----------

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello from plain func")
}

// 中间件：接收 Handler，返回 Handler，返回值靠 HandlerFunc 把闭包适配成接口
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("->", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// ---------- 2. 自己写一个函数类型适配器 ----------

// Greeter 1. Define the Single-Method Interface
type Greeter interface {
	Greet(name string) string
}

// GreeterFunc 2. Define the Function Type Adapter
type GreeterFunc func(string) string

// Greet 3. Make the function type implement the interface
func (f GreeterFunc) Greet(name string) string {
	return f(name) // Call the underlying function itself
}

// 调用方只依赖 Greeter 接口，不关心背后是 struct 还是函数
func welcome(g Greeter, name string) {
	fmt.Println(g.Greet(name))
}

func english(name string) string {
	return "Hello, " + name
}

func main() {
	// http.HandlerFunc(hello) 是类型转换，不是调用
	var h http.Handler = logging(http.HandlerFunc(hello))

	// 用 httptest 直接触发，不用真的监听端口
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	h.ServeHTTP(rec, req)
	fmt.Print("response: ", rec.Body.String())

	// GreeterFunc 本身是类型名，不是函数名。Go 里 T(x) 形式、T 是类型时，一律是类型转换，和 int64(x)、[]byte(s) 同一语法。
	// 普通函数：签名匹配，直接类型转换
	welcome(GreeterFunc(english), "Gopher")

	// 闭包：可以捕获状态
	count := 0
	welcome(GreeterFunc(func(name string) string {
		count++
		return fmt.Sprintf("你好，%s（第 %d 次）", name, count)
	}), "Gopher")
}
