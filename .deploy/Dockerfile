FROM hub.class100.com/ops/golang:1.12-alpine

WORKDIR /app

COPY . .

ENV GO111MODULE on
RUN GOOS=linux CGO_ENABLED=0 \
  go build -mod=vendor -ldflags="-s -w" -installsuffix cgo -o eas main.go

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

CMD ["./eas", "server", "-c", "./config/config.yaml"]
