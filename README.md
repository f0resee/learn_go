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
1. RPC框架 [kite](https://github.com/koding/kite)/[kitex](https://github.com/cloudwego/kitex)
2. Web框架 [gin](https://gin-gonic.com/)
3. ORM框架 [gorm](https://gorm.io/)
4. Mock框架 [gomock](https://github.com/golang/mock) / [goconvey](https://github.com/smartystreets/goconvey)
5. 断言库 [assert](https://github.com/stretchr/testify)
6. flag包 [pflag](https://github.com/spf13/pflag)
7. validate [govalidator](http://github.com/asaskevich/govalidator)
8. websocket [websocket](https://github.com/gorilla/websocket)
9. redis key-value存储 [redis](https://redis.io/)

## 常用技术
1. 稳定性：缓存、降级、限流（熔断，[hystrix-go](https://gitee.com/mirrors/hystrix-go) ，[uber rate limit](https://pkg.go.dev/go.uber.org/ratelimit) ，[go x rate](https://pkg.go.dev/golang.org/x/time/rate) ）
2. docker（[Docker镜像仓库](https://hub.docker.com/search?image_filter=official&q=) ）
3. kafka/rocketmq([官网](https://rocketmq.apache.org/ )，[控制面板](https://github.com/apache/rocketmq-dashboard )，使用参考go rocketmq examples)
4. k8s（全称Kubernetes，[Kubernetes](https://kubernetes.io/) ）
5. Elasticsearch
6. nginx

## 中间件执行顺序
同一个中间件，他的前置逻辑越早执行，他的后置逻辑执行的越晚。
![输入图片说明](img/image1.png)
![输入图片说明](img/image.png)

## 

# MySql
1. mysql 只能root登录
   alter user 'root'@'localhost' identified with mysql_native_password by '123456';
2. mysql 修改密码
3. mysql 创建用户
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

## Redis
#### 1. 缓存穿透、击穿、雪崩
+ 穿透：指查询一个缓存和数据库都不存在的数据，导致尽管数据不存在但是每次都会到数据库查询，在访问量大的时候数据库可能挂掉。
+ 击穿：单个key值的缓存失效过期
+ 雪崩：redis缓存中大量的key同时失效，同时刚好有大量的请求，会直接访问数据库，造成数据库阻塞甚至宕机

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

    熟悉 golang；
    熟悉 Shell编程，TCP, HTTP 协议；
    熟练使用 Linux 操作系统，熟悉 Linux 下多线程开发；
    有低延迟、Docker、Kubernetes，数据可视化等经验加分；
    对技术有追求，对量化行业有兴趣。

goland无限试用，ide-eval-reset，2021.2.2及之前版本。