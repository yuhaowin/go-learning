package main

import (
	"runtime"
	"testing"
	"time"
)

// 验证 Go 1.14+ 的异步抢占：即使 GOMAXPROCS(1) 把所有 goroutine 挤到一个核上，
// 纯 CPU 循环（没有函数调用/协作点）也不会永久占住线程，main 的 time.Sleep 依然能被唤醒。
func TestAsyncPreemption(t *testing.T) {
	prev := runtime.GOMAXPROCS(1)
	defer runtime.GOMAXPROCS(prev)

	arr := [10]int{}
	for i := range 10 {
		go func(i int) {
			for {
				arr[i]++
			}
		}(i)
	}

	time.Sleep(1 * time.Second)

	//  用 GOMAXPROCS(1) 模拟单核最坏情况，起 10 个纯 CPU 死循环 goroutine，验证 1 秒后每个 goroutine 都被调度过至少一次（arr[i] != 0），实测证明 Go 1.14+ 的异步抢占确实会打断没有函数调用的纯 CPU 循环。
	for i, v := range arr {
		if v == 0 {
			t.Errorf("arr[%d] = 0，goroutine %d 从未被调度到，异步抢占可能未生效", i, i)
		}
	}
}
