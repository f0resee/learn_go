# Golang
## Tools
+ delve
+ pprof
  + top
  + list func
  + web
## grpc
1. protoc install
   ```shell
   sudo apt install protobuf-compiler
   ```

2. generate proto 

   ```shell
   protoc --go_out=. --go_opt=paths=source_relative \
   --go-grpc_out=. --go-grpc_opt=paths=source_relative \
   helloworld/helloworld.proto
   ```
## 一、Golang基础
### 1. context
`context`包定义了`Context`类型，它包含了最后期限、取消信号以及其他其请求范围内跨API边界以及过程之间的值。

到达server的请求应该创建一个`Context`，向其他server发起的请求应该接收一个`Context`。在请求中的函数调用链应该传递`Context`，或者将它替换为一个用`WithCancel`、`WithDeadline`、`WithTimeout`或者是`WithValue`创建的`Context`。当一个`Context`被取消，从它派生出的所有`Context`都会被取消。

`WithCancel`、`WithDeadline`、`WithTimeout`这三个函数接收一个`Context`参数作为父`Context`，返回一个子`Context`以及一个`CancelFunc`。调用`CancelFunc`会取消子`Context`以及它的子`Context`，并移除父`Context`对子`Context`的引用，并停止所有相关的计时器。不调用`CancelFunc`会导致子`Context`以及它的子`Context`泄漏，直到父`Context`被取消或者是计时器超时。go vet工具会检查在每条控制流上`CancelFunc`都会被调用。

使用`Context`的程序应该遵守以下规则，以保证不同包之间接口一致，并使静态分析工具可以在检查`Context`的传播。
+ 不要将`Context`存储在一个结构体内，而是在需要它的函数之间进行显式地传递。`Context`应该是第一个参数，一般命名为`ctx`;
+ 不要传递一个空`Context`，即便函数允许这样做。如果不知道使用哪个`Context`，那么使用`context.TODO`;
+ 只在过程和API之间利用context的`Values`传递请求可见的数据，而不是将它用来传递函数中的可选参数;
+ 同一个`Context`可以被传递到不同协程中的函数，`Context`是并发安全的;

变量：

+ ```var Canceled = errors.New("context canceled")```

+ ```var DeadlineExceeded error = deadlineExceededError{} ```

函数：
+ ```func WithCancel(parent Context) (ctx Context, cancel CancelFunc)```
+ ```func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)```
+ ```WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)```

类型：

+ ```type CancelFunc func()```
+ type Context
```go
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
}
```
+ ```func Background() Context```
+ ```func TODO() Context```
+ ```func WithValue(parent Context, key, val any) Context```

## 二、常用命令
1. go help
2. go test
   ```shell
   go test xxx_test.go
   go test -v xxx_test.go # -v显示详细信息
   go test -run Testxxx xxx_test.go # -run运行某一个测试
   go test -bench=. -v xxx_test.go # -bench运行某一个benchmark
   ```
3. go设定proxy
   ```shell
   go env -w GOPROXY=https://goproxy.cn,direct
   ```
4

## 三、面试题目
### 基础语法
#### 1. = 和 := 的区别？
#### 2. 指针的作用
#### 3. Go 允许多个返回值吗？
#### 4. Go 有异常类型吗？
#### 5. 什么是协程（Goroutine）
#### 6. 如何高效地拼接字符串
#### 7. 什么是 rune 类型
#### 8. 如何判断 map 中是否包含某个 key ？
#### 09. Go 支持默认参数或可选参数吗？
#### 10. defer 的执行顺序
#### 11. 如何交换 2 个变量的值？
#### 12. Go 语言 tag 的用处？
#### 13. 如何判断 2 个字符串切片（slice) 是相等的？
#### 14. 字符串打印时，%v 和 %+v 的区别
#### 15. Go 语言中如何表示枚举值(enums)？
#### 16. 空 struct{} 的用途

### 实现原理

#### 1. init() 函数是什么时候执行的？
特点：
+ init函数先于main函数自动执行，不能被其他函数调用
+ init函数没有输入参数、返回值
+ 每个包可以有多个init函数
+ 包的每个源文件也可以有多个init函数
+ 同一个包的init执行顺序，golang没有明确定义，编程时要注意不要依赖这个执行顺序
+ 不同包的init函数按照包导入的依赖关系决定执行顺序

初始化执行顺序：

    a. 初始化导入的包（包的初始化顺序并不是按照导入顺序（“从上到下”）执行的，runtime需要解析包依赖关系，没有依赖的包最先初始化，与变量初始化依赖关系类似）
    b. 初始化包作用域的变量，runtime解析依赖关系，没有依赖的变量最先初始化，[golang变量初始化]()
#### 2. Go 语言的局部变量分配在栈上还是堆上？
#### 3. 2 个 interface 可以比较吗 ？
#### 4. 2 个 nil 可能不相等吗？
#### 5. 简述 Go 语言GC(垃圾回收)的工作原理
#### 6. 函数返回局部变量的指针是否安全？
#### 7. 非接口非接口的任意类型 T() 都能够调用 *T 的方法吗？反过来呢？

### 并发编程


#### 1. 无缓冲的 channel 和有缓冲的 channel 的区别？
#### 2. 什么是协程泄露(Goroutine Leak)？
#### 3. Go 可以限制运行时操作系统线程的数量吗？

### 代码输出

#### 1. 变量与常量
#### 2. 作用域
#### 3. defer 延迟调用

    熟悉 golang；
    熟悉 Shell编程，TCP, HTTP 协议；
    熟练使用 Linux 操作系统，熟悉 Linux 下多线程开发；
    有低延迟、Docker、Kubernetes，数据可视化等经验加分；
    对技术有追求，对量化行业有兴趣。

### Golang

#### 1.golang GC 回收介绍 (标记清除, 三色标记法, 混合写屏障)

#### 2. 什么时候会触发 golang GC 呢
#### 3. golang channel 你会用在什么地方 (一个是控制 goroutine数量 一个是主main 控制关闭 子 goroutine)
#### 4. golang 如何做超时控制 ? (time.After 和 context.WithTimeOut)
#### 5. golang select 一般使用在什么场景
#### 6. golang 如何 比较两个 map 使用相等
+ 不能使用```==```直接比较，但可以用来判断map是否为```nil```
+ 方式1：可以通过分别枚举map1和map2中的key-value来比较两个map是否相等
+ 方式2：使用```reflect.DeepEqual```进行比较

#### 7. golang sync.Cond

#### 8. interface
interface是只要你实现了接口定义的类型（方法），就可以作为接口的实例拿来用，在语言层面上不再需要其它的约束。

在结构体中嵌入匿名类型成员，就可以继承匿名类型的方法。

不仅可以嵌入匿名类型成员，还可以嵌入匿名的接口，或者匿名的指针对象，是一种纯虚继承，继承的只是接口的规范，而实现还得靠自己。

struct 成员小写导致rpc的结果全为默认值

二维切片中反复添加同一个切片，修改一个切片的值会影响二维切片中的所有值

安装go包：可以直接在程序中引用，然后运行使用go mod（GO111MODULE=on）