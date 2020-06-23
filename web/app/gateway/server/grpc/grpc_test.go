package grpc_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"google.golang.org/grpc"

	base_pb "go-web/web/protobuf/base"
	pb_gateway "go-web/web/protobuf/gateway"
)

func TestDemo3(t *testing.T)  {
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := pb_gateway.NewServiceClient(conn)
	res, err := client.Demo3(context.Background(), &base_pb.Base{
		Code: 1,
		Msg:  "ok",
	})
	fmt.Println(res,err)

}
