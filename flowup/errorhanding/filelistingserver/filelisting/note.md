# userError 错误处理笔记

## 1. web.go 中的 userError 接口

```go
type userError interface {
	error
	Message() string
}
```

- `error` 是接口嵌入，要求实现 `Error() string`（内置 `error` 接口）。
- 额外要求 `Message() string`，作为可以直接展示给终端用户的提示文案。
- 用途：区分"用户错误"（如参数校验失败，可把消息原样返回给客户端）和"系统错误"（如文件不存在、权限不足，只记日志、不暴露细节）。

`errWrapper` 里的用法（web.go:33-36）：

```go
if userErr, ok := err.(userError); ok {
	http.Error(writer, userErr.Message(), http.StatusBadRequest)
	return
}
```

- 若断言成功，说明这是"用户错误"，返回 400 + `Message()` 内容。
- 否则走系统错误分支，按 `os.IsNotExist` / `os.IsPermission` 等映射成 404 / 403 / 500，不泄露内部错误信息。

⚠️ 已知问题：`err.(userError)` 是直接类型断言，若 `err` 被 `fmt.Errorf("...: %w", err)` 包装过会断言失败。更稳妥的写法是 `errors.As(err, &userErr)`。

## 2. filelisting/handler.go 中的 userError 类型

```go
type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}
```

- 基于 `string` 定义了一个具体类型（和 web.go 里的 `userError` 接口同名但不是一回事，只是恰好同名）。
- 同时实现了 `Error()` 和 `Message()`，因此天然满足 web.go 里 `userError` 接口的要求。
- 用法（handler.go:27）：

```go
return userError(fmt.Sprintf("path %s must start with %s", request.URL.Path, prefix))
```

直接把格式化字符串包装成 `userError` 返回，回到 `errWrapper` 里就会被识别为用户错误。

## 3. 为什么用 string 而不是 struct

用 `struct` 也能实现同样效果：

```go
type userError struct {
	msg string
}

func (e userError) Error() string   { return e.msg }
func (e userError) Message() string { return e.msg }
```

- `string` 版本：错误信息只有一个字段，直接把字符串当底层类型最简洁，构造时 `userError("xxx")`，省去字段定义和取值样板代码。
- `struct` 版本：如果以后要携带更多信息（错误码、HTTP 状态码、原始 error 等），struct 更容易扩展。
- 当前场景只需要一条文本消息，用 `string` 是刚好够用、不过度设计的写法。
