package goods

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/internal/dao"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/model/do"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

type (
	sGoods struct{}
)

func init() {
	service.RegisterGoods(New())
}

func New() *sGoods {
	return &sGoods{}
}

func (s *sGoods) Create(ctx context.Context, in model.GoodsInput) (err error) {

	_, err = dao.Goods.Ctx(ctx).Data(do.Goods{
		Name: in.Name,
		Num:  in.Num,
	}).Insert()

	return err
}
