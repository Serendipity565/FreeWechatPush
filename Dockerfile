# 使用官方 Go 镜像作为构建阶段
FROM golang:1.23 AS builder

# 设置工作目录
WORKDIR /hello

# 将代码复制到容器
COPY . /hello

ENV GOPROXY https://goproxy.cn,direct

# 下载依赖并构建
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app main.go

# 使用更小的基础镜像运行应用
FROM alpine:latest

# 设置工作目录
WORKDIR /hello

# 复制构建的二进制文件
COPY --from=builder /hello/app .

# 运行程序
CMD ["./app"]
