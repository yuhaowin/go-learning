### 接口值的内部结构

接口变量在运行时不是直接存数据，而是存一个二元组：

```
(动态类型 type, 动态值 value)
```

- **动态类型**：接口当前实际装的具体类型，例如 `*scheduler.SimpleScheduler`。
- **动态值**：该具体类型对应的实际数据（如果是指针，就是指针本身）。

例如：

```shell
var s engine.Scheduler = &scheduler.SimpleScheduler{}
// s 内部 = (type: *scheduler.SimpleScheduler, value: 0xc0000.../指向该实例的指针)
```

如果换成另一个满足接口的具体类型赋值，`(type, value)` 会整体替换：

```shell
s = FakeScheduler{}
// s 内部 = (type: FakeScheduler, value: FakeScheduler{} 的拷贝)
```

#### 由此引出的两个坑/要点

1. **接口 nil 判断**：只有当 `(type, value)` 都是零值 `(nil, nil)` 时，接口才 `== nil`。 如果 type 已经被设成某个具体类型（哪怕
   value 是 nil 指针），接口本身就不等于 nil。
2. **方法调用即多态**：调用接口方法时，运行时查 `type` 对应的方法表，找到具体实现执行——这是接口实现多态的机制来源。

---

### 为什么"值类型"字段能接收指针（`ConcurrentEngine.Scheduler`）

`ConcurrentEngine.Scheduler` 字段的类型是 **接口** `Scheduler`（engine/concurrent.go），不是某个具体 struct。
"值类型 / 指针类型" 的拷贝语义只针对具体类型；接口类型本身就是上面说的 `(type, value)` 二元组，
能装入任何实现了该方法集的具体类型，不管这个具体类型是值类型还是指针类型。

所以：

```shell
type ConcurrentEngine struct {
    Scheduler Scheduler // 接口类型字段，不是具体 struct
}
```

赋值 `Scheduler: &scheduler.SimpleScheduler{}` 完全合法：接口内部存的动态类型就是 `*SimpleScheduler`。

如果字段类型换成具体的 `scheduler.SimpleScheduler`（而不是接口），才会出现真正的"值类型字段不能接收指针"的编译错误。

### 为什么 SimpleScheduler 的方法必须用指针接收者

`SimpleScheduler.ConfigureMasterWorkerChan` 会给 `s.workerChan` 赋值（scheduler/simple.go）。
如果用值接收者，方法内修改的只是调用者传入值的一份拷贝，外部（包括 `Submit` 里用到的 `s.workerChan`） 看不到这次赋值，
`workerChan` 永远是 nil。

另外，Go 的方法集规则：只要有一个方法用了指针接收者，就只有 `*SimpleScheduler` 能满足接口，
`SimpleScheduler`（值类型）不满足。所以 main.go 里必须写 `&scheduler.SimpleScheduler{}` 传指针， 否则编译报错「
`SimpleScheduler` does not implement `Scheduler`」。
