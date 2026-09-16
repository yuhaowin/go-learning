# 函数类型适配器（Function Type Adapter / HandlerFunc 模式）

来源：`net/http/server.go`

```go
type Handler interface {
ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func (ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
f(w, r)
}
```

## 这是什么

这种写法在 Go 里叫 **函数类型适配器**，因为它出自 `net/http`，也常被直接称为 **HandlerFunc 模式**。
本质是 **适配器模式（Adapter Pattern）** 在 Go 中的惯用实现：把一个"普通函数"适配成一个"接口值"。

拆开看只有两步：

1. **给函数签名起一个类型名**：`type HandlerFunc func(ResponseWriter, *Request)`。
   Go 允许对任何自定义类型定义方法， **函数类型也不例外**（和 `type MyInt int` 一样）。
2. **让这个函数类型实现接口**：给 `HandlerFunc` 定义 `ServeHTTP` 方法，方法体只做一件事：调用自身 `f(w, r)`。
   于是 `HandlerFunc` 就满足了 `Handler` 接口。

## 解决什么问题

`http.Handle` 只收 `Handler` 接口。如果没有 `HandlerFunc`，每写一个处理逻辑都得先定义一个 struct 再实现方法：

```go
type helloHandler struct{}

func (helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
w.Write([]byte("hi"))
}

http.Handle("/", helloHandler{})
```

有了 `HandlerFunc`，一个普通函数做一次 **类型转换**就变成 `Handler`，不需要 struct：

```go
func hello(w http.ResponseWriter, r *http.Request) {
w.Write([]byte("hi"))
}

http.Handle("/", http.HandlerFunc(hello)) // 类型转换，不是函数调用
```

注意 `http.HandlerFunc(hello)` 是 **类型转换**（和 `int64(x)` 一个语法），不是调用 `hello`。
`hello` 的签名和 `HandlerFunc` 的底层类型一致，所以可以直接转换。
`http.HandleFunc` 就是在内部帮你做了这一步：

```go
func (mux *ServeMux) HandleFunc(pattern string, handler func (ResponseWriter, *Request)) {
mux.Handle(pattern, HandlerFunc(handler))
}
```

## 为什么值得学

- **接口和函数之间架了一座桥**：接口是 Go 的抽象手段，函数是最轻量的行为单元。适配器让两者可以互相替换，调用方只依赖接口，实现方可以给
  struct、也可以给一个函数或闭包。
- **中间件（middleware）的基础**：中间件的标准签名 `func(http.Handler) http.Handler`，返回值通常就是一个 `http.HandlerFunc`
  包裹的闭包。没有这个适配器，每层中间件都要新建 struct。
- **与 [[functional.md]] 里"函数是一等公民"呼应**：闭包能捕获状态，再套一层 `HandlerFunc` 转换，就得到一个带状态的接口实现。

## 中间件示例

```go
func logging(next http.Handler) http.Handler {
return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
log.Println(r.Method, r.URL.Path)
next.ServeHTTP(w, r)
})
}

http.Handle("/", logging(http.HandlerFunc(hello)))
```

## 自己写一个

模式是通用的，任何单方法接口都能这么做，三步走：

```go
// 1. Define the Single-Method Interface
type Greeter interface {
Greet(name string) string
}

// 2. Define the Function Type Adapter
type GreeterFunc func (string) string

// 3. Make the function type implement the interface
func (f GreeterFunc) Greet(name string) string {
return f(name) // Call the underlying function itself
}
```

之后一个普通函数或闭包做一次类型转换，就是一个 `Greeter`：

```go
func english(name string) string { return "Hello, " + name }

var g Greeter = GreeterFunc(english) // 普通函数
var h Greeter = GreeterFunc(func (n string) string {         // 闭包
return "你好，" + n
})
```

多方法接口就不适合了，因为一个函数只能对应一个方法。

## 标准库里的同类例子

- `net/http`：`HandlerFunc` / `Handler`（本篇）
- `sort`：`sort.Slice` 内部用函数适配 `Less`
- `strings.NewReplacer` 之外，`io` 包里也有类似手法把函数包装成 `io.Reader` / `io.Writer`（第三方常见 `WriterFunc`）
- `fs.WalkDirFunc`、`flag.Func`：都是"函数类型 + 方法"的变体

## 一句话总结

**给函数类型定义一个"转发给自己"的方法，让函数不经 struct 就能满足接口。**
