# version5：monitor goroutine 实现并发安全的 memo

## 什么是 monitor goroutine

Monitor goroutine（监控 goroutine）是一种并发模式：**规定某份数据只能由一个 goroutine 访问**，其他 goroutine 想读写它，都通过 channel 发消息给这个 goroutine，由它代劳。

这个"专职看守数据"的 goroutine 就是 monitor goroutine。它用「通信」代替「加锁」来保护共享状态。

对应 Go 谚语：

> Do not communicate by sharing memory; instead, share memory by communicating.
> 不要用共享内存来通信，而要用通信来共享内存。

## 本实现的结构

```go
type Memo struct{ requests chan request }   // 对外只暴露一个 channel

func New(f Func) *Memo {
    memo := &Memo{requests: make(chan request)}
    go memo.server(f)   // 启动 monitor goroutine
    return memo
}
```

`server` 就是 monitor goroutine：

```go
func (memo *Memo) server(f Func) {
    cache := make(map[string]*entry)   // 只属于 server 一个 goroutine
    for req := range memo.requests {
        e := cache[req.key]
        if e == nil {
            e = &entry{ready: make(chan struct{})}
            cache[req.key] = e          // 只有 server 读写 cache，天然无竞争
            go e.call(f, req.key)
        }
        go e.deliver(req.response)
    }
}
```

关键点：`cache` 这个 map **没有任何锁**，却是并发安全的。因为它是 `server` 的局部变量，只有 `server` 一个 goroutine 访问它。相比 version3/version4 用 `sync.Mutex` 保护 `cache`，这里完全不需要锁。

## 客户端如何交互

```go
func (memo *Memo) Get(key string) (any, error) {
    response := make(chan result)
    memo.requests <- request{key, response}  // 把请求(含回信 channel)发给 monitor
    res := <-response                         // 等 monitor 把结果送回来
    return res.value, res.err
}
```

- 任意多个 goroutine 并发调用 `Get`，只是往 `requests` channel 里塞消息。
- `server` 一次处理一个请求，串行操作 `cache`。
- 每个请求自带一个 `response` channel，用于接收专属回信。
- `Close()` 里 `close(memo.requests)`，`for range` 结束，monitor 退出。

## 为什么 call 和 deliver 要另开 goroutine

`server` 循环体里 `go e.call(...)` 和 `go e.deliver(...)` 都加了 `go`，是为了**不阻塞 monitor**：

- `e.call` 里执行真正耗时的 `f(key)`
- `e.deliver` 里有 `<-e.ready`，会阻塞等结果就绪

若 `server` 同步调用它们，整个 monitor 被卡住，其他 key 的请求也进不来。开 goroutine 后 `server` 立刻回到 `for` 接收下一个请求，实现「不同 key 的请求并行处理」。

而 `cache` 的读写始终留在 `server` 手里 —— 子 goroutine 只碰局部变量 `e`，不碰 `cache`，所以没有把共享状态泄露出去。

## duplicate suppression（重复抑制）

```go
type entry struct {
    res   result
    ready chan struct{} // res 就绪后 close
}
```

- 第一个请求某 key 时，建 `entry`、开 `call`，`ready` 尚未关闭。
- 同一 key 的后续请求发现 `e != nil`，不再重复调用 `f`，直接 `go e.deliver`。
- `call` 算完后 `close(e.ready)` 广播；所有等在 `<-e.ready` 的 `deliver` 一起被唤醒，各自把结果发回对应客户端。

对应包注释：

> Concurrent requests for the same key block until the first completes.

## 与其他版本对比

| 版本 | 保护 cache 的方式 | 特点 |
|------|------------------|------|
| version3 | `sync.Mutex` 全程持锁 | 简单，但同一时刻只有一个 `f` 在跑 |
| version4 | `sync.Mutex` + per-entry `ready` channel | 有重复抑制，锁只保护 map |
| version5 | monitor goroutine + channel | 无锁；用通信共享内存 |
