[contaienrd详解](https://www.rectcircle.cn/posts/)
1. 进程：
通过`ps xao pid,ppid,uid,cmd|grep containerd`查看
    + containerd守护进程通过`/run/containerd/containerd.sock`提供grpc接口
    + 每个容器都有一个守护进程containerd-shim-runc-v2管理容器的生命周期，通过`/run/containerd/s/`下的socket文件提供shim grpc server的接口。职责为执行runc命令启动容器、监控容器进程状态、容器1号进程被杀死后reap掉其所有子进程
2. 存储
    + `/var/lib/containerd/`存储持久化的数据
        + `io.containerd.content.v1.content`存储oci image
        + `io.containerd.metadata.v1.bolt`存储镜像、容器、快照的元数据
        + `io.containerd.snapshotter.v1`snapshotter快照目录
    + `/run/containerd/`存储临时数据
        + `containerd.sock`
        + `containerd.sock.ttrpc`低内存环境的grpc服务
        + `fifo`容器进程的stdin、stdout、stderr对接到目录下的fifo文件中
        + `io.containerd.runtime.v2.task/<namespace>/<name>`容器数据
    + `runc/<default>/<name>/state.json`容器状态文件
    + `s/xxx`与shim进程通讯的sock文件