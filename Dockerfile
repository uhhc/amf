FROM hub.docker.com/uhhc/amf-base:latest as builder

WORKDIR /workspace

# Build
COPY . ./
RUN make build

FROM alpine:3.10

WORKDIR /workspace

COPY --from=builder /go/bin/goose /usr/bin/goose
COPY --from=builder /workspace/bin/amf .
COPY --from=builder /workspace/assets ./assets
COPY --from=builder /workspace/db ./db

# Set time location to China
ENV TZ=Asia/Shanghai
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache tzdata mysql-client curl; \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone; \
    rm -rf /var/cache/apk/*

CMD ["/workspace/amf", "serve"]
