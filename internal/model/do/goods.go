// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Goods is the golang structure of table goods for DAO operations like Where/Data.
type Goods struct {
	g.Meta   `orm:"table:goods, do:true"`
	Id       interface{} // id
	Name     interface{} // goods name
	Num      interface{} // number
	CreateAt *gtime.Time // create time
	UpdateAt *gtime.Time // update time
}
