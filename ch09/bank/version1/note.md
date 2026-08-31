这是 Go 并发里经典的 monitor goroutine（管家协程） 模式，用一个专属 goroutine 独占 balance 变量，外部完全靠 channel 通信来读写它，从而避免用锁。

逐行看这个 select：

- deposits 和 balances 都是无缓冲 channel（第 4、5 行）
- case amount := <-deposits: —— 如果有人调用 Deposit(amount)（第 7 行 deposits <- amount），这个分支就绪，接收到金额后执行 balance += amount
- case balances <- balance:（第 16 行）—— 如果有人调用 Balance()（第 8 行 <-balances），这个分支就绪，把当前 balance 的值发送出去

第 16 行的分支体是空的（{}），因为"发送"这个动作本身（balances <- balance）就是它要做的全部事情——把当前 balance 值交给等待读取的调用方，不需要额外操作。

select 会在这两个 case 之间随时等待哪个先就绪（谁先有人来通信就执行谁），执行完一次立刻回到 for 循环继续等待下一次操作。因为只有 teller 这一个 goroutine 会读写 balance 变量（第 11 行注释 "balance is confined to teller
goroutine"），所以不会有数据竞争——即使有很多个 goroutine 同时调用 Deposit/Balance，它们都要通过 channel 排队跟 teller "对话"，天然是串行、安全的。

init()（第 21-23 行）在包加载时就启动这个 teller goroutine，一直在后台跑着处理请求。                                                                                                          
                                                                                                       