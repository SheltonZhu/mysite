package servers

import (
	"context"
	"log"
	"mysite/dao"
	"mysite/grpc/pbs"
)

type BibleServiceServer struct {
	pbs.UnimplementedBibleServiceServer
}

func (p *BibleServiceServer) Get(ctx context.Context, id *pbs.Int32Value) (*pbs.Bible, error) {
	log.Println("bible_get")
	b, err := dao.GetBibleById(int(id.Value))
	return &pbs.Bible{Id: int32(b.Id), Text: b.Text}, err
}

func (p *BibleServiceServer) List(ctx context.Context, _ *pbs.Empty) (*pbs.Bibles, error) {
	log.Println("bible_list")
	bibles := make([]*pbs.Bible, 0)
	if b, err := dao.ListBibles(); err != nil {
		return &pbs.Bibles{}, err
	} else {
		for _, v := range b {
			pbsBible := pbs.Bible{Id: int32(v.Id), Text: v.Text}
			bibles = append(bibles, &pbsBible)
		}
	}
	return &pbs.Bibles{Bibles: bibles}, nil
}

func (p *BibleServiceServer) Create(ctx context.Context, text *pbs.StringValue) (*pbs.Int32Value, error) {
	log.Println("bible_create")
	b := dao.Bible{Text: text.Value}
	id, err := dao.CreateBible(&b)
	if err != nil {
		return &pbs.Int32Value{Value: int32(0)}, err
	}
	return &pbs.Int32Value{Value: int32(id)}, err
}

func (p *BibleServiceServer) Delete(ctx context.Context, id *pbs.Int32Value) (*pbs.Int32Value, error) {
	log.Println("bible_delete")
	count, err := dao.DeleteBible(int(id.Value))
	if err != nil {
		return &pbs.Int32Value{}, err
	}
	return &pbs.Int32Value{Value: int32(count)}, err
}
