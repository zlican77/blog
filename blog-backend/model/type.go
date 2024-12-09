package model

import (
	"time"
)

type Articles struct {
	ArticleId  int       `gorm:"primaryKey;column:article_id;autoIncrement"`
	Title      string    `gorm:"column:title;type:varchar(255)"`
	Excerpt    string    `gorm:"column:excerpt;type:varchar(500)"`
	Content    string    `gorm:"column:content;type:text"`
	CategoryId int       `gorm:"column:category_id"`
	ViewCount  int       `gorm:"column:view_count;default:0"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;not null"`
	Status     int       `gorm:"column:status;type:tinyint;default:1"`
	Url        string    `gorm:"column:url;type:varchar(255)"`

	Category Category `gorm:"foreignKey:CategoryId;references:CategoryId"`                                                                //多对一 关联表
	Tags     []Tags   `gorm:"many2many:article_tags;foreignKey:ArticleId;joinForeignKey:ArticleId;References:TagId;joinReferences:TagId"` //多对多 关联表

}

type Category struct {
	CategoryId   int    `gorm:"primaryKey;column:category_id;autoIncrement"`
	CategoryName string `gorm:"column:category_name;type:varchar(255)"`
}

type Tags struct {
	TagId   int64  `gorm:"primaryKey;column:tag_id;autoIncrement"`
	TagName string `gorm:"column:tag_name;type:varchar(255)"`
}

type ArticleTags struct { //多对多
	ArticleId int   `gorm:"primaryKey;column:article_id"`
	TagId     int64 `gorm:"primaryKey;column:tag_id"`
}

type UserInfo struct {
	PhoneNumber string     `json:"phone_number" gorm:"PrimaryKey;column:phone_number;type:varchar(20);not null"`
	UserName    string     `json:"user_name" gorm:"column:user_name;type:varchar(50);not null"`
	PassWord    string     `json:"password" gorm:"column:password;type:varchar(100);not null"`
	Comments    []Comments `gorm:"foreignKey:PhoneNumber;references:PhoneNumber"` // 一对多关系
}

func (UserInfo) TableName() string {
	return "user_info"
}

type Comments struct {
	CommentId   int       `gorm:"primaryKey;column:comment_id;autoIncrement"`
	Content     string    `gorm:"column:content;type:text;not null"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime;not null;"`
	PhoneNumber string    `gorm:"column:phone_number;type:varchar(20);not null;"`
}
