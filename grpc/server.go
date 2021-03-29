package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"mysite/dao"
	"mysite/grpc/bible"
	"mysite/grpc/heartbeat"
	"net"
)

func InitServer(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	log.Printf("grpc served at %d", port)
	if err != nil {
		log.Fatalf("grpc server start failed. failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	heartbeat.RegisterHeartbeatServiceServer(grpcServer, &heartbeatServiceServer{})
	bible.RegisterBibleServiceServer(grpcServer, &bibleServiceServer{})

	if err = grpcServer.Serve(listener); err != nil {
		panic(fmt.Sprintf("grpc serve at %d start failed.", port))
	}
}

type heartbeatServiceServer struct {
	heartbeat.UnimplementedHeartbeatServiceServer
}

func (p *heartbeatServiceServer) Heartbeat(ctx context.Context, e *heartbeat.Empty) (*heartbeat.StringValue, error) {
	return &heartbeat.StringValue{Value: "pong_grpc"}, nil
}

type bibleServiceServer struct {
	bible.UnimplementedBibleServiceServer
}

func (p *bibleServiceServer) Get(ctx context.Context, id *bible.Int32Value) (*bible.Bible, error) {
	b := dao.GetBibleById(int(id.Value))
	return &bible.Bible{Id: int32(b.Id), Text: b.Text}, nil
}
func (p *bibleServiceServer) List(ctx context.Context, _ *bible.Empty) (*bible.Bibles, error) {
	return &bible.Bibles{}, nil
}
func (p *bibleServiceServer) Create(ctx context.Context, text *bible.StringValue) (*bible.Int32Value, error) {
	b := dao.Bible{Text: text.Value}
	id, err := dao.CreateBible(&b)
	return &bible.Int32Value{Value: int32(id)}, err
}
