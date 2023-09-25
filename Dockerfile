# 使用一个基础的 Golang 镜像
FROM golang:alpine as build

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

# 设置工作目录
WORKDIR /app

# 复制 Golang 项目的源代码到容器中
COPY . /app

# 在容器中编译 Golang 项目
RUN go build -o blog cmd/main.go

# 创建最终的生产镜像
FROM alpine:latest as prod

# 设置工作目录
WORKDIR /app

# 从之前的阶段复制二进制文件
COPY --from=build /app/blog /app/
COPY --from=build /app/internal/conf/* /app/internal/conf/

# 设置环境变量等
ENV PORT=18080

# 暴露端口
EXPOSE 18080

# 启动应用
CMD ["./blog"]
