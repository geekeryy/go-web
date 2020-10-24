FROM golang:1.15 as builder

WORKDIR /app

COPY . .

ENV GO111MODULE on

RUN GOOS=linux CGO_ENABLED=0 \
  go build -mod=vendor -ldflags="-s -w" -installsuffix cgo -o server main.go


FROM golang:1.15 as runner

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/config.yaml .

EXPOSE 8080

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

CMD ["./server", "http", "-c", "./config.yaml"]
