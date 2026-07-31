package main

import (
	"fmt"
	"time"
)

// 协程 Coroutine 轻量级“线程”
// 非抢占式多任务处理，由协程主动交出控制权，由于非抢占式，才能做到轻量，主动交出控制权时，可确保需要保存的上下文信息比较少。
// 是编译器/解释器/虚拟机层面的多任务，不是操作系统层面的多任务。在 go 中，会有调度器调度协程。

// go build 出来的 binary，内嵌了一整套 Go runtime —— goroutine 调度器、垃圾回收器(GC)、内存分配器全都在这一个文件里。

func goroutine1() {
	for i := range 10 {
		go func(i int) {
			for {
				fmt.Printf("Hello form goroutine %d\n", i) // fmt.Printf() 是 IO 操作，会使得 goroutine 主动交出控制权。
			}
		}(i)
	}
	time.Sleep(time.Minute)
}

func goroutine2() {
	arr := [10]int{}
	for i := range 10 {
		go func(i int) {
			for {
				arr[i]++ // cpu 操作，不会交出控制权。(小于Go 1.14 )
			}
		}(i)
	}
	time.Sleep(time.Second) // main 本身也跑在一个 goroutine 上，由于没有人交出控制权，该语句也得不到执行。(小于Go 1.14 )
	fmt.Println(arr)
}

// 不越界的原因是 Go 1.22 起改变了 for 循环变量的作用域语义。
//
// 关键点：for i := 0; i < 10; i++ { go func() { ... arr[i] ... }() } 里的 i，在 Go 1.22 之前是整个循环共享一个变量——所有 10 个 goroutine 闭包捕获的是同一个 i 的地址。由于 goroutine 是异步调度的，等它们真正执行到
// arr[i]++ 时，主 goroutine 的 for 循环可能已经把 i 增到 10（循环条件判断失败退出时的值），这时候读到的 i 就是 10，arr[10] 越界 panic。而且这是对同一个变量的无同步并发读写，本身也是 data race。
//
// Go 1.22 把这个坑改掉了：每次循环迭代都会创建一个新的 i 副本，闭包捕获的是"当次迭代"那个独立的变量，值固定为 0、1、2...9，不会再被后续迭代覆盖。所以现在这段代码里每个 goroutine 拿到的 i 永远是自己创建时的那个值，arr[i] 永远落在 [0,9]，不会越界。
func goroutine3() {
	arr := [10]int{}
	for i := range 10 {
		go func() {
			for {
				arr[i]++ // 这里不会发生数组越界了。
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(arr)
}

func goroutine4() {
	arr := [10]int{}
	var i int
	for i = 0; i < 10; i++ { // i 是循环共享变量，一边读，一边写，会出现 race condition
		// i 是循环共享变量，等 goroutine 真正跑到 arr[i]++ 时 i 已经是 10，arr[10] 越界 panic。
		go func() { arr[i]++ }() // 会 out of index
	}
	time.Sleep(time.Second)
}

// - for i = 0; i < 10; i++：循环体跑完 i=9 那次之后，还会再执行一次 i++ 把 i 变成 10，然后条件 i<10 才为假、退出循环。所以循环结束后共享变量 i 停在 10——一个越界的下标。等 goroutine 真正运行到 arr[i]++ 时大概率读到这个 10，panic。
// - for i = range 10：range-over-int 语义不一样，它只是把 0..9 依次赋给 i，最后一次迭代之后不会再多"进一"一步。循环结束后共享变量 i 停在 9——依然是个合法下标。所以就算 goroutine 全部在循环结束后才执行，读到的也是 9，不会越界。
func goroutine5() {
	arr := [10]int{}
	var i int
	for i = range 10 {
		go func() { arr[i]++ }()
	}
	time.Sleep(time.Second)
}

func main() {
	//goroutine1()
	//goroutine2()
	//goroutine3()
	//goroutine4()
	goroutine5()
}
