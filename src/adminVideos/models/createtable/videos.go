package createtable

import (
	"github.com/jinzhu/gorm"
)

/*
	上传视频视频
*/
type Videos struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar;"json:"name"` //视频名称

	Date string  `gorm:"column:date;type:varchar;"json:"date"` //上映时间

	VideoPersons VideoPersons `gorm:"FOREIGNKEY:VideoPersonsId;"json:"video_persons_id"` //演员表
	VideoPersonId int `gorm:"column:video_persons_id;type:integer;"json:"video_persons_id"`

	Introduce string  `gorm:"column:introduce;type:varchar;"json:"introduce"` //视频简介

	CommentId int `gorm:"column:comment_id;type:integer;"json:"comment_id"` // 评论 一对多
	Comment []Comment `gorm:"FOREIGNKEY:VideosId;"json:"actor_id"`

	ImageId int `gorm:"column:image_id;type:integer;"json:"image_id"` //引入视频封面表
	Image Image `gorm:"FOREIGNKEY:ImageId;"json:"image"` //封面 belongs to 视频

	SrcId int `gorm:"column:src_id;type:integer;"json:"src_id"` //引入视频路径表
	Src Src `gorm:"FOREIGNKEY:SrcId;"json:"src_id"` //视频路径

	ClassesId int `gorm:"column:classes_id;type:integer;"json:"classes_id"`  //分类表的id
	Classes Classes `gorm:"FOREIGNKEY:ClassesId;"json:"classes"`

}



//视频路径
type Src struct {
	gorm.Model
	Patch string `gorm:"column:path;type:varchar;"json:"path"`
}

//视频封面
type Image struct {
	gorm.Model
	Path string `gorm:"column:path;type:varchar;"json:"path"`
}








