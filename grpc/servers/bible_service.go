package servers

import (
	"context"
	"mysite/dao"
	"mysite/grpc/pbs"
)

type BibleServiceServer struct {
	pbs.UnimplementedBibleServiceServer
}

func (p *BibleServiceServer) Get(ctx context.Context, id *pbs.Int32Value) (*pbs.Bible, error) {
	b := dao.GetBibleById(int(id.Value))
	return &pbs.Bible{Id: int32(b.Id), Text: b.Text}, nil
}

func (p *BibleServiceServer) List(ctx context.Context, _ *pbs.Empty) (*pbs.Bibles, error) {
	bibles := make([]*pbs.Bible, 0)
	for _, v := range dao.ListBibles() {
		pbsBible := pbs.Bible{Id: int32(v.Id), Text: v.Text}
		bibles = append(bibles, &pbsBible)
	}
	return &pbs.Bibles{Bibles: bibles}, nil
}

func (p *BibleServiceServer) Create(ctx context.Context, text *pbs.StringValue) (*pbs.Int32Value, error) {
	b := dao.Bible{Text: text.Value}
	id, err := dao.CreateBible(&b)
	return &pbs.Int32Value{Value: int32(id)}, err
}
