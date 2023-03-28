### init
```
# 将工程 clone 到 gopath 的src目录下 然后执行如下命令
go mod init golang-module_one
go mod tidy
```

### builder
```
#  -t 表示给镜像取名字
docker build . -t go_module_one
```

### run
```
# 语法： docker run --name 设置容器名 -it -d -p [宿主机端口号]:[容器启动端口号] 要启动的镜像
# -d 作用在镜像启动之后不进入容器内部，使容器在后台运行
# -i 交互式操作
# -t 终端
docker run --name goModuleOne  -it -d  -p 8888:8888 go_module_one
```

### push image to DockerHub
```
# 语法：docker push [OPTIONS] NAME[:TAG]
# 登录docker 输入用户名  密码
docker login 

# 打标签并设置版本号  用户名/镜像名：版本
docker tag  go_module_one:latest pladahuzi/go_module_one:v1.0
docker push pladahuzi/go_module_one:v1.0
# 
```

### 通过 nsenter 进入容器查看 IP 配置
```
# 查看运行的容器
docker ps

# 获取pid
# 370c4bf0dc36 是容器运行的 CONTAINER ID 或者使用启动容器时指定的 --name 的值
PID=$(docker inspect --format "{{ .State.Pid}}" 370c4bf0dc36)
# or
PID=$(docker inspect --format "{{ .State.Pid}}" goModuleOne)
# 或者直接查看 
# -f 等效于 --format
docker inspect -f {{.State.Pid}} goModuleOne

# 进入网络命名空间
# -t, --target pid：指定被进入命名空间的目标进程的pid
# -m, --mount[=file]：进入mount命令空间。如果指定了file，则进入file的命令空间
# -u, --uts[=file]：进入uts命令空间。如果指定了file，则进入file的命令空间
# -i, --ipc[=file]：进入ipc命令空间。如果指定了file，则进入file的命令空间
# -n, --net[=file]：进入net命令空间。如果指定了file，则进入file的命令空间
# -p, --pid[=file]：进入pid命令空间。如果指定了file，则进入file的命令空间
# -U, --user[=file]：进入user命令空间。如果指定了file，则进入file的命令空间
# -G, --setgid gid：设置运行程序的gid
# -S, --setuid uid：设置运行程序的uid
# -r, --root[=directory]：设置根目录
# -w, --wd[=directory]：设置工作目录
nsenter -n -t 13212
# 查看网络
ip a
# 退出
exit

# 查看网络 
nsenter --target $PID --mount --uts --ipc --net --pid ip a && ip r
```


### k8s 安装
[参考文档](https://u.geekbang.org/lesson/482?article=608883&utm_source=time_web&utm_medium=menu&utm_term=timewebmenu&utm_identify=geektime&utm_content=menu&utm_campaign=timewebmenu&gk_cus_user_wechat=university)
###### 遇到的问题
```
    问题1. 22:16:28 ubantu-01 kubelet[39726]: E0328 22:16:28.220696   39726 server.go:294] "Failed to run kubelet" err="failed to run Kubelet: misconfiguration: kubelet cgroup driv>
    
    解决: docker 没有设置Cgroup Driver，默认使用的是 cgroupfs ，k8s推荐使用的 Cgroup Driver 是 systemd，所以报错
    ps：docker默认的Cgroup Driver是cgroupfs ，cgroupfs是cgroup为给用户提供的操作接口而开发的虚拟文件系统类型，可以向用户展示cgroup的hierarchy，通知kernel用户对cgroup改动，对cgroup的查询和修改只能通过cgroupfs文件系统来进行
        Kubernetes 推荐使用 systemd 来代替 cgroupfs，  因为systemd是Kubernetes自带的cgroup管理器, 负责为每个进程分配cgroups， 但docker的cgroup driver默认是cgroupfs，这样就同时运行有两个cgroup控制管理器，当资源有压力的情况时，有可能出现不稳定的情况

更新 cgroupdriver 为 systemd
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://uy35zvn6.mirror.aliyuncs.com"],
  "exec-opts": ["native.cgroupdriver=systemd"]
}
EOF
 
systemctl daemon-reload
systemctl restart docker
```

















