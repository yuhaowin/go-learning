# var 声明变量 vs type 声明类型

```go
var x float64   // 声明一个变量，类型是 float64
type y float64  // 声明一个新类型，底层类型是 float64
```

- `var x float64` 声明的是**变量**，可以直接赋值、参与运算。
- `type y float64` 声明的是**新的类型**，`y` 本身不是变量，底层虽然是 `float64`，
  但 Go 认为它和 `float64` 是不同类型，不能直接混用，必须显式转换：

```go
type y float64

var v y = 3.14
var f float64 = 2.0

// v + f        // 编译错误：类型不匹配
v + y(f)         // 正确，需要显式转换
```

常用于给类型附加方法，增加类型安全性：

```go
type Celsius float64
type Fahrenheit float64

func (c Celsius) ToFahrenheit() Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}
```

这样 `Celsius` 和 `Fahrenheit` 不能直接相加，避免摄氏度、华氏度混用出错。

## 类型天然有构造函数吗？

没有。Go 没有构造函数这个语言机制，类型天然自带的是**零值（zero value）**：

```go
var c Celsius   // 自动是 0
var s string    // 自动是 ""
var p *int      // 自动是 nil
```

设计理念：让零值本身就是可用、合理的状态。

"构造函数"只是社区约定，写一个普通函数，通常叫 `New` 或 `NewXxx`：

```go
func NewPoint(x, y int) Point {
    return Point{X: x, Y: y}
}
```

编译器不强制调用它，可以绕过直接用字面量或零值。

## Fahrenheit(c\*9/5 + 32) 为什么可以这样用

这里叠加了两件事：

1. **`T(v)` 是类型转换**，不是函数调用。当 `T` 是类型名时，`T(v)` 表示把 `v`
   转换成类型 `T`，前提是底层类型相同或兼容。

2. **无类型常量的隐式适配**。`9`、`5`、`32` 是无类型常量，参与运算时会隐式转换成
   对方的类型去匹配，所以 `c*9/5 + 32` 的结果类型仍然是 `Celsius`。

   如果换成有类型的变量就会报错：

   ```go
   var x float64 = 9
   c * x   // 编译错误！Celsius 和 float64 是不同类型
   ```

整个表达式的过程：

1. `c*9/5 + 32` → 结果是 `Celsius` 类型的值
2. `Fahrenheit(...)` → 把这个 `Celsius` 值转换成 `Fahrenheit` 类型
3. 返回

## type y float64 这种定义的目的是什么（底层类型明明相同）

核心目的不是"数据不同"，而是给这个数据赋予**独立的身份和行为**，尽管底层存储完全一样。

**1. 类型安全，防止逻辑上不该混合的值被混用**

```go
type Celsius float64
type Fahrenheit float64
type Meter float64
```

编译器把它们当作不同类型，`Celsius` 和 `Meter` 就算数值一样也不能相加、不能互相赋值。
能防住"单位搞混了"这种典型 bug（比如把摄氏度当英尺用）。裸的 `float64` 编译器毫无察觉。

**2. 可以给它定义方法 —— 这是最关键的一点**

Go 里不能给内置类型（`float64`、`int`、`string` 等）直接加方法：

```go
func (f float64) IsFreezing() bool { ... }  // 编译错误！
```

但定义了自己的类型之后就可以：

```go
type Celsius float64

func (c Celsius) IsFreezing() bool {
    return c <= 0
}

func (c Celsius) String() string {
    return fmt.Sprintf("%g°C", float64(c))
}
```

有了方法，就能让这个类型**实现接口**（比如 `fmt.Stringer`、`sort.Interface` 等），
获得完整的"对象"行为，而 `float64` 本身做不到。

**3. 表达意图，自文档化**

`func Boil() Celsius` 比 `func Boil() float64` 传达的信息多——调用者一看就知道
返回的是摄氏度，不用猜单位或看注释。

**4. 零运行时开销**

`Celsius` 编译后和 `float64` 内存布局完全相同，转换 `Celsius(x)` 是编译期操作，
没有额外运行时成本，不像结构体包装那样有间接层。

标准库里 `time.Duration`（底层 `int64`）、`http.Header`（底层 `map[string][]string`）
都是用这种方式定义的。
