package services

import (
	"context"
	"mysite/dto"
	"mysite/grpc"
	"mysite/grpc/heartbeat"
	"time"
)

func Heartbeat() *dto.Result {
	message := "pong"
	if MasterAddr != "" {
		conn := grpc.InitClient(MasterAddr)
		client := heartbeat.NewHeartbeatServiceClient(conn)
		defer conn.Close()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		msg, _ := client.Heartbeat(ctx, &heartbeat.Empty{})
		message = msg.Value
	}
	return &dto.Result{Code: 0, Message: message}
}
