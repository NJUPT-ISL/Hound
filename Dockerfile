

FROM golang:1.13-alpine

# Build Project
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && apk add --no-cache gcc g++

# Create DIR
RUN mkdir /Hound
COPY .  /Hound

# Set Build ENV
ENV GO111MODULE on
ENV GOPROXY https://goproxy.io
ARG ROLE
RUN cd /Hound \
    && go build ./$ROLE/main.go \
    && chmod +x main \
    && apk del gcc g++

