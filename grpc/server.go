package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"mysite/dao"
	"mysite/grpc/pb"
	"net"
)

func InitServer(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	log.Printf("grpc served at %d", port)
	if err != nil {
		log.Fatalf("grpc server start failed. failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHeartbeatServiceServer(grpcServer, &heartbeatServiceServer{})
	pb.RegisterBibleServiceServer(grpcServer, &bibleServiceServer{})

	if err = grpcServer.Serve(listener); err != nil {
		panic(fmt.Sprintf("grpc serve at %d start failed.", port))
	}
}

type heartbeatServiceServer struct {
	pb.UnimplementedHeartbeatServiceServer
}

func (p *heartbeatServiceServer) Heartbeat(ctx context.Context, e *pb.Empty) (*pb.StringValue, error) {
	return &pb.StringValue{Value: "pong_grpc"}, nil
}

type bibleServiceServer struct {
	pb.UnimplementedBibleServiceServer
}

func (p *bibleServiceServer) Get(ctx context.Context, id *pb.Int32Value) (*pb.Bible, error) {
	b := dao.GetBibleById(int(id.Value))
	return &pb.Bible{Id: int32(b.Id), Text: b.Text}, nil
}
func (p *bibleServiceServer) List(ctx context.Context, _ *pb.Empty) (*pb.Bibles, error) {
	return &pb.Bibles{}, nil
}
func (p *bibleServiceServer) Create(ctx context.Context, text *pb.StringValue) (*pb.Int32Value, error) {
	b := dao.Bible{Text: text.Value}
	id, err := dao.CreateBible(&b)
	return &pb.Int32Value{Value: int32(id)}, err
}
