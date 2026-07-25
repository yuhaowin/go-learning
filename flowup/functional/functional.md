

函数是一等公民：参数，变量，返回值都可以是函数


“正统”函数式编程
+ 不可变性：不能有状态，只有常量和函数
+ 函数只能有一个参数


## 函数类型 与 闭包（adder 示例）

```go
func adder() func(int) int {
	sum := 0 // 自由变量，闭包
	return func(v int) int {
		sum += v
		return sum
	}
}
```

### `func(int) int` 是什么

`func(int) int` 是一个**函数类型**，表示“接收一个 int 参数、返回一个 int 值的函数”，描述的是函数签名（参数类型 + 返回值类型），不关心函数名字。

函数类型和 `int`、`string` 一样可以：
- 作为变量类型
- 作为参数类型
- 作为返回值类型

### 拆解 adder

- `adder` 本身不接收参数
- `adder` 的返回值类型是 `func(int) int`，也就是说 `adder()` 返回的不是普通值，而是**另一个函数**
- `return func(v int) int {...}` 返回一个匿名函数（函数字面量），它的类型正好是 `func(int) int`，与 `adder` 声明的返回类型匹配

### 闭包

匿名函数内部引用了外部的局部变量 `sum`，这就是**闭包**：内部函数“捕获”了 `sum`。即使 `adder()` 执行完毕，正常局部变量该被销毁，但因为内部函数还持有对 `sum` 的引用，Go 会把 `sum` 分配到堆上，让它的生命周期跟随这个闭包延续。

### 用法示例

```go
a := adder()       // a 类型是 func(int) int，内部绑定了它自己的 sum
fmt.Println(a(1))  // sum: 0+1=1  -> 输出 1
fmt.Println(a(2))  // sum: 1+2=3  -> 输出 3
fmt.Println(a(3))  // sum: 3+3=6  -> 输出 6

b := adder()        // 重新调用 adder()，产生一个全新、独立的 sum
fmt.Println(b(10))  // sum: 0+10=10 -> 输出 10，不受 a 影响
```

每次调用 `adder()` 都会创建一个**新的、独立的** `sum`，返回一个绑定该 `sum` 的新闭包，因此 `a` 和 `b` 各自维护自己的累加状态，互不干扰。