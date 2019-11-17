ARG ROLE

FROM golang:1.13-alpine

# Build Project
RUN apk add --no-cache git \
    && git clone https://github.com/NJUPT-ISL/Hound.git \
    && export GOMOD=$PWD/Hound/go.mod \
    && mkdir -p /root/$ROLE/log \
    && touch /root/$ROLE/log/$ROLE.log \
    && cd Hound/$ROLE \
    && go env \
    && export GO111MODULE=on \
    && go build main.go \
    && mv $ROLE /root/$ROLE \
    && chmod +x /root/$ROLE \
    && cd ../.. \
    && rm -rf Hound \
    && apk del git

EXPOSE [8080, 8081]

ENTRYPOINT ["/root/$ROLE/$ROLE"]
