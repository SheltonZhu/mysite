package services

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
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

func ListBibles() (*dto.Result, error) {
	var data []dao.Bible
	var err error
	switch MasterAddr {
	case "":
		data, err = dao.ListBibles()
	default:
		tmpData := make([]dao.Bible, 0)
		GetBibleGrpcClient(func(ctx context.Context, client pbs.BibleServiceClient) {
			pbBibles, _ := client.List(ctx, &pbs.Empty{})
			for _, v := range pbBibles.Bibles {
				bible := dao.Bible{Id: int(v.Id), Text: v.Text}
				tmpData = append(tmpData, bible)
			}
			data = tmpData
		})
	}
	if err != nil {
		return &dto.Result{Code: 1, Message: "list failed.", Data: nil}, err
	}
	return &dto.Result{Code: 0, Message: "", Data: data}, nil
}
func GetBible(id int) (*dto.Result, error) {
	var data interface{}
	var err error
	switch MasterAddr {
	case "":
		data, err = dao.GetBibleById(id)
	default:
		GetBibleGrpcClient(func(ctx context.Context, client pbs.BibleServiceClient) {
			data, err = client.Get(ctx, &pbs.Int32Value{Value: int32(id)})
		})
	}
	if err != nil {
		return &dto.Result{Code: 1, Message: "get failed", Data: nil}, err
	}
	return &dto.Result{Code: 0, Message: "", Data: data}, nil
}

func CreateBible(text string) (*dto.Result, error) {
	b := dao.Bible{Text: text}
	var err error
	switch MasterAddr {
	case "":
		b.Id, err = dao.CreateBible(&b)
	default:
		var idWrapper *pbs.Int32Value
		GetBibleGrpcClient(func(ctx context.Context, client pbs.BibleServiceClient) {
			idWrapper, err = client.Create(ctx, &pbs.StringValue{Value: text})
		})
		if err == nil {
			b.Id = int(idWrapper.Value)
		}
	}
	if err != nil {
		log.Println(err)
		return &dto.Result{Code: 1, Message: "create failed.", Data: nil}, err
	}
	return &dto.Result{Code: 0, Message: "", Data: b}, nil
}

func DeleteBible(id int) (*dto.Result, error) {
	var deleted int
	var err error
	switch MasterAddr {
	case "":
		deleted, err = dao.DeleteBible(id)
	default:
		var count *pbs.Int32Value
		GetBibleGrpcClient(func(ctx context.Context, client pbs.BibleServiceClient) {
			count, err = client.Delete(ctx, &pbs.Int32Value{Value: int32(id)})
		})
		deleted = int(count.Value)
	}
	if err != nil {
		return &dto.Result{Code: 1, Message: "delete failed.", Data: nil}, err
	}
	return &dto.Result{Code: 0, Message: "delete success!", Data: gin.H{"deletedCount": deleted}}, nil
}
