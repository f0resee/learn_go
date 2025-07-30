# kubelet源码
## kubelet工作机制
### 1. 如何发现有新pod创建
   * 初始化一个[PodConfig](https://github.com/kubernetes/kubernetes/blob/4ce5a8954017644c5420bae81d72b09b735c21f0/pkg/kubelet/kubelet.go#L260), 之后添加文件、http路径、apiserver三个配置来源
   * 持续监听PodConfig给出的[updates](https://github.com/kubernetes/kubernetes/blob/4ce5a8954017644c5420bae81d72b09b735c21f0/cmd/kubelet/app/server.go#L1180)
   * 监听到update后，根据更新是ADD、UPDATE、REMOVE、RECONCILE、DELETE、SET进行处理[syncLoopIteration](https://github.com/kubernetes/kubernetes/blob/4ce5a8954017644c5420bae81d72b09b735c21f0/pkg/kubelet/kubelet.go#L2047)

### 2. 如何管理cpu

### 3. 拓扑管理
numa分区
#### a. memory manager
##### 官方文档
+ [Utilizing the NUMA-aware Memory Manager](https://kubernetes.io/docs/tasks/administer-cluster/memory-manager/)
+ [KEP-1769: Memory Manager](https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/1769-memory-manager#kep-1769-memory-manager)
##### 核心思路
实现了[MemoryManager](https://github.com/kubernetes/kubernetes/tree/release-1.24/pkg/kubelet/cm/memorymanager)，MemoryManger内包括：
+ Policy：用于决策对于一个给定内存需求的容器，怎么给其分配numa
+ State：持久化内存分配结果，读取/写入到`/var/lib/kubelet/memory_manager_state`文件中
+ runtimeService：更新容器的资源配置
+ ActivePodsFunc：获取node上所有存活的pod
+ PodStatusProvider：获取pod状态
+ ContainerMap：根据containerID查找*v1.Pod, *v1.Container
+ SourcesReady
+ []state.Block：allocatableMemory

其核心是policy，以一个有2个numa node(node0，node1)的节点为例：
+ 2个numa节点共有[0],[1],[0,1]三种组合
+ 分别判断三种组合是否可以满足容器内存需求，假设都可以，[0],[1]是preferred，但是[0]的序号更小，因此选择[0]作为best hint
+ 继续判断是否可以满足所有内存资源需求，如果不是，则需要扩大hint范围，选择[0,1]


#### b. cpu manager

