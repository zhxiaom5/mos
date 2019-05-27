package annotations

import (
	"github.com/jinzhu/gorm"
)

type Annotation struct {
	IsRegion  bool   `gorm:"not null;column:is_region"`
	Project   string `gorm:"type:varchar(256);not null;column:project"`
	Text      string `gorm:"type:varchar(256);not null;column:text"`
	RegionID  uint   `gorm:"not null;column:region_id"`
	Tag       string `gorm:"type:varchar(256);not null;column:tag"`
	StartTime int64  `gorm:"not null;column:start_time"`
	EndTime   int64  `gorm:"not null;column:end_time"`
	gorm.Model
}

type AnnonPost struct {
	IsRegion  bool   `json:"is_region"`
	Project   string `json:"project"`
	Text      string `json:"text"`
	RegionID  uint   `json:"region_id"`
	Tag       string `json:"tag"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

// TableName Annotation 表名
func (Annotation) TableName() string {
	return "Annotation"
}

// // Create 插入Annotation记录
// func (a *Annotation) Create(p *AnnonPost) (res bool, err error) {
// 	a.Project = p.Project
// 	a.Text = p.Text
// 	a.RegionID = p.RegionID
// 	a.Tag = p.Tag
// 	a.PotTime = p.PotTime
// 	if err = glo.Db.Create(&a).Error; err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }
