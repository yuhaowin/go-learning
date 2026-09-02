# 学习笔记

## net/http 请求处理链路

从 `http.ListenAndServe("localhost:8080", nil)` 到 `handler(w, r)` 被调用，完整链路（源码：`net/http/server.go`，Go 版本见本机 `/opt/homebrew/opt/go/libexec/src/net/http/server.go`）：

```
ListenAndServe (L3673)
  → Server.Serve (accept 连接，每个连接开一个 goroutine)
    → conn.serve (读取解析请求，构造 w/r)
      → serverHandler.ServeHTTP (L3302)
          handler := srv.Handler
          if handler == nil { handler = DefaultServeMux }   // 传 nil 就回退到默认 mux
        → ServeMux.ServeHTTP (L2814)
            findHandler(r) 按路径做最长前缀匹配
          → HandlerFunc.ServeHTTP (L2285)
              f(w, r)   // 真正调用你写的函数
```

参考代码位置：`ch1/server1/main.go:10-11`（`http.HandleFunc("/", handler)` + `http.ListenAndServe(..., nil)`）

## HandlerFunc：函数类型实现接口的适配器模式

源码位置：`net/http/server.go:2282-2287`

```go
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

**理解要点：**

1. `type HandlerFunc func(ResponseWriter, *Request)` 定义了一个新类型，其底层类型是函数 `func(ResponseWriter, *Request)`。Go 允许给具名类型（包括函数类型）定义方法。

2. `net/http` 的 `Handler` 接口要求实现 `ServeHTTP(ResponseWriter, *Request)`：
   ```go
   type Handler interface {
       ServeHTTP(ResponseWriter, *Request)
   }
   ```
   只要给 `HandlerFunc` 定义了这个方法，`HandlerFunc` 类型就自动满足 `Handler` 接口。

3. 方法体 `f(w, r)`：因为接收者 `f` 本身底层就是函数，直接调用它、把 `w, r` 传进去执行。

**串联 `http.HandleFunc("/", handler)` 的转换过程：**

```go
// handler 只是个普通函数，本身不满足 Handler 接口（没有方法）
func handler(w http.ResponseWriter, r *http.Request) { ... }

// HandleFunc 内部：
mux.register(pattern, HandlerFunc(handler))
//                     ^^^^^^^^^^^^^^^^^^^
//     类型转换：把普通函数“伪装”成 HandlerFunc 类型
//     转换后就“获得”了 ServeHTTP 方法，从而满足 Handler 接口
```

之后 `ServeMux` 统一用 `Handler` 接口调度：
```go
h.ServeHTTP(w, r)
// h 的动态类型是 HandlerFunc，实际执行 f(w, r)
// f 就是最初写的 handler，等价于直接调用 handler(w, r)
```

**为什么要绕这一圈：**

`ServeMux` 内部要用统一的 `Handler` 接口存储和调度所有路由（既支持传普通函数，也支持传实现了 `ServeHTTP` 方法的结构体）。但 Go 类型系统里函数值本身没有方法，不能直接满足接口。`HandlerFunc` 这种"类型转换 + 方法转发"的写法，是 Go 里"用最小适配层把函数值适配成接口"的标准手法，也是 `http.HandlerFunc` 被反复引用为经典范例的原因。

## HandleFunc 使用方式

```go
func main() {
    http.HandleFunc("/", handler)        // 子树匹配：匹配所有未被更具体规则匹配的路径
    http.HandleFunc("/about", about)     // 精确匹配 /about
    http.HandleFunc("/api/", apiHandler) // 子树匹配 /api/ 开头的所有路径

    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
```

匹配规则：
- pattern 以 `/` 结尾 → 子树匹配（前缀匹配）。
- pattern 不以 `/` 结尾 → 精确匹配。
- 多个 pattern 都能匹配时，选最长（最具体）的那个。

Go 1.22+ 支持带方法和路径参数的增强写法：
```go
http.HandleFunc("GET /items/{id}", func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    fmt.Fprintf(w, "item id = %s\n", id)
})
```

## 函数字面量赋值给变量 vs 具名函数声明

```go
handler := func(w http.ResponseWriter, r *http.Request) {
    lissajous(w)
}
```

等价于具名函数声明：

```go
func handler(w http.ResponseWriter, r *http.Request) {
    lissajous(w)
}
```

区别：
- `handler := func(...) {...}`：函数字面量（匿名函数），在函数体内部动态创建一个局部变量，类型是 `func(http.ResponseWriter, *http.Request)`，只在当前作用域可见。
- `func handler(...) {...}`：包级函数声明，全局（包内）可见。

常见使用场景：需要闭包捕获外部变量（比如根据参数生成不同行为的 handler），或者只是临时用一下、不想污染包级命名空间。
