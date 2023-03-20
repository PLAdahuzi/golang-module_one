# 指定基础镜像 docker官网公开的基础镜像地址 https://hub.docker.com/_/golang
FROM golang:alpine

# 设置环境变量
# GO111MODULE: on 开启GO111MODULE 解决导包问题
# GOPROXY: 使用国内代理
# CGO_ENABLED： 当CGO_ENABLED=1， 进行编译时， 会将文件中引用libc的库（比如常用的net包），以动态链接的方式生成目标文件。
#               当CGO_ENABLED=0， 进行编译时， 则会把在目标文件中未定义的符号（外部函数）一起链接到可执行文件中。
# GOOS: 指定在什么系统环境中运行
# GOARCH： 指定运行的系统是多少位
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录中
WORKDIR /build

#将代码复制到容器中
COPY . .

RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep

#初始化go项目
#RUN go mod init golang-module_one
# 编译代码 编译成可执行的二进制文件 应用的名字叫做 module_one
RUN go build -o module_one .

# 移动到用于存放生成的二进制文件的 /dist 目录
WORKDIR /dist

# 将二进制文件从 /build 目录复制到这里 /dist
RUN cp /build/module_one .

# 声明服务端口
EXPOSE 8888

# 启动容器时需要启动的应用
ENTRYPOINT ["module_one"]

# 定义启动容器应用的时候的参数
CMD ["--help"]