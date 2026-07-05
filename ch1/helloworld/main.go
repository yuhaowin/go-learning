//这里 package main 是 Go 语言的约定，作用是告诉编译器：这个包要被编译成一个可独立执行的程序，而不是一个供别人 import 的库。
//
//具体原因：
//
//1. Go 用包名 main 作为特殊标记——编译器看到包声明是 main 时，会认为这是程序的入口包，而不是普通库包（普通库包名通常和目录名相关，比如 package helloworld）。
//2. 必须配合 func main() 使用——package main 里必须有一个无参数、无返回值的 main() 函数，这个函数就是程序运行的起点，类似 Java 里的 main 入口。
//3. 可执行 vs 库的区别：
//- package main → go build/go run 会生成一个可执行文件（如 go run ch1/helloworld/main.go 能直接跑）。
//- package xxx（非 main）→ 只能被其他代码 import，自身不能单独运行，也不会生成可执行文件。

//go 文件名称不需要一定叫 main.go Go 语言的入口判断只看两点：
//
//1. 包声明是 package main
//2. 该包内有 func main() 函数
//
//文件名任意，不影响是否是程序入口。 比如你把 ch1/helloworld/main.go 改名成 hello.go、app.go 甚至 xyz.go，只要文件里还是 package main + func main()，一样可以 go
//build/go run 生成可执行程序。
//
//之所以大家习惯叫 main.go，只是社区约定俗成，方便一眼看出这是入口文件，Go 工具链本身并不强制要求这个文件名。
//
//补充几点相关规则：
//- 同一个目录下的所有 .go 文件必须属于同一个包（这里都得是 package main）。
//- 一个目录（也就是一个包）里只能有一个 func main()，多了会编译报错（重复定义）。
//- 如果目录里有多个文件，func main() 放在哪个文件里都行，不要求非在叫 main.go 的文件里。

package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
