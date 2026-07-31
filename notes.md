# 学习笔记（工具链 / 环境）

## GOROOT

Go **自身的安装目录**——标准库源码、编译器、工具链都在这里。装好 Go 后由 `go` 命令自动推断，基本不需要手动设置。

```bash
go env GOROOT
# /opt/homebrew/opt/go/libexec
```

## GOPATH

Go Modules 出现前（~1.10 之前）的 **工作区目录**，所有第三方包和自己的项目都必须放在 `$GOPATH/src/<import-path>/`
下，编译产物和依赖分别放在 `$GOPATH/pkg`、`$GOPATH/bin`。

```bash
go env GOPATH
# /Users/sqb/go
```

进入 Modules 时代后 GOPATH 没有消失，退化为：

- `go install` 生成的可执行文件默认放到 `$GOPATH/bin`
- 下载的依赖模块缓存放在 `$GOPATH/pkg/mod`

一句话： **GOROOT 是 Go 装在哪，GOPATH 是（旧模式下）代码/依赖放在哪**。

### Global GOPATH vs Project GOPATH（GoLand/IntelliJ 概念）

这是 GoLand 等 IDE 的概念，不是 `go` 命令本身的术语——`go env` 只有一个全局 `GOPATH`。GoLand 在 GOPATH 模式（Modules 出现前）下把
GOPATH 配置拆成两层：

- **Global GOPATH**：即 `go env GOPATH` 返回的路径（本机是 `/Users/sqb/go`），对所有 Go 项目生效，是默认共享工作区。对应
  GoLand `Preferences → Languages & Frameworks → Go → GOPATH` 里的 "Global GOPATH"。
- **Project GOPATH**：GoLand 允许给单个项目额外指定一个只在该项目生效的 GOPATH 目录，配置存在项目的 `.idea/`
  里，不影响其他项目。用途是在没有 Modules 时给项目做依赖隔离（不同项目可能需要同一个包的不同版本）。解析时 Project GOPATH
  优先于 Global GOPATH，效果类似把两者用 `:` 拼进 `GOPATH` 环境变量。

本仓库 `.idea/*.xml` 里没有任何 GOPATH 相关配置，因为项目用的是 Go Modules（有 `go.mod`），GoLand 检测到 Modules
模式后会切换依赖解析方式，Global/Project GOPATH 完全不参与当前项目的 import 解析，只在未启用 Modules 的老项目里才用得上。

## Go Modules

Go 官方的 **依赖管理机制**（1.11 引入，1.16 起默认开启），解决两个问题：项目叫什么名字（import 路径前缀）、依赖哪些包及版本。

核心文件 `go.mod`（本仓库根目录）：

```
module golearning

go 1.26
```

- `module golearning`：声明模块路径，仓库内所有包的 import 路径都以它为前缀，例如 `flowup/queue` 包的完整 import 路径是
  `golearning/flowup/queue`。
- `go 1.26`：声明使用的 Go 语言版本。

引入第三方依赖后会多出 `require github.com/xxx/yyy v1.2.3` 这样的行；同时会生成 `go.sum` 记录依赖内容的哈希值用于校验完整性（类似
`package-lock.json`）。本仓库暂无第三方依赖，所以没有 `go.sum`。

**相比 GOPATH 解决的问题**：GOPATH 时代依赖全局共享、没有版本概念，升级一个项目的依赖可能悄悄影响另一个项目。Modules
让每个项目可以放在任意目录（本仓库就不在 GOPATH 下），自行声明依赖及精确版本，项目间互不干扰；依赖仍缓存在 `$GOPATH/pkg/mod`
复用下载，但各项目按自己的 `go.mod` 锁定版本。

常用命令：

| 命令                        | 作用                                                                     |
|-----------------------------|--------------------------------------------------------------------------|
| `go mod init <module-path>` | 初始化，生成 `go.mod`                                                    |
| `go get <pkg>@<version>`    | 添加/升级依赖，写入 `go.mod`/`go.sum`                                    |
| `go mod tidy`               | 按代码实际 import 情况补全缺失依赖、清掉未用到的                         |
| `go mod why <pkg>`          | 查看某依赖为什么被引入                                                   |
| `go mod download`           | 把 `go.mod` 里的依赖下载到本地模块缓存，不改 `go.mod`/`go.sum`，也不编译 |

`go mod download` 的常见用法：

- `go mod download` — 下载当前模块所有依赖（及其间接依赖）到本地缓存。
- `go mod download all` — 下载构建/测试所需的全部模块，包括测试依赖等。
- `go mod download -x` — 打印详细执行过程（调试用）。
- `go mod download <module>` — 只下载指定的某个模块。
