# 使用官方Go镜像作为构建环境
FROM golang:1.21 as builder

# 设置工作目录
WORKDIR /app

# 复制go.mod文件
COPY go.mod ./

# 下载依赖
RUN go mod download

# 复制源代码到容器内
COPY . .

# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o be101-golang

# 使用scratch作为最小运行环境
FROM scratch

# 从构建器镜像中复制构建的可执行文件
COPY --from=builder /app/be101-golang /app/

# 设置运行时的工作目录
WORKDIR /app

# 设置容器启动时运行的命令
ENTRYPOINT ["./be101-golang"]