## Docker
### 1. 常用命令
#### 1.1 安装
```bash
# 卸载旧版
sudo apt-get remove docker docker-engine docker.io containerd runc

# 安装证书并添加源
sudo apt update
sudo apt-get install ca-certificates curl gnupg lsb-release
curl -fsSL http://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -

# 添加源
sudo add-apt-repository "deb [arch=amd64] http://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable"

# 安装
sudo apt-get install docker-ce docker-ce-cli containerd.io

# 启动
systemctl start docker
```
#### 1.2 命令
```bash
# 拉镜像
docker pull nginx  # 默认最新版
docker pull nginx:1.20.1 # 指定版本

# 查看镜像
docker images

# 移除镜像
docker rmi 镜像名:版本号/镜像id

# 启动容器
docker run [OPTIONS] IMAGE [COMMAND] [ARG...]

#【docker run  设置项   镜像名  】 镜像启动运行的命令（镜像里面默认有的，一般不会写）
# -d：后台运行
# --restart=always: 开机自启
docker run --name=mynginx   -d  --restart=always -p  88:80   nginx

# 查看正在运行的容器
docker ps

# 查看所有
docker ps -a

# 删除容器
docker rm  容器id/名字
docker rm -f mynginx   #强制删除正在运行中的

#停止容器
docker stop 容器id/名字

#再次启动
docker start 容器id/名字

#应用开机自启
docker update 容器id/名字 --restart=always

# 进入容器内部的系统，修改容器内容
docker exec -it 容器id  /bin/bash

docker run --name=mynginx   \
-d  --restart=always \
-p  88:80 -v /data/html:/usr/share/nginx/html:ro  \ # 主机目录:容器
nginx

# 提交修改
docker commit [OPTIONS] CONTAINER [REPOSITORY[:TAG]]
docker commit -a "username"  -m "首页变化" 341d81f7504f mynginx:v1.0

# 将镜像保存成压缩包
docker save -o abc.tar mynginx:v1.0

# 别的机器加载这个镜像
docker load -i abc.tar

# 打tag并推送远端
docker tag local-image:tagname new-repo:tagname
docker push new-repo:tagname

# 把旧镜像的名字，改成仓库要求的新版名字
docker tag guignginx:v1.0 leifengyang/guignginx:v1.0

# 登录到docker hub
docker login       
docker logout（推送完成镜像后退出）

# 推送
docker push leifengyang/guignginx:v1.0


# 容器->主机
docker cp 5eff66eec7e1:/etc/nginx/nginx.conf  /data/conf/nginx.conf
# 主机->容器
docker cp  /data/conf/nginx.conf  5eff66eec7e1:/etc/nginx/nginx.conf
```