## linux

### 常用命令
#### system
+ ps
    ```shell
    ps -ef
    ps -p 1
    ps -o pid,ppid,cmd -p 3086311
    ps -p 3086311 -o cmd
    ```
+ top
+ dstat
+ sar
+ vmstat
+ pidstat
+ atop: [usage](https://www.digitalocean.com/community/tutorials/atop-command-in-linux)
+ strace
#### memory
   ```shell
  cat /proc/meminfo
   ```
+ free
+ slabtop
+ pmap

#### IO
+ iotop
+ iostat
+ blktrace
+ perf
+ stap
+ iosnoop

#### network
+ iftop
+ nethogs
+ ifstat
+ netstat
+ tcpdump
+ ip
+ ping
+ iperf
+ iptables

#### file
+ ldd
+ nm
+ size
+ readleaf
+ ls
+ objdump
+ du
+ df
+ dd
+ lsof

#### performance
+ perf
+ sysrq
+ ftrace
+ ebpf
+ pprof


linux settings 
```bash
sudo apt-get install ntpdate					//在Ubuntu下更新本地时间
sudo ntpdate time.windows.com
sudo hwclock --localtime --systohc			//将本地时间更新到硬件上
```

[cgroups](https://github.com/containerd/cgroups)

### 本机sock通信
### cgroups资源限制
### 硬件信息获取
### daemon set 

[linux kernel map](https://makelinux.github.io/kernel/map/)

[linux kernel docs](https://www.kernel.org/doc/html/latest/)

chsh -s /bin/zsh

/sbin/addgroup nansan sudo