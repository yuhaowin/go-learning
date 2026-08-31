// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			// Ready when someone is blocked on deposits <- amount (i.e. Deposit was called).
			balance += amount
		case balances <- balance:
			// Ready when someone is blocked on <-balances (i.e. Balance was called);
			// an unbuffered channel only completes a send when a receiver is
			// simultaneously waiting, so this case rendezvous with that receive
			// and hands over the current balance. Empty body: the send is the
			// whole point, nothing else to do.
		}
	}
}

// init() 是 Go 里的一个特殊函数，会自动执行，不需要手动调用。关键特性：
//
// - 触发时机：包被导入（import）时，Go runtime 会自动调用该包里所有的 init() 函数，发生在 main() 执行之前
// - 执行顺序：先初始化包级变量（比如这里的 deposits、balances），再执行 init()
// - 不能手动调用：不能在代码里写 bank.init()，也不能取它的地址，编译器不允许
// - 可以有多个：一个文件甚至一个包里可以定义多个 init()，会按源文件名的字典序、以及文件内出现的顺序依次执行
// - 没有参数也没有返回值：签名固定是 func init()
//
// 在这段代码里的作用：当外部代码 import 了 bank 这个包（比如 bank_test.go 里的 import bank "github.com/.../bank/version1"），init() 会自动跑起来，启动 teller() 这个 goroutine 在后台常驻监听 deposits/balances 两个
// channel。这样调用方不需要显式调用任何"启动"函数，一 import 包，这个"账户管家"协程就已经在运行了。
func init() {
	go teller() // start the monitor goroutine
}
