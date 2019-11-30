FROM golang:latest
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
WORKDIR $GOPATH/src/wallet-service
COPY . $GOPATH/src/wallet-service
RUN go build .
#暴露端口
EXPOSE 9002
#最终运行docker的命令
ENTRYPOINT  ["./wallet-service"]