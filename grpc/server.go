package grpc

import (
	"fmt"
	"github.com/SheltonZhu/mysite/grpc/pbs"
	. "github.com/SheltonZhu/mysite/grpc/servers"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitServer(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	log.Printf("grpc served at %d", port)
	if err != nil {
		log.Fatalf("grpc server start failed. failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pbs.RegisterHeartbeatServiceServer(grpcServer, &HeartbeatServiceServer{})
	pbs.RegisterBibleServiceServer(grpcServer, &BibleServiceServer{})

	if err = grpcServer.Serve(listener); err != nil {
		panic(fmt.Sprintf("grpc serve at %d start failed.", port))
	}
}
