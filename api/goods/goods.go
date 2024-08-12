// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package goods

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/api/goods/v1"
)

type IGoodsV1 interface {
	GoodsInsert(ctx context.Context, req *v1.GoodsInsertReq) (res *v1.GoodsInsertRes, err error)
}
