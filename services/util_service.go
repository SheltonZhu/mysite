package services

import (
	"context"
	"github.com/SheltonZhu/mysite/dto"
	"github.com/SheltonZhu/mysite/grpc"
	"github.com/SheltonZhu/mysite/grpc/pbs"
	"time"
)

func Heartbeat() *dto.Result {
	message := "pong"
	if MasterAddr != "" {
		conn := grpc.InitClient(MasterAddr)
		client := pbs.NewHeartbeatServiceClient(conn)
		defer conn.Close()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		msg, _ := client.Heartbeat(ctx, &pbs.Empty{})
		message = msg.Value
	}
	return &dto.Result{Code: 0, Message: message}
}
