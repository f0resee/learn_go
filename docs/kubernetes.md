## 一、安装
[kubernetes](https://kubernetes.io/)
### 1. 准备
```shell
swapoff -a      //关闭swap，立即生效
vim /etc/fstabs //关闭swap
```

### 2.下载安装containerd
[containerd](https://github.com/containerd/containerd)
```shell
tar Cxzvf /usr/local containerd-1.6.2-linux-amd64.tar.gz //解压
mkdir -p /usr/local/lib/systemd/system
//创建 containerd.service并启动
systemctl daemon-reload
systemctl enable --now containerd
install -m 755 runc.amd64 /usr/local/sbin/runc
containerd config default > /etc/containerd/config.toml
vim /etc/containerd/config.toml // 修改SystemdGroup为true
sudo systemctl restart containerd
```

### 3. kuberadm
```shell
apt-get update && apt-get install -y apt-transport-https
curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add - 
cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main
EOF
apt-get update
apt-get install -y kubelet kubeadm kubectl
```

### 4. 
```
hostnamectl set-hostname master01  
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
overlay
br_netfilter
EOF

sudo modprobe overlay
sudo modprobe br_netfilter

# sysctl params required by setup, params persist across reboots
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-iptables  = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.ipv4.ip_forward                 = 1
EOF

# Apply sysctl params without reboot
sudo sysctl --system

kubeadm config images pull --image-repository registry.aliyuncs.com/google_containers
vim /etc/containerd/config.toml // sandbox : registry.aliyuncs.com/google_containers/pause:3.9
kubeadm init --pod-network-cidr=10.244.0.0/16 --image-repository registry.aliyuncs.com/google_containers

```

### 5. 安装dashboard
```shell
// https://github.com/kubernetes/dashboard
kubectl apply -f recommended.yaml  
kubectl edit svc kubernetes-dashboard -n kubernetes-dashboard  //type：ClusterIP修改为NodePort
kubectl get svc -A |grep kubernetes-dashboard 
kubectl -n kubernetes-dashboard create token admin-user
```


### 6. minikube
https://minikube.sigs.k8s.io/docs/start/

https://www.jeeinn.com/2022/07/1715/

minikube stop

minikube delete

minikube start --kubernetes-version=v1.23.8

## 二、主要组件
 ![输入图片说明](../img/kubernetes.png)

### 控制面组件
### 1. API server(kube-apiserver)
API server是k8s控制面中的一个组件，暴露k8s API，是k8s控制面的前端。其主要实现是kube-apiserver，可以水平
缩放，通过部署更多的实例来进行缩放。可以运行多个kube-apiserver，并在实例间进行负载均衡。

### 2. etcd
完备、高可用key-value存储，用于k8s后端存储所有的集群数据。

### 3. kube-scheduler
用于监听新创建的Pods和未分配的node，并选择node来运行这些Pod。调度时考虑的因素：资源需求、硬件/软件/策略约束、亲和/反亲和要求、数据局部性、工作负载间的影响及生命周期。

### 4. kube-controller-manager
是运行controller进程的控制面组件。每个controller是一个单独的过程，但是都被编译成一个二进制文件，在一个进程中运行。

### 节点组件
### 5. kubelet

### 6. kube-proxy
维护节点网络规则，网络规则允许从集群内或者从集群外通过网络连接访问Pod。
## 三、源码