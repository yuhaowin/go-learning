# 学习笔记

## 无类型常量（untyped constant）

Go 里不写类型、直接用字面量声明的常量是**无类型常量**，跟"有类型常量/变量"最大的区别是：它没有固定类型，会根据使用它的表达式上下文**自动适配**成需要的类型，不需要显式转换。

```go
const Typed int = 100   // 有类型常量，类型固定为 int
const Untyped = 100     // 无类型常量，没有固定类型
```

### 默认类型

没有上下文强制转换时，无类型常量会使用一个"默认类型"：

| 字面量形式 | 默认类型 |
|---|---|
| `100` | `int` |
| `3.14` | `float64` |
| `'A'` | `rune` (int32) |
| `"hello"` | `string` |
| `true` | `bool` |
| `1 + 2i` | `complex128` |

### 核心特性：按上下文自动适配

```go
const size = 100

var a int32 = size        // size 在这里变成 int32
var b float64 = size      // size 在这里变成 float64
var c int64 = size * 3    // 整个表达式都在 int64 上下文里计算

fmt.Println(math.Sqrt(size)) // size 变成 float64，因为 Sqrt 需要 float64 参数
```

如果 `size` 是 `var size = 100`（有类型变量，固定为 `int`），上面这些赋值全部会编译报错，因为 Go 不允许不同具体类型之间的隐式转换。

### 为什么这样设计

1. **精度不截断**：常量运算在编译期用任意精度计算，不受具体类型宽度限制，比如 `const Big = 1 << 100`，只要不真正赋给某个类型的变量就不会溢出报错。
2. **写代码更自然**：可以写 `var x float64 = 1`，不用写成 `1.0`，因为无类型整数常量 `1` 能自动转成 `float64`。
3. **常量之间可以自由混合运算**：`const Pi = 3.14`、`const TwoPi = 2 * Pi`，`2` 和 `Pi` 默认类型不同（int 和 float64）也能直接相乘，因为都还没"定型"。

### 什么时候会"定型"

- 赋值给一个有类型的变量：`var x int32 = size`
- 参与到一个已经有类型的表达式运算中
- 用 `:=` 声明时：`x := size` 会让 `x` 采用 `size` 的**默认类型**（这里是 `int`），从此 `x` 就固定为 `int` 了

### 实例：ch1/lissajous 里 const 改成 var 会编译报错

`lissajous` 函数里原本用 `const` 声明了一组配置值：

```go
const (
    cycles  = 5
    res     = 0.001
    size    = 100
    nframes = 64
    delay   = 8
)
```

`size` 在 `2*size+1`（int 上下文）里被当成 `int`，在 `x*size`（`x` 是 `math.Sin` 返回的 `float64`）里被当成 `float64`，两处互不冲突，因为它是无类型常量。

如果把 `const` 误改成 `var`：

```go
var (
    cycles  = 5     // 定型为 int
    size    = 100   // 定型为 int
    ...
)
```

`size` 从此固定为具体类型 `int`，再用在 `x*size`（`float64` 上下文）里就会报错：

```
./main.go:41:30: math.Pi (untyped float constant 3.14159) truncated to int
./main.go:44:31: invalid operation: x * size (mismatched types float64 and int)
./main.go:44:53: invalid operation: y * size (mismatched types float64 and int)
```

- 41 行 `cycles*2*math.Pi`：`cycles` 定型为 `int` 后，`math.Pi` 这个无类型浮点常量被迫转成 `int`，发生截断报错。
- 44 行 `x*size`：`x` 是具体类型 `float64`，`size` 定型为具体类型 `int`，两个不同具体类型不能隐式相乘。

**结论**：用 `const` 而非 `var` 声明这类"配置值"，不仅语义上表达"不可变"，还能利用无类型常量按上下文自动适配的特性，避免类型不匹配问题。

参考代码位置：`ch1/lissajous/main.go:27-33`

## var (...) / const (...) 分组声明

`var (...)` 和 `const (...)` 是 Go 的**分组声明（grouped declaration）**语法，用括号把多个变量/常量声明放在一起，避免重复写 `var`/`const` 关键字。

### 单个声明 vs 分组声明

```go
// 单个声明，重复写关键字
var cycles = 5
var res = 0.001
var size = 100

// 分组声明，等价效果，只写一次关键字
var (
    cycles = 5
    res    = 0.001
    size   = 100
)
```

两种写法**完全等价**，只是分组写法更简洁，也更方便对齐注释，比如：

