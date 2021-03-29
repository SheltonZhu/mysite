package services

import (
	"context"
	"mysite/dao"
	"mysite/dto"
	"mysite/grpc"
	"mysite/grpc/bible"
	"time"
)

func ListBibles() *dto.Result {
	return &dto.Result{Code: 0, Message: "", Data: dao.ListBibles()}
}
func GetBible(id int) *dto.Result {
	var data interface{}
	var _ error
	if MasterAddr != "" {
		conn := grpc.InitClient(MasterAddr)
		client := bible.NewBibleServiceClient(conn)
		defer conn.Close()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		data, _ = client.Get(ctx, &bible.Int32Value{Value: int32(id)})
	} else {
		data = dao.GetBibleById(id)
	}
	return &dto.Result{Code: 0, Message: "", Data: data}
}
