package model

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	Vid         uint    `gorm:"comment:'所属视频';index"`
	Uid         uint    `gorm:"comment:'所属用户';index"`
	Title       string  `gorm:"type:varchar(50);comment:'分P使用的标题'"`
	Url         string  `gorm:"type:varchar(255);comment:'视频链接'"`
	OriginalUrl string  `gorm:"type:varchar(255);comment:'原始mp4链接'"`
	Duration    float64 `gorm:"comment:'视频时长';default:0"`
	Status      int     `gorm:"comment:'审核状态';not null;index"`
	Quality     int     `gorm:"comment:'视频最大质量';"`
}

func (table *Resource) TableName() string {
	return "resource"
}
