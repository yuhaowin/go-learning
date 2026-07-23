# 学习笔记：接口与指针接收者

## 问题现象

```go
r = real.Retriever{
    UserAgent: "Mozilla/5.0",
    Timeout:   10,
}
```

编译报错：

```
cannot use real.Retriever{…} (value of struct type real.Retriever) as Retriever value in assignment:
real.Retriever does not implement Retriever (method Get has pointer receiver)
```

## 原因：方法集规则

`real.go` 中 `Get` 方法用的是指针接收者：

```go
func (r *Retriever) Get(url string) string
```

Go 的方法集规则：

- 类型 `T` 的方法集：只包含**接收者是 `T`** 的方法
- 类型 `*T` 的方法集：包含**接收者是 `T` 或 `*T`** 的所有方法

因此：

- `real.Retriever`（值类型）方法集里**没有** `Get` → 不满足 `Retriever` 接口
- `*real.Retriever`（指针类型）方法集里**有** `Get` → 满足 `Retriever` 接口

## 修复

```go
r = &real.Retriever{
    UserAgent: "Mozilla/5.0",
    Timeout:   10,
}
```

`&T{...}` 是对复合字面量取地址，Go 特别允许这种写法（复合字面量本身不可寻址，但 `&T{...}` 是例外），等价于：

```go
tmp := real.Retriever{UserAgent: "Mozilla/5.0", Timeout: 10}
r = &tmp
```

得到一个 `*real.Retriever`，其方法集包含 `Get`，可以赋值给接口变量 `r`。

## 为什么 `r`（声明为接口类型，非指针）能装下一个指针？

```go
var r Retriever
```

`r` 的**静态类型**是接口 `Retriever`。接口变量在底层可以理解为一个装着两部分信息的"盒子"：

```
(动态类型, 动态值)
```

- `r = mock.Retriever{...}` → 盒子里是 `(mock.Retriever, {Contents:"..."})`，动态类型是值类型
- `r = &real.Retriever{...}` → 盒子里是 `(*real.Retriever, 0xc0001...)`，动态类型是指针类型

接口只关心动态类型是否满足方法集要求（这里是 `Get(string) string`），不关心动态类型本身是值还是指针。调用 `r.Get(...)` 时，Go 根据盒子里存的动态类型分派到对应的方法实现。

这也是为什么 `download(r Retriever)` 完全不用关心调用方传进来的是值还是指针，只要满足接口即可。

## 小结

- 值接收者方法：值类型和指针类型都能调用，两者的方法集都包含该方法
- 指针接收者方法：只有指针类型的方法集包含该方法，值类型不包含
- 接口变量能存放任何满足其方法集的具体类型（值或指针），接口本身的声明形式与存入的是值还是指针无关
