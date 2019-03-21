package createtable

import (
	"github.com/jinzhu/gorm"
)


// 无限极分类表
type Classes struct {
	gorm.Model
	Name  string  `gorm:"column:name;type:varchar;"json:"name"` // 类名
	ParentId int `gorm:"column:parent_id;type:integer;DEFAULT:0;"json:"parent_id"` // 父母分类的id  默认值是0
	ParentList  string  `gorm:"column:parent_list;type:varchar;DEFAULT:0;"json:"parent_list"` // 分类的层级关系，从最高级到自己 默认值是0
	Depth int  `gorm:"column:depth;type:integer;DEFAULT:1;"json:"depth"` // 深度  假设1,2,3  那么就是3层  默认值是1
	Status int  `gorm:"column:status;type:integer;DEFAULT:1;"json:"status"`  // 0是禁用  1是启用  默认值是1
	Priority string `gorm:"column:priority;type:varchar;DEFAULT:0;"json:"priority"`  //Priority  默认值是0
}
