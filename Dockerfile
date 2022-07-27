FROM golang:1.16.6-alpine

#创建文件夹
RUN mkdir ./dist

#将Dockerfile 中的文件存储到/app下
ADD . /dist

# 设置工作目录
WORKDIR /dist

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn,direct"
# 因为已经是在 /app下了，所以使用  ./
RUN go build -o main ./main.go ./index.go

# 暴露的端口
EXPOSE 8081

#设置容器的启动命令，CMD是设置容器的启动指令
CMD /dist/main

