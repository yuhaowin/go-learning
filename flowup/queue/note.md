# Test 与 Example 的区别

## Test 函数

- 写法：`func TestXxx(t *testing.T)`
- 接收 `*testing.T`，通过 `t.Error` / `t.Fatal` 等方法显式断言成功或失败
- 用于常规单元测试：边界条件、错误路径等，断言逻辑需要自己写
- 不会出现在 GoDoc 文档里

## Example 函数

- 写法：`func ExampleXxx()`，不接收参数
- 没有 `t.Error` 这类断言，靠函数末尾的 `// Output:` 注释生效：
  go test 会捕获 `fmt.Println` / `fmt.Print` 等标准输出，与注释里的期望内容逐字符比对，不一致则失败
- 如果没有 `// Output:` 注释，函数只会被编译检查（保证代码能跑通），但**不会被执行**，也不算测试用例
- 命名关联文档：`ExampleQueue_Pop` 表示这是 `Queue` 类型 `Pop` 方法的示例，会被 `go doc` / pkg.go.dev 关联展示到该方法文档下

## 一句话总结

- `Test`：给自己验证逻辑用的常规测试
- `Example`：既是测试（校验输出），又是"可执行的文档"，会被提取展示为 API 使用示例

## 命令行运行方式

`Test` 和 `Example`（带 `// Output:` 注释的）都是通过 `go test` 统一执行的，不需要额外命令：

- `go test`：默认运行所有 `TestXxx`，以及有 `// Output:` 注释的 `ExampleXxx`
- `go test -run 正则`：按名字过滤，只跑匹配的用例，`Test`/`Example` 共用同一个 `-run` 参数
- `go test -v`：显示每个用例的运行结果（`=== RUN` / `--- PASS`）；不加 `-v` 只在失败时才输出细节

还有第三类 `BenchmarkXxx`，同样写在 `_test.go` 里，但默认**不会**被 `go test` 执行，需要显式加 `-bench` 参数（如 `go test -bench=.`）才会跑。

## 示例（flowup/queue/queue_test.go）

```go
func ExampleQueue_Pop() {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	// Output:
	// 1
	// 2
	// false
	// 3
	// true
}
```

运行 `go test ./flowup/queue/... -v`，输出：

```
=== RUN   ExampleQueue_Pop
--- PASS: ExampleQueue_Pop (0.00s)
PASS
```
### 生成 doc 网页

```shell
godoc -http :6060
```