ARG ROLE

FROM golang:1.13-alpine

# Build Project
RUN apk add --no-cache git \
    && git clone https://github.com/NJUPT-ISL/Hound.git \
    && mkdir -p /root/$ROLE/log \
    && touch /root/$ROLE/log/$ROLE.log \
    && cd Hound/$ROLE \
    && go build . -o /root/$ROLE \
    && chmod +x /root/$ROLE \
    && cd ../.. \
    && rm -rf Hound \
    && apk del git

EXPOSE [8080, 8081]

ENTRYPOINT ["/ROOT/$ROLE"]
