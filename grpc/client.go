package grpc

import (
	"google.golang.org/grpc"
	"log"
)

func InitClient(serverAddr string) *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	log.Printf("%v", serverAddr)
	return conn
}
