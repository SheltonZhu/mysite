package servers

import (
	"context"
	"github.com/SheltonZhu/mysite/grpc/pbs"
	"log"
)

type HeartbeatServiceServer struct {
	pbs.UnimplementedHeartbeatServiceServer
}

func (p *HeartbeatServiceServer) Heartbeat(ctx context.Context, e *pbs.Empty) (*pbs.StringValue, error) {
	log.Println("pong_grpc")
	return &pbs.StringValue{Value: "pong_grpc"}, nil
}
