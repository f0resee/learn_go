## 开发所用库
1. RPC框架 [kite](https://github.com/koding/kite )/[kitex](https://github.com/cloudwego/kitex ), [gRPC](https://grpc.io/docs/)
2. Web框架 [gin](https://gin-gonic.com/)
3. ORM框架 [gorm](https://gorm.io/)
4. Mock框架 [gomock](https://github.com/golang/mock) /[testify/mock](https://github.com/stretchr/testify )/ [goconvey](https://github.com/smartystreets/goconvey )。mock本质上是对测试过程中所依赖的一些函数/接口进行接管， 即便真实调用尚未实现/不可直接调用也可以返回预期值。
   如何编写可mock的代码
   1. mock 作用的是接口，因此将依赖抽象为接口，而不是直接依赖具体的类。
   2. 不直接依赖的实例，而是使用依赖注入降低耦合性。
6. 断言库 [assert](https://github.com/stretchr/testify)
7. flag包 [pflag](https://github.com/spf13/pflag)
8. validate [govalidator](http://github.com/asaskevich/govalidator)
9. websocket [websocket](https://github.com/gorilla/websocket)
10. redis key-value存储 [redis](https://redis.io/)
11. yaml [yaml.v3](https://gopkg.in/yaml.v3)
12. pretty [pretty](https://github.com/kr/pretty)

## 常用技术
1. 稳定性：缓存、降级、限流（熔断，[hystrix-go](https://gitee.com/mirrors/hystrix-go )，[uber rate limit](https://pkg.go.dev/go.uber.org/ratelimit )，[go x rate](https://pkg.go.dev/golang.org/x/time/rate )）
2. 池化（线程、协程、连接池）
3. 缓存cache（ [freecache](https://github.com/coocood/freecache )、 [groupcache](https://github.com/golang/groupcache )、[bigcache](https://github.com/allegro/bigcache) ，此外还有fastcache、offheap、ristretto等）
4. docker（[Docker镜像仓库](https://hub.docker.com/search?image_filter=official&q=) ,[Docker文档](https://docs.docker.com/get-started/overview/) ,[go docker库](https://github.com/moby/moby )）
5. kafka/rocketmq([官网](https://rocketmq.apache.org/ )，[控制面板](https://github.com/apache/rocketmq-dashboard )，使用参考go rocketmq examples)
6. k8s（全称Kubernetes，[Kubernetes](https://kubernetes.io/) ）
7. Elasticsearch
8. nginx
9. 监控[prometheus](https://prometheus.io/ )以及[grafana](https://grafana.com/ )

## 设计
1. [设计模式](./docs/design_pattern.md )

## 中间件执行顺序
同一个中间件，他的前置逻辑越早执行，他的后置逻辑执行的越晚。

![输入图片说明](img/image1.png)

![输入图片说明](img/image.png)

## 项目
1. [6.824](http://nil.csail.mit.edu/6.824/2020/schedule.html)
2. 数据库项目


### 1. 缓存穿透、击穿、雪崩
+ 穿透：指查询一个缓存和数据库都不存在的数据，导致尽管数据不存在但是每次都会到数据库查询，在访问量大的时候数据库可能挂掉。
+ 击穿：单个key值的缓存失效过期
+ 雪崩：redis缓存中大量的key同时失效，同时刚好有大量的请求，会直接访问数据库，造成数据库阻塞甚至宕机

### 2. 消息队列
+ 异步
+ 削峰
+ 解耦

//  mockgen -source=foo.go -package=mock -destination mock_foo.go


goland无限试用，ide-eval-reset，2021.2.2及之前版本。

## 计划
## 1. 后端
## 2. 嵌入式：树莓派->电机控制->四旋翼

#### [设计模式](./docs/design_pattern.md)
#### [Golang](./docs/golang.md)
#### [kubernetes](./docs/kubernetes.md)

#### 基础服务

containerd/lazy load/sidecar/daemonset

### zookeeper

### grpc
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
