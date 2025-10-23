
## 二、主要组件
 ![输入图片说明](../img/kubernetes.png)

### 控制面组件
### 1. API server(kube-apiserver)
API server是k8s控制面中的一个组件，暴露k8s API，是k8s控制面的前端。其主要实现是kube-apiserver，可以水平缩放，通过部署更多的实例来进行缩放。可以运行多个kube-apiserver，并在实例间进行负载均衡。

### 2. etcd
完备、高可用key-value存储，用于k8s后端存储所有的集群数据。
```sh
# 查看etcd容器启动命令
kubectl -n kube-system describe pod etcd-master01
# 查看etcd中的所有key
./etcdctl --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/peer.crt --key=/etc/kubernetes/pki/etcd/peer.key get --prefix "" --keys-only
# 查看etcd中的所有key并解析为json
./etcdctl --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/peer.crt --key=/etc/kubernetes/pki/etcd/peer.key get --prefix "" --keys-only -w json|python3 -m json.tool
# 查看某个key对应的字符串
base64 -d <<<L3JlZ2lzdHJ5L3NlcnZpY2VzL3NwZWNzL2t1YmVybmV0ZXMtZGFzaGJvYXJkL2t1YmVybmV0ZXMtZGFzaGJvYXJk
# etcdhelper查看key对应的值，https://github.com/openshift/origin/tree/master/tools/etcdhelper
./etcdhelper --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/peer.crt --key=/etc/kubernetes/pki/etcd/peer.key get /registry/clusterroles/system:controller:endpointslice-controller
```

### 3. kube-scheduler
用于监听新创建的Pods和未分配的node，并选择node来运行这些Pod。调度时考虑的因素：资源需求、硬件/软件/策略约束、亲和/反亲和要求、数据局部性、工作负载间的影响及生命周期。

### 4. kube-controller-manager
是运行controller进程的控制面组件。每个controller是一个单独的过程，但是都被编译成一个二进制文件，在一个进程中运行。

### 节点组件
### 5. kubelet
在每个Node节点上运行的主要 “节点代理”。它接受通过各种机制(主要是kube-apiserver)提供的一组PodSpec(描述pod 的YAML或JSON对象)并保证这些PodSpec中描述的pod正常健康运行。

### 6. kube-proxy
维护节点网络规则，网络规则允许从集群内或者从集群外通过网络连接访问Pod。

## 三、基础概念

[Kubernetes核心实战](https://www.yuque.com/leifengyang/oncloud/ctiwgo#3ykv9)
[Kubernetes](https://iximiuz.com/en/categories/?category=Kubernetes)
[kubernetes乱世浮生](https://atbug.com/deep-dive-k8s-network-mode-and-communication/)
[kubernetes notes](https://github.com/rfyiamcool/notes)

### 1. Master节点
### 2. Node节点
```sh
kubectl get node
kubectl describe node worker02
```
### 3. Namespace
```sh
kubectl get ns
```
### 4. Pod
```sh
kubectl get pod -A
kubectl get pod -n test  
kubectl get pods -A -o wide
kubectl describe pod mynginx -n test
kubectl logs mynginx -n test
kubectl exec -it mynginx -n test -- /bin/bash
kubectl exec -it mynginx -n test -c nginx -- /bin/bash
```

```yaml
kind: Pod
apiVersion: v1
metadata:
  name: mynginx
  containers:
    - name: mynginx
      image: nginx
```

```yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    run: myapp
  name: myapp
spec:
  containers:
  - image: nginx
    name: nginx
  - image: tomcat
    name: tomcat
```
