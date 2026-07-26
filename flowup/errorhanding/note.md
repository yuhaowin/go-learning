defer 调用

确保调用在函数结束时发生

参数在defer语句时计算

defer列表为后进先出


这是 Go 的**类型别名(type alias)**语法,和普通的类型定义(type definition)不一样。

## 两种写法的区别

```go
type PathError = fs.PathError   // 类型别名(带等号)
type PathError fs.PathError     // 类型定义(不带等号)
```

### 1. 类型别名(`type A = B`)

`PathError` 和 `fs.PathError` 是**同一个类型**,只是多了一个名字,完全可以互换使用。就像给一个人起了个外号,但外号和真名指的是同一个人。

```go
type PathError = fs.PathError

var e PathError = fs.PathError{}
var e2 fs.PathError = e  // 完全没问题,因为是同一个类型
```

### 2. 类型定义(`type A B`,不带等号)

这是创建一个**新的、独立的类型**,底层类型(underlying type)虽然和 `fs.PathError` 一样,但它是一个不同的类型,不能直接互相赋值,需要显式转换。

```go
type PathError fs.PathError  // 新类型

var e PathError = PathError(fs.PathError{})  // 需要显式转换
var e2 fs.PathError = fs.PathError(e)         // 也需要显式转换
```

## 为什么标准库要用别名

在 Go 标准库里,你能看到这样的代码(比如在 `os` 包里):

```go
package os

type PathError = fs.PathError
```

这是历史原因造成的:`PathError` 最早是在 `os` 包里定义的,后来 Go 团队把这个类型挪到了新的 `io/fs` 包里(Go 1.16 引入 `io/fs`),但是又想保持向后兼容——不想让所有用 `os.PathError` 的老代码全部报错或者需要修改。

于是他们用 `type PathError = fs.PathError` 这种别名写法,让 `os.PathError` 和 `fs.PathError` 变成**完全相同的类型**。这样:

- 老代码里用 `os.PathError` 的地方,继续可以正常工作
- 新代码可以直接用 `fs.PathError`
- 两者可以无缝互换、比较、赋值,因为它们本质上是同一个类型,只是路径(包名)不同

## 一个简单的判断方法

看到 `type A = B` 这种**带等号**的写法,就知道是"起别名",两个名字实际上引用的是同一底层类型,可以完全互换。这在 Go 做 API 迁移、包重构、保持向后兼容时很常见。


### panic

停止当前函数执行

一直向上返回，执行每一层defer

如果没有遇见recover，程序退出

### recover

仅在defer调用中使用

获取panic的值

如果无法处理，可重新panic