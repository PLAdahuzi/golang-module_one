# 指定基础镜像 docker官网公开的基础镜像地址 https://hub.docker.com/_/golang
FROM golang:1.20.2 AS builder

# 设置环境变量
# GO111MODULE: on 开启GO111MODULE 解决导包问题
# GOPROXY: 使用国内代理
# CGO_ENABLED： 当CGO_ENABLED=1， 进行编译时， 会将文件中引用libc的库（比如常用的net包），以动态链接的方式生成目标文件。
#               当CGO_ENABLED=0， 进行编译时， 则会把在目标文件中未定义的符号（外部函数）一起链接到可执行文件中。
# GOOS: 指定在什么系统环境中运行
# GOARCH： 指定运行的系统是多少位
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.io,direct

# 移动到工作目录中
WORKDIR /build

#将代码复制到容器中
COPY . .

# 编译代码 编译成可执行的二进制文件 应用的名字叫做 module_one
# -installsuffix：为了使当前的输出目录与默认的编译输出目录分离，可以使用这个标记。此标记的值会作为结果文件的父目录名称的后缀。
#       其实，如果使用了-race标记，这个标记会被自动追加且其值会为race。如果我们同时使用了-race标记和-installsuffix，
#       那么在-installsuffix标记的值的后面会再被追加_race，并以此来作为实际使用的后缀。
# -o：编译指定输出到的文件。
# cgo：
RUN go build -installsuffix cgo -o module_one .

FROM alpine:3.17.0
WORKDIR /app

COPY --from=builder /build/module_one .
# 声明服务端口
EXPOSE 8888

# 启动容器时需要启动的应用
ENTRYPOINT ["/app/module_one"]