package models

import (
	"dwd-api/app/response"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Id        int32  `gorm:"primary_key; column:id" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	State     int    `gorm:"default:1;column:state" json:"state"`
	IsDel     int    `gorm:"default:0;column:is_del" json:"is_del"`
	CreatedAt uint32 `gorm:"column:created_at" json:"created_at"`
	UpdatedAt uint32 `gorm:"column:updated_at" json:"updated_at"`
}

func (Tag) TableName() string {
	return "tag"
}

func (t Tag) Info(db *gorm.DB) (*response.TagInfoResponse, error) {
	var tags = &response.TagInfoResponse{}
	err := db.Table("tag").Where("id=?", t.Id).Find(tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tag) List(db *gorm.DB, offset, limit int) ([]*response.TagListResponse, error) {
	tags := []*response.TagListResponse{}
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	if t.Name != "" {
		db = db.Where("name=?", t.Name)
	}
	err := db.Table("tag").Where("is_del = ?", 0).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name=?", t.Name)
	}
	err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
