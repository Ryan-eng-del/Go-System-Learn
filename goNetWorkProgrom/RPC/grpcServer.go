package rpc_server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"rpc/proto"

	"google.golang.org/grpc"
)

type ProductServer struct {
	proto.UnimplementedProductServer
}

func (*ProductServer) ProductInfo(context.Context, *proto.ProductRequest) (*proto.ProductResponse, error) {
	return &proto.ProductResponse{
		Id: 1,
		Name: "iphone",
		IsSale: false,
	}, nil
}

var port = flag.Int("port", 5051, "The GRPC Server Port")

func RPCServer() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterProductServer(grpcServer, &ProductServer{})
	log.Printf("GRPC Server listening on %s", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}



