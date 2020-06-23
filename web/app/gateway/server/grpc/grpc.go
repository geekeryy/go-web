package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"go-web/web/app/gateway/service/user"
	pb_gateway "go-web/web/protobuf/gateway"
)

func Init() {
	listen, err := net.Listen("tcp", ":8888")
	if err!=nil {
		log.Fatalln("grpc err:",err)
	}
	svr:=grpc.NewServer()
	pb_gateway.RegisterServiceServer(svr,user.NewUser())
	if err := svr.Serve(listen);err!=nil{
		log.Fatal("grpc listen err:",err)
	}

}