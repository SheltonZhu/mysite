package services

import (
	"context"
	"mysite/dto"
	"mysite/grpc"
	"mysite/grpc/pb"
	"time"
)

func Heartbeat() *dto.Result {
	message := "pong"
	if MasterAddr != "" {
		conn := grpc.InitClient(MasterAddr)
		client := pb.NewHeartbeatServiceClient(conn)
		defer conn.Close()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		msg, _ := client.Heartbeat(ctx, &pb.Empty{})
		message = msg.Value
	}
	return &dto.Result{Code: 0, Message: message}
}
