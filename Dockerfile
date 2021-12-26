FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /home

# 将代码复制到容器中
COPY ./ .

# 将我们的代码编译成二进制可执行文件  可执行文件名为 app
RUN go build -o app .

# 移动到用于存放生成的二进制文件的 /dist 目录
WORKDIR /dist

# 将二进制文件从 /home/app 目录复制到这里
COPY --from=builder /home/app /dist/

# 声明服务端口
EXPOSE 9000

# 启动容器时运行的命令
ENTRYPOINT ["/dist/app"]

