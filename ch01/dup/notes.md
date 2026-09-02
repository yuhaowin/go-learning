# 学习笔记

## os.Stdin API

`os.Stdin` 是 `os` 包提供的标准输入抽象，本质是一个已经打开的 `*os.File`。

### 基本定义

```go
var (
	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)
```

三兄弟对应 Unix 的三个标准文件描述符：`0`（stdin）、`1`（stdout）、`2`（stderr）。类型都是 `*os.File`，所以 `os.File` 有的方法 `os.Stdin` 都能用。

### os.File 常用方法（都可用在 os.Stdin 上）

| 方法 | 作用 |
|---|---|
| `Read(b []byte) (n int, err error)` | 实现了 `io.Reader` 接口，最底层的读取方式 |
| `Close() error` | 关闭文件描述符（一般不需要手动关闭 Stdin） |
| `Fd() uintptr` | 返回底层文件描述符（整数） |
| `Name() string` | 返回文件名，比如 `/dev/stdin` |
| `Stat() (FileInfo, error)` | 获取文件信息，可用来判断是终端输入还是管道输入 |

### 因为实现了 io.Reader，常见配套用法

**1. `bufio.Scanner`（dup1/dup2 用的方式）—— 按行/按词读取**

```go
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
	line := scanner.Text()
}
```

**2. `bufio.Reader` —— 更细粒度控制**

```go
reader := bufio.NewReader(os.Stdin)
line, err := reader.ReadString('\n')
```

**3. `fmt.Scan` / `Fscan` 系列 —— 格式化解析**

```go
var name string
var age int
fmt.Scan(&name, &age)
```

**4. `io.ReadAll` —— 一次性读取全部内容**

```go
data, err := io.ReadAll(os.Stdin)
```

**5. `io.Copy` —— 直接转发到别处**

```go
io.Copy(os.Stdout, os.Stdin) // 简单的 echo 管道
```

### 判断输入来源（终端 vs 管道/文件）

```go
stat, _ := os.Stdin.Stat()
if (stat.Mode() & os.ModeCharDevice) == 0 {
	// 数据来自管道或重定向文件
} else {
	// 数据来自终端交互输入
}
```

### 小结

`os.Stdin` 本身 API 很简单（就是个 `*os.File`），真正好用的能力来自它实现了 `io.Reader` 接口，可以直接"插"进 `bufio.Scanner`、`bufio.Reader`、`io.ReadAll`、`io.Copy` 等标准库工具里，这也是 Go 里 `io.Reader`/`io.Writer` 接口设计的典型体现——各种 I/O 组件可以互相组合。

参考代码位置：`ch1/dup1/main.go:11`

## map 和 struct 的值传递区别

记录在 `ch1/dup2/main.go:32-37` 的源码注释中：

- **struct**（如 `os.File`）是值类型，传值会复制整个结构体，函数内部拿到的是独立副本，修改不影响外部，所以需要传指针（如 `*os.File`）。
- **map** 是引用类型，变量内部是指向底层哈希表的描述符（含指针），传值时只复制描述符，底层数据仍是同一份，所以函数参数不需要传指针，函数内部对它的增删改外部也能看到。

## go run 时如何传入 os.Stdin

`go run` 运行时，`os.Stdin` 就是运行该命令的终端会话的标准输入，程序本身不需要做任何特殊操作去"传入"，只需要在执行命令时用外部方式喂数据：

**1. 交互式手动输入**

```bash
go run main.go
```
直接在终端敲字，回车换行，输完后按 **Ctrl+D**（Mac/Linux，代表 EOF）结束输入。

**2. 重定向文件**

```bash
go run main.go < sample.txt
```
`<` 把文件内容接到标准输入上。

**3. 管道**

```bash
echo -e "a\nb\na" | go run main.go
cat sample.txt | go run main.go
```
`|` 把左边命令的标准输出接到右边命令的标准输入。

### 原理

这三种方式的本质是在 **shell 层面**决定文件描述符 `0`（stdin）连接到什么地方——终端、文件、还是另一个进程的输出管道。Go 程序里的 `os.Stdin` 只是去读这个由 shell 准备好的文件描述符 `0`，代码本身不感知也不需要关心数据来源。

严格来说不存在"给 `go run` 传入 os.Stdin"这个操作——`os.Stdin` 永远存在（连到某个数据源），能控制的只是这个数据源具体是什么，且这与 `go run` 本身无关（`go build` 编译出的二进制文件同理）。
