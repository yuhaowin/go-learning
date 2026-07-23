package main

import (
	"fmt"
	. "net" // 点导入（dot import） 意思是：把 net 包里所有导出的标识符（类型、函数、常量等）直接导入到当前文件的作用域，使用时不需要写包名前缀。
)

//  import "net"
//  使用时要写: net.Flags, net.FlagUp
//
//  import . "net"
//  使用时直接写: Flags, FlagUp

//  注意：
//  - 这种写法在实际工程中不推荐，因为它把包的所有导出名字都塞进当前作用域，容易和其他标识符冲突，也让读代码的人搞不清楚 Flags、FlagUp 到底来自哪个包，可读性变差。
//  - 常见用途仅限于：测试文件里为了简化调用（比如 . "some/testpackage"），或者像这里一样是《The Go Programming Language》教材里为了演示简洁而故意这么写的示例代码。

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776             (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424    (exceeds 1 << 64) 溢出
	YiB // 1208925819614629174706176 溢出
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {

	fmt.Println(YiB / ZiB) // 不会溢出

	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}
