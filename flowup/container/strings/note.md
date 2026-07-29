## Go 没有 char 类型

Go 用两种类型表示字符：

- `byte`：`uint8` 的别名，表示一个 ASCII 字符（1 字节）
- `rune`：`int32` 的别名，表示一个 Unicode 码点（4 字节），用于处理多字节字符（比如中文）

```go
var b byte = 'A'      // 65
var r rune = '中'      // Unicode 码点
```

因为 Go 源码默认是 UTF-8 编码，字符串按字节遍历会出问题，遇到中文这类多字节字符需要用 `rune` 或 `for range` 遍历字符串才能正确处理。
