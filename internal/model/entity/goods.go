// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Goods is the golang structure for table goods.
type Goods struct {
	Id       uint        `json:"id"       orm:"id"        description:"id"`
	Name     string      `json:"name"     orm:"name"      description:"goods name"`
	Num      int         `json:"num"      orm:"num"       description:"number"`
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:"create time"`
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" description:"update time"`
}
