# 第一阶段：构建应用程序
FROM golang:1.23 AS builder
# 定义构建参数，在构建镜像的时候需要传入 SERVICE 来表示构建的模块
ARG SERVICE

# 设置环境变量
ENV TZ=Asia/Shanghai
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct
ENV GOOS=linux
ENV GOARCH=amd64
# 创建工作目录
RUN mkdir -p /app
RUN mkdir -p /app/output

WORKDIR /app

# 复制所有文件到工作目录
COPY . .

# 下载依赖
RUN go mod tidy

# 编译应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ${SERVICE} ./cmd/${SERVICE}/main.go

# 第二阶段：创建最终运行环境
FROM alpine

ARG SERVICE

# 环境变量
ENV TZ=Asia/Shanghai
ENV SERVICE=${SERVICE}

# 换源，更新软件依赖
RUN sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#https://mirrors.tuna.tsinghua.edu.cn/alpine#g' /etc/apk/repositories
RUN apk update --no-cache && apk --no-cache add ca-certificates tzdata bash

# 创建工作目录
WORKDIR /app

# 复制 Binary
COPY --from=builder /app/${SERVICE} /app/${SERVICE}

CMD ["./${SERVICE}"]




