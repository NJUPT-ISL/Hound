ARG ROLE

FROM golang:1.13-alpine

# Build Project
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && apk add --no-cache gcc g++
COPY $ROLE /root
COPY go.mod /root
COPY go.sum /root
RUN mkdir -p /root/$ROLE/log \
    && export GO111MODULE=on \
    && export GOPROXY=https://goproxy.io \ 
    && go build /root/$ROLE/main.go \
    && mv main /root/$ROLE \
    && chmod +x /root/$ROLE/main \
    && apk del gcc g++


