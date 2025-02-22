# kubelet源码
1. 如何发现有新pod创建
    * 初始化一个[PodConfig](https://github.com/kubernetes/kubernetes/blob/4ce5a8954017644c5420bae81d72b09b735c21f0/pkg/kubelet/kubelet.go#L260), 之后添加文件、http路径、apiserver三个配置来源
    * 持续监听PodConfig给出的[updates](https://github.com/kubernetes/kubernetes/blob/4ce5a8954017644c5420bae81d72b09b735c21f0/cmd/kubelet/app/server.go#L1180)
    * 监听到update后，根据更新是ADD、UPDATE、REMOVE、RECONCILE、DELETE、SET进行处理[syncLoopIteration](https://github.com/kubernetes/kubernetes/blob/4ce5a8954017644c5420bae81d72b09b735c21f0/pkg/kubelet/kubelet.go#L2047)