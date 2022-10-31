//interface是只要你实现了接口定义的类型（方法），就可以作为接口的实例拿来用，在语言层面上不再需要其它的约束。

//在结构体中嵌入匿名类型成员，就可以继承匿名类型的方法。
//
//不仅可以嵌入匿名类型成员，还可以嵌入匿名的接口，或者匿名的指针对象，是一种纯虚继承，继承的只是接口的规范，而实现还得靠自己。

struct 成员小写导致rpc的结果全为默认值

二维切片中反复添加同一个切片，修改一个切片的值会影响二维切片中的所有值

安装go包：可以直接在程序中引用，然后运行使用go mod（GO111MODULE=on）

# linux

### 常用命令
1. ps
2. kill
3. ip
4. netstat
5. scp
6.

# Golang
## 常用命令
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
## 开发所用库
1. RPC框架 [kite](https://github.com/koding/kite )/[kitex](https://github.com/cloudwego/kitex ), gRPC
2. Web框架 [gin](https://gin-gonic.com/)
3. ORM框架 [gorm](https://gorm.io/)
4. Mock框架 [gomock](https://github.com/golang/mock) / [goconvey](https://github.com/smartystreets/goconvey)
5. 断言库 [assert](https://github.com/stretchr/testify)
6. flag包 [pflag](https://github.com/spf13/pflag)
7. validate [govalidator](http://github.com/asaskevich/govalidator)
8. websocket [websocket](https://github.com/gorilla/websocket)
9. redis key-value存储 [redis](https://redis.io/)
10. yaml [yaml.v3](https://gopkg.in/yaml.v3)
11. pretty [pretty](https://github.com/kr/pretty)

## 常用技术
1. 稳定性：缓存、降级、限流（熔断，[hystrix-go](https://gitee.com/mirrors/hystrix-go )，[uber rate limit](https://pkg.go.dev/go.uber.org/ratelimit )，[go x rate](https://pkg.go.dev/golang.org/x/time/rate )）
2. 池化（线程、协程、连接池）
3. 缓存cache（ [freecache](https://github.com/coocood/freecache )、 [groupcache](https://github.com/golang/groupcache )、[bigcache](https://github.com/allegro/bigcache) ，此外还有fastcache、offheap、ristretto等）
4. docker（[Docker镜像仓库](https://hub.docker.com/search?image_filter=official&q=) ）
5. kafka/rocketmq([官网](https://rocketmq.apache.org/ )，[控制面板](https://github.com/apache/rocketmq-dashboard )，使用参考go rocketmq examples)
6. k8s（全称Kubernetes，[Kubernetes](https://kubernetes.io/) ）
7. Elasticsearch
8. nginx

## 工程设计
1. 设计模式

## 中间件执行顺序
同一个中间件，他的前置逻辑越早执行，他的后置逻辑执行的越晚。

![输入图片说明](img/image1.png)

![输入图片说明](img/image.png)

# MySql
0. mariadb安装及设置
### 注意mysql和mariadb使用配置文件目录不同
### 主从备份需要手动创建数据库及数据表。
sudo apt-get install mariadb-server
切换root用户
使用mysql_secure_installation命令设置密码
设置新用户，并为其授权远程登录：
修改sudo vim /etc/mysql/mariadb.conf.d/50-server.cnf
重启服务  sudo systemctl restart mariadb.service
1. mysql 只能使用root登录
   alter user 'root'@'localhost' identified with mysql_native_password by '123456';
2. mysql 修改密码 
mysql -u root update mysql.user set authentication_string=PASSWORD('123456') where User='root'; flush privileges;
   set password for root@localhost = password('123456'); flush privileges;
3. mysql 创建用户并授权某个数据库:
```sql
CREATE USER 'test'@'localhost' IDENTIFIED BY '123456';
update user set host='%' where user='test';
grant all privileges on test.* to 'test'@'%'; --（两次）

```

4. MySql执行sql文件，a. source xxx.sql  b. mysql -u用户名 -p用户密码 < xxx.sql 
   
# 项目
1. [6.824](http://nil.csail.mit.edu/6.824/2020/schedule.html)
2. 数据库项目

## Docker
#### 1. 常用命令
```bash
docker images
docker ps -a
docker build -t getting-started . 
docker exec -it 0ec47bf44530 sh 
docker run -dp 3000:3000 getting-started  
```


### 1. 缓存穿透、击穿、雪崩
+ 穿透：指查询一个缓存和数据库都不存在的数据，导致尽管数据不存在但是每次都会到数据库查询，在访问量大的时候数据库可能挂掉。
+ 击穿：单个key值的缓存失效过期
+ 雪崩：redis缓存中大量的key同时失效，同时刚好有大量的请求，会直接访问数据库，造成数据库阻塞甚至宕机

### 2. 消息队列
+ 异步
+ 削峰
+ 解耦

## Golang

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
1. 初始化导入的包（包的初始化顺序并不是按照导入顺序（“从上到下”）执行的，runtime需要解析包依赖关系，没有依赖的包最先初始化，与变量初始化依赖关系类似）
2. 初始化包作用域的变量，runtime解析依赖关系，没有依赖的变量最先初始化，[golang变量初始化]()
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

//  mockgen -source=foo.go -package=mock -destination mock_foo.go


goland无限试用，ide-eval-reset，2021.2.2及之前版本。

# 计划
## 1. 后端
## 2. 嵌入式：树莓派->电机控制->四旋翼

#### [设计模式](./docs/design_pattern.md)