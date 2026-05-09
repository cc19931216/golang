# Go语言基础示例项目 Code Wiki

## 1 项目概述

### 1.1 项目简介

本项目（`github.com/cncamp/golang`）是一个Go语言基础示例代码仓库，包含Go语言核心语法和特性的演示示例。项目中包含httpserver服务示例以及多个功能性示例模块，涵盖了Go语言从基础到进阶的核心知识点。

### 1.2 项目信息

| 属性 | 值 |
|------|-----|
| 模块名称 | `github.com/cncamp/golang` |
| Go版本 | 1.16+ |
| 项目类型 | 示例代码仓库 |
| 编程语言 | Go |
| 许可证 | 未指定 |

### 1.3 目录结构

```
/workspace/
├── .idea/                          # GoLand IDE 配置目录
├── examples/                       # 示例代码目录
│   ├── busyloop/                    # Busy Loop 示例
│   ├── callbacks/                   # 回调函数示例
│   ├── context/                     # Context 使用示例
│   │   ├── context/                # Context.Context 用法
│   │   └── donechannel/            # Done Channel 用法
│   ├── cpuprofiling/               # CPU 性能分析示例
│   ├── defer/                       # defer 语句示例
│   ├── forloop/                    # for 循环示例
│   ├── govet/                      # go vet 工具示例
│   ├── helloworld/                 # Hello World 示例
│   ├── init/                        # init 函数执行顺序示例
│   ├── interface/                  # 接口示例
│   ├── malloc/                     # 内存分配示例
│   ├── map/                        # Map 数据结构示例
│   ├── once/                       # sync.Once 示例
│   ├── pointer/                    # 指针示例
│   ├── reflect/                    # 反射示例
│   ├── slice/                      # Slice 示例
│   │   ├── forrange/               # for range 循环
│   │   └── makenew/                 # make 与 new 对比
│   ├── struct/                     # 结构体示例
│   ├── structs/                    # 结构体嵌套示例
│   ├── task/                       # 云原生任务示例
│   │   ├── CloudNative1.1.go
│   │   └── CloudNative1.2.go
│   └── waitgroup/                  # WaitGroup 示例
├── httpserver/                     # HTTP 服务器示例
│   ├── bin/amd64/                  # 编译输出目录
│   ├── Dockerfile
│   ├── Makefile
│   ├── build.sh
│   ├── httpserver
│   └── main.go
├── go.mod                          # Go 模块依赖文件
├── go.sum                          # Go 模块校验文件
└── readme.MD                       # 项目说明文件
```

---

## 2 项目依赖

### 2.1 Go Modules 配置

项目使用Go Modules进行依赖管理，`go.mod`文件内容如下：

```go
module github.com/cncamp/golang

go 1.16

require (
    github.com/golang/glog v0.0.0-20210429001901-424d2337a529
    github.com/stretchr/testify v1.7.0
)
```

### 2.2 依赖说明

| 依赖包 | 版本 | 用途 |
|--------|------|------|
| `github.com/golang/glog` | v0.0.0-20210429001901-424d2337a529 | Google开发的日志库，提供分级日志输出功能 |
| `github.com/stretchr/testify` | v1.7.0 | 常用的Go测试工具库，提供断言和mock功能 |

### 2.3 依赖管理命令

```bash
# 下载并整理依赖
go mod download
go mod tidy

# 查看依赖
go list -m all

# 清理未使用的依赖
go mod tidy
```

---

## 3 主要模块说明

### 3.1 Hello World 模块 (`examples/helloworld/`)

**功能说明**：演示Go语言命令行参数解析的基础用法

**核心代码**：[main.go](file:///workspace/examples/helloworld/main.go)

**关键函数**：

| 函数名 | 功能 | 参数 | 返回值 |
|--------|------|------|--------|
| `flag.String()` | 定义命令行字符串参数 | name(参数名)、default(默认值)、usage(说明) | *string指针 |
| `flag.Parse()` | 解析命令行参数 | 无 | 无 |
| `fmt.Sprintf()` | 格式化字符串 | format(格式)、args(参数) | string |

**示例用法**：

```bash
go run main.go -name=Alice
# 输出: Hello Alice from Go
```

---

### 3.2 指针模块 (`examples/pointer/`)

**功能说明**：演示Go语言指针的声明、使用及值传递与指针传递的区别

**核心代码**：[main.go](file:///workspace/examples/pointer/main.go)

**关键概念**：

- `&` 操作符：获取变量地址
- `*` 操作符：解引用（获取指针指向的值）
- 值传递：函数参数为结构体副本，修改不影响原值
- 指针传递：函数参数为结构体指针，修改会影响原值

**关键结构体**：

```go
type ParameterStruct struct {
    Name string
}
```

**关键函数**：

| 函数名 | 功能 | 说明 |
|--------|------|------|
| `changeParameter(para *ParameterStruct, value string)` | 修改指针指向的结构体 | 传指针，可修改原值 |
| `cannotChangeParameter(para ParameterStruct, value string)` | 尝试修改结构体 | 传值，修改不影响原值 |

---

### 3.3 循环模块 (`examples/forloop/`)

**功能说明**：演示Go语言for循环的基本用法和字符串遍历

**核心代码**：[main.go](file:///workspace/examples/forloop/main.go)

**循环类型**：

| 循环类型 | 示例 | 说明 |
|----------|------|------|
| 标准for循环 | `for i := 0; i < 3; i++` | 类似C语言语法 |
| for-range循环 | `for i, c := range fullString` | 遍历字符串、数组、切片、map |

---

### 3.4 Slice 切片模块 (`examples/slice/`)

**功能说明**：演示Go语言Slice（切片）的创建、切片操作和修改

**核心代码**：

- [main.go](file:///workspace/examples/slice/main.go) - 基础切片操作
- [forrange/main.go](file:///workspace/examples/slice/forrange/main.go) - for range遍历
- [makenew/main.go](file:///workspace/examples/slice/makenew/main.go) - make与new对比

**Slice创建方式对比**：

| 创建方式 | 代码 | 说明 |
|----------|------|------|
| 切片表达式 | `myArray[1:3]` | 从数组或切片中截取 |
| make函数 | `make([]int, 0)` | 创建空切片 |
| make带长度 | `make([]int, 10)` | 创建长度为10的切片 |
| make带容量 | `make([]int, 10, 20)` | 创建长度10、容量20的切片 |
| new函数 | `new([]int)` | 返回*[]int指针，值为nil |

**关键函数**：

```go
func deleteItem(slice []int, index int) []int {
    return append(slice[:index], slice[index+1:]...)
}
```

该函数演示了如何删除切片指定索引的元素。

---

### 3.5 Map 映射模块 (`examples/map/`)

**功能说明**：演示Go语言Map数据结构的创建和使用

**核心代码**：[main.go](file:///workspace/examples/map/main.go)

**创建方式**：

```go
myMap := make(map[string]string, 10)  // 指定初始容量
myFuncMap := map[string]func() int{   // 值为函数的Map
    "funcA": func() int { return 1 },
}
```

**操作方法**：

| 操作 | 代码 | 说明 |
|------|------|------|
| 赋值 | `myMap["a"] = "b"` | 添加或更新键值对 |
| 读取 | `value, exists := myMap["a"]` | 返回值和是否存在 |
| 遍历 | `for k, v := range myMap` | 遍历所有键值对 |

---

### 3.6 Struct 结构体模块 (`examples/struct/`)

**功能说明**：演示Go语言结构体的定义和使用

**核心代码**：[main.go](file:///workspace/examples/struct/main.go)

**结构体定义**：

```go
type MyType struct {
    Name string `json:"name"`  // 带JSON标签的字段
}
```

**结构体嵌套（嵌入）**：

```go
type TypeA struct {
    P1 string
}

type TypeB struct {
    P2 string
    TypeA  // 匿名嵌入，可直接访问 tb.P1
}
```

---

### 3.7 Interface 接口模块 (`examples/interface/`)

**功能说明**：演示Go语言接口的定义和实现

**核心代码**：[main.go](file:///workspace/examples/interface/main.go)

**接口定义**：

```go
type IF interface {
    getName() string
}
```

**实现类型**：

| 类型 | 说明 |
|------|------|
| `Human` | 结构体，通过指针receiver实现接口 |
| `Plane` | 结构体，通过值receiver实现接口 |
| `Car` | 结构体，通过指针receiver实现接口 |

**方法定义示例**：

```go
func (h *Human) getName() string {
    return h.firstName + "," + h.lastName
}

func (p Plane) getName() string {
    return fmt.Sprintf("vendor: %s, model: %s", p.vendor, p.model)
}
```

---

### 3.8 Defer 延迟执行模块 (`examples/defer/`)

**功能说明**：演示Go语言defer语句的执行机制

**核心代码**：[main.go](file:///workspace/examples/defer/main.go)

**核心概念**：

- defer语句在函数返回前执行
- 多个defer按LIFO（后进先出）顺序执行
- defer常用于资源释放和异常处理

**示例代码**：

```go
func main() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
    // 输出顺序: 3, 2, 1
}
```

---

### 3.9 Context 上下文模块 (`examples/context/`)

#### 3.9.1 Context.Context 用法

**核心代码**：[context/main.go](file:///workspace/examples/context/context/main.go)

**关键功能**：

| 函数 | 功能 |
|------|------|
| `context.Background()` | 创建根Context |
| `context.WithValue()` | 创建带值的Context |
| `context.WithTimeout()` | 创建带超时的Context |
| `ctx.Done()` | 返回完成信号channel |
| `ctx.Value()` | 获取Context中的值 |

**示例**：

```go
baseCtx := context.Background()
ctx := context.WithValue(baseCtx, "a", "b")
timeoutCtx, cancel := context.WithTimeout(baseCtx, time.Second)
```

#### 3.9.2 Done Channel 用法

**核心代码**：[donechannel/main.go](file:///workspace/examples/context/donechannel/main.go)

**模式说明**：使用无缓冲channel实现父子协程间的通信和取消信号传递

**核心组件**：

| 组件 | 类型 | 用途 |
|------|------|------|
| `messages` | `chan int` | 带缓冲的消息通道 |
| `done` | `chan bool` | 完成信号通道 |

---

### 3.10 HTTP 服务器模块 (`httpserver/`)

**功能说明**：完整的HTTP服务器示例，包含路由处理和性能分析

**核心代码**：[main.go](file:///workspace/httpserver/main.go)

**关键函数**：

| 函数名 | 功能 | 参数 |
|--------|------|------|
| `main()` | 主函数，启动HTTP服务器 | 无 |
| `rootHandler(w http.ResponseWriter, r *http.Request)` | 根路径处理器 | 响应写入器、请求对象 |
| `healthz(w http.ResponseWriter, r *http.Request)` | 健康检查处理器 | 响应写入器、请求对象 |

**服务器配置**：

```go
http.HandleFunc("/", rootHandler)
http.HandleFunc("/healthz", healthz)
err := http.ListenAndServe(":80", nil)
```

**pprof性能分析**（已注释）：

```go
import _ "net/http/pprof"  // 启用pprof

// 路由配置（注释状态）
mux.HandleFunc("/debug/pprof/", pprof.Index)
mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
```

---

## 4 其他示例模块

### 4.1 Callbacks 回调模块 (`examples/callbacks/`)

**功能说明**：演示Go语言函数作为参数传递的回调模式

**核心代码**：[main.go](file:///workspace/examples/callbacks/main.go)

### 4.2 Once 单次执行模块 (`examples/once/`)

**功能说明**：演示`sync.Once`确保代码只执行一次

### 4.3 WaitGroup 等待组模块 (`examples/waitgroup/`)

**功能说明**：演示`sync.WaitGroup`用于协程同步

### 4.4 Malloc 内存分配模块 (`examples/malloc/`)

**功能说明**：演示Go语言的内存分配机制

### 4.5 Reflect 反射模块 (`examples/reflect/`)

**功能说明**：演示Go语言反射API的使用

**关键用法**：

```go
mt := MyType{Name: "test"}
myType := reflect.TypeOf(mt)
name := myType.Field(0)
tag := name.Tag.Get("json")  // 获取json标签
```

### 4.6 CPU Profiling 模块 (`examples/cpuprofiling/`)

**功能说明**：演示Go程序CPU性能分析的方法

### 4.7 Busy Loop 模块 (`examples/busyloop/`)

**功能说明**：演示忙等待循环的实现

### 4.8 Go Vet 模块 (`examples/govet/`)

**功能说明**：演示`go vet`静态分析工具的使用

### 4.9 Init 执行顺序模块 (`examples/init/`)

**功能说明**：演示Go程序init函数的执行顺序

---

## 5 项目构建与运行

### 5.1 环境要求

- Go SDK 1.16 或更高版本
- 支持Linux/Unix/macOS/Windows系统

### 5.2 快速开始

```bash
# 克隆项目后，进入项目目录
cd /workspace

# 下载依赖
go mod download

# 运行Hello World示例
cd examples/helloworld
go run main.go -name=Alice

# 运行HTTP服务器
cd httpserver
go build -o httpserver main.go
sudo ./httpserver  # 需要80端口权限
```

### 5.3 运行测试

```bash
# 运行所有测试
go test ./...

# 运行指定包测试
go test ./examples/callbacks/...
```

### 5.4 静态分析

```bash
# 使用go vet检查代码
go vet ./...

# 使用gofmt格式化代码
gofmt -w .
```

### 5.5 HTTP服务器Docker部署

```bash
cd httpserver

# 构建Docker镜像
docker build -t cncamp/httpserver:latest .

# 运行容器
docker run -p 80:80 cncamp/httpserver:latest
```

---

## 6 依赖关系图

```
github.com/cncamp/golang
├── examples/
│   ├── helloworld/      → 依赖: fmt, flag, os
│   ├── pointer/        → 依赖: fmt
│   ├── forloop/        → 依赖: fmt
│   ├── slice/          → 依赖: fmt
│   ├── map/            → 依赖: fmt
│   ├── struct/         → 依赖: reflect
│   ├── interface/      → 依赖: fmt
│   ├── defer/          → 依赖: fmt, sync, time
│   ├── context/        → 依赖: context, fmt, time
│   ├── callbacks/      → 依赖: (测试相关)
│   ├── waitgroup/      → 依赖: sync, time
│   ├── reflect/        → 依赖: reflect
│   └── once/           → 依赖: sync
│
└── httpserver/         → 依赖: flag, fmt, io, log, net/http, glog
```

---

## 7 关键概念速查

### 7.1 Go语言特点

| 特点 | 说明 |
|------|------|
| 静态类型 | 编译时检查类型安全 |
| 并发支持 | 原生支持goroutine和channel |
| 垃圾回收 | 自动内存管理 |
| 接口隐式实现 | 无需显式声明实现接口 |
| defer语句 | 函数退出时延迟执行 |

### 7.2 常用标准库

| 包名 | 用途 |
|------|------|
| `fmt` | 格式化I/O操作 |
| `os` | 操作系统功能 |
| `flag` | 命令行参数解析 |
| `net/http` | HTTP客户端和服务端 |
| `context` | 上下文控制和取消信号 |
| `sync` | 同步原语（Mutex, WaitGroup, Once） |
| `reflect` | 运行时反射 |
| `time` | 时间操作 |

### 7.3 内存分配对比

| 函数 | 返回类型 | 说明 |
|------|----------|------|
| `new(T)` | `*T` | 分配零值内存，返回指针 |
| `make(T, ...)` | `T` | 分配并初始化切片、map、channel |

---

## 8 附录

### 8.1 相关资源

- [Go语言官方文档](https://go.dev/doc/)
- [Go语言标准库](https://pkg.go.dev/std)
- [Effective Go中文版](https://go-zh.org/doc/effective_go.html)

### 8.2 文档版本

| 版本 | 日期 | 说明 |
|------|------|------|
| 1.0 | 2024年 | 初始文档 |

---

*本文档由Code Wiki自动生成*
