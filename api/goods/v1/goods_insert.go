package v1

import "github.com/gogf/gf/v2/frame/g"

type GoodsInsertReq struct {
	g.Meta `path:"/goods/insert" method:"post" tags:"GoodsService" summary:"Insert a goods"`
	Name   string `v:"required"`
	Num    int    `v:"required"`
}

type GoodsInsertRes struct{}
