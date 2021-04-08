package services

import (
	"context"
	"mysite/dao"
	"mysite/dto"
	"mysite/grpc"
	"mysite/grpc/pbs"
	"time"
)

func GetBibleGrpcClient(f func(ctx context.Context, client pbs.BibleServiceClient)) {
	conn := grpc.InitClient(MasterAddr)
	client := pbs.NewBibleServiceClient(conn)
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	f(ctx, client)
}

func ListBibles() *dto.Result {
	var data []dao.Bible
	var _ error

	if MasterAddr != "" {
		tmpData := make([]dao.Bible, 0)
		GetBibleGrpcClient(func(ctx context.Context, client pbs.BibleServiceClient) {
			pbBibles, _ := client.List(ctx, &pbs.Empty{})
			for _, v := range pbBibles.Bibles {
				bible := dao.Bible{Id: int(v.Id), Text: v.Text}
				tmpData = append(tmpData, bible)
			}
			data = tmpData
		})
	} else {
		data = dao.ListBibles()
	}
	return &dto.Result{Code: 0, Message: "", Data: data}
}
func GetBible(id int) *dto.Result {
	var data interface{}
	var _ error
	if MasterAddr != "" {
		conn := grpc.InitClient(MasterAddr)
		client := pbs.NewBibleServiceClient(conn)
		defer conn.Close()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		data, _ = client.Get(ctx, &pbs.Int32Value{Value: int32(id)})
	} else {
		data = dao.GetBibleById(id)
	}
	return &dto.Result{Code: 0, Message: "", Data: data}
}
