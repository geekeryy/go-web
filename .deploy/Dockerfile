FROM golang:1.15-alpine as builder

COPY . .

ENV GO111MODULE on

RUN GOOS=linux CGO_ENABLED=0 \
  go build -mod=vendor -ldflags="-s -w" -installsuffix cgo -o server main.go


FROM golang:1.15-alpine as runner

WORKDIR /app

COPY --from=builder /server /

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

CMD ["./server", "http", "-c", "./config.yaml"]