```go
const (
    cycles  = 5     // number of complete x oscillator revolutions
    res     = 0.001 // angular resolution
    size    = 100   // image canvas covers [-size..+size]
    nframes = 64    // number of animation frames
    delay   = 8     // delay between frames in 10ms units
)
```

### 常见用法：const 分组 + iota

`const (...)` 分组常搭配 `iota` 自动生成一组递增的常量，是 Go 里定义"枚举"的惯用方式：

```go
const (
    whiteIndex = iota // 0
    blackIndex        // 1
)
```

`ch1/lissajous/main.go:16-19` 的 `whiteIndex = 0` / `blackIndex = 1` 也可以用 `iota` 简化成上面这样，效果一样。

### import 也是同样的语法

这个分组写法不止 `var`/`const`，`import` 也用同样的模式：

```go
import (
    "fmt"
    "os"
)
```

**小结**：`var (...)` / `const (...)` 是"批量声明"的语法糖，把多条同类型的声明语句合并到一个块里，纯粹是为了代码整洁，不影响语义。

参考代码位置：`ch1/lissajous/main.go:16-19`、`ch1/lissajous/main.go:27-33`

## 有没有无类型的变量（var）？

没有。**`var` 声明的变量永远有确定、具体的类型**，"无类型（untyped）"这个概念只存在于**常量**身上，变量不存在无类型的说法。

### 为什么变量必须有具体类型

变量是要在内存里实际存储的东西——编译器需要知道它占几个字节、怎么解释这些字节（`int32` 是 4 字节整数编码，`float64` 是 8 字节浮点编码……），所以每个变量在编译期就必须绑定一个确定的类型，运行时不能再变。

常量则不同，它在编译期只是一个"值"，不占用内存、不需要立即确定表示方式，所以可以保持"无类型"状态，直到真正用到某个类型的地方才临时转换。

### var 的类型从哪来

```go
var a int32 = 100     // 显式指定类型：a 就是 int32

var b = 100            // 没写类型，编译器用初始值推导
                        // 100 是无类型常量，默认类型是 int
                        // 所以 b 被推导为 int，从此固定

var c = size           // size 若是 const 无类型常量，c 会取 size 的默认类型
                        // size 若是 var 变量，c 直接沿用 size 的具体类型
```

不管哪种写法，`var` 声明完成的那一刻，变量的类型就**确定并锁死**了，之后不能再像无类型常量那样"随上下文变化"。

### 对比小结

| | const | var |
|---|---|---|
| 有类型形式 | `const x int = 100` | `var x int = 100` |
| 无类型/推导形式 | `const x = 100` → **无类型常量**，用到哪就适配哪个类型 | `var x = 100` → **有类型变量**，类型在声明时就定死为 `int` |

`const x = 100` 和 `var x = 100` 看起来很像，但结果完全不同：前者 `x` 是无类型常量，后者 `x` 是类型已经固定为 `int` 的变量——这正是上面 `const` 改成 `var` 会编译报错的根本原因。

## slice 能不能声明为 const？

不能。`const` **只能声明基本类型（布尔、数字、字符串）的值**，slice 不行。

```go
const a = []int{1, 2, 3}  // 编译错误：const initializer []int{…} is not a constant
```

### 为什么不行

`const` 声明的常量必须是**编译期就能确定、且不占用运行时内存分配**的值——纯粹的数字、字符串、布尔值可以直接以字面形式"刻"进编译产物里。

而 slice 本质上是一个运行时结构体（指针 + 长度 + 容量），`[]int{1, 2, 3}` 这行代码实际上要在**运行时**：
1. 分配一个底层数组
2. 把 1、2、3 写进去
3. 构造出包含指针的 slice header

这个过程涉及内存分配和指针，根本不是"编译期能直接确定"的值，所以 Go 不允许它作为常量。

### Go 里能用 const 的类型只有这些

- 布尔型：`bool`
- 数字类型：`int`、`float64`、`complex128` 等
- 字符串：`string`

**不能** const 的类型：slice、map、struct、array、pointer、interface、channel、function 等所有需要运行时构造或包含引用/指针语义的类型。

### 想要"不可变的 slice" 怎么办

Go 没有真正的"只读 slice"机制，惯用做法是用 `var` 声明包级变量，靠约定（不去修改它）来保证不变性：

```go
var palette = []color.Color{color.White, color.Black}  // ch1/lissajous/main.go:14 就是这么写的
```

如果确实需要编译期常量且是"一组值"，通常改用 `const (...)` 分组声明多个独立常量，而不是一个 slice：

```go
const (
    whiteIndex = 0
    blackIndex = 1
)
```
