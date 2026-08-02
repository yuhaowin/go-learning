# select + nil channel 学习笔记

## 代码回顾（main.go）

```go
func main() {
	var c1, c2 chan int // c1 and c2 = nil
	select {
	case n := <-c1:
		fmt.Println("Received form c1:", n)
	case n := <-c2:
		fmt.Println("Received form c2:", n)
	}
}
```

## 现象

不是编译错误，运行时报：

```
fatal error: all goroutines are asleep - deadlock!
```

## 原因

- `var c1, c2 chan int` 只声明未 `make`，值是 **nil channel**。
- 对 nil channel 做发送/接收会 **永久阻塞**（Go 规范定义的行为，既不 panic 也不返回）。
- `select` 里所有 `case` 都是 nil channel，又没有 `default`，两个 case 永远不会就绪 → `select` 永久阻塞。
- `main` 是唯一在跑的 goroutine，永久阻塞且无人能唤醒 → runtime 判定死锁，直接 fatal error 终止（无法用 `recover` 捕获，因为不是
  panic）。

## 关键知识点：nil channel 在 select 中的作用

nil channel 在 `select` 里会被当作**"永不就绪"**的 case，这是一个常见技巧：可以在运行时把某个 channel 变量置为 `nil`
，从而动态"关闭"/禁用 select 里对应的分支，而不需要真正 close 掉 channel。

但前提是： **至少要有一个 case 是真实可用的 channel，或者要有 `default` 分支**，否则全员 nil 就会死锁。

## 修复方式

1. 让 `c1`、`c2` 真正 `make` 出来，并有其他 goroutine 往里发送数据；或
2. 加 `default` 分支避免永久阻塞：

```go
select {
case n := <-c1:
	fmt.Println("Received from c1:", n)
case n := <-c2:
	fmt.Println("Received from c2:", n)
default:
	fmt.Println("no channel ready")
}
```

## 与前面几个例子的对比（同类"运行时而非编译期"问题）

| 例子                                        | 位置                                  | 问题类型                      | 触发原因                                                          |
|---------------------------------------------|---------------------------------------|-------------------------------|-------------------------------------------------------------------|
| `chan<- int` 命名返回值内部 `<-c`           | flowup/channel/channel.go             | **编译错误**                  | 单向 channel 类型静态检查，receive-only 操作用在 send-only 类型上 |
| `wg := sync.WaitGroup`                      | flowup/channel/waitgroup/waitgroup.go | **编译错误**                  | 类型名当值用，`:=` 右边必须是表达式                               |
| `close(out)` 放在回调里，每个节点都执行一次 | flowup/tree/traversal.go              | **运行时 panic**              | 对已关闭 channel 重复 send / 重复 close                           |
| `close(out)` 紧跟在 `go func(){...}()` 后面 | flowup/tree/traversal.go              | **运行时 panic（竞态）**      | close 与内部 goroutine 的 send 没有同步顺序保证                   |
| 全 nil channel 的 `select` 无 `default`     | flowup/channel/select/v1/main.go      | **运行时死锁（fatal error）** | nil channel 在 select 中永不就绪，且无出口                        |

一句话总结： **编译错误是类型系统在静态检查阶段就能发现的问题（比如 channel 方向、类型 vs 值），而 panic / deadlock
是并发时序、channel 状态（nil / closed）在运行时才暴露的问题**——这类 bug 编译器帮不了忙，只能靠对 channel 语义的理解去规避。

## select 关键词的使用场景

`select` 的每个 `case` **只能是 channel 的发送或接收操作**，不能是任意布尔表达式——这是它和 `switch`/`if` 最本质的区别：

```go
select {
case v := <-c1:      // 接收
case c2 <- x:         // 发送
case v, ok := <-c3:   // 接收 + 判断是否 closed
default:              // 唯一的例外，不涉及 channel
}
```

不合法：`case x > 5:` 这种普通条件判断，编译器会直接报错。

常见使用场景，本质都是"同时监听/操作多个 channel"：

1. **多路等待**：多个 goroutine 往不同 channel 发数据，谁先就绪就处理谁；多个 case 同时就绪时随机选一个（避免饥饿）。
2. **超时控制**：配合 `time.After(d)` 实现经典的超时写法。
   ```go
   select {
   case res := <-resultChan:
       // ...
   case <-time.After(2 * time.Second):
       // timeout
   }
   ```
3. **非阻塞发送/接收**：配合 `default`，channel 没就绪时立刻走别的逻辑，不阻塞等待。
4. **优雅退出/取消**：监听 `ctx.Done()` 或自定义 `done` channel，随时响应外部的停止信号。
5. **心跳/定时任务**：配合 `time.NewTicker(d).C` 做周期性工作。
6. **扇入（fan-in）**：把多个 channel 的数据汇总到一个循环里处理，pipeline 模式常用组件。
7. **动态启用/禁用某个 case**：把某个 channel 变量置为 `nil`，对应 case
   就永远不会被选中，从而运行时"关闭"某条分支——但必须保证至少留一个可用出口（或有 `default`），否则会像本文件开头的例子一样死锁。
</content>
