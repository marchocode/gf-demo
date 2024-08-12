package goods

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/goods/v1"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

func (c *ControllerV1) GoodsInsert(ctx context.Context, req *v1.GoodsInsertReq) (res *v1.GoodsInsertRes, err error) {

	err = service.Goods().Create(ctx, model.GoodsInput{
		Name: req.Name,
		Num:  req.Num,
	})
	return
}
