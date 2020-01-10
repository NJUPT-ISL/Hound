ARG ROLE

FROM golang:1.13-alpine

# Install Dep
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && apk add --no-cache gcc g++ git

# Build Project
RUN git clone https://github.com/NJUPT-ISL/Hound \
    && export GO111MODULE=on \
    && export GOPROXY=https://goproxy.io \ 
    && cd Hound/$ROLOE/main.go \
    && go build . \
    && chmod +x /root//$ROLE/$ROLE \
    && apk del gcc g++

ENTRYPOINT ["/root/Hound/$ROLE/$ROLE"]
