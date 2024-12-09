# Blog  数据库

## 初始化创建4张表



```sql
-- 创建分类表
CREATE TABLE categories (
    category_id INT PRIMARY KEY AUTO_INCREMENT,
    category_name VARCHAR(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建标签表
CREATE TABLE tags (
    tag_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    tag_name VARCHAR(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建文章表
CREATE TABLE articles (
    article_id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    excerpt VARCHAR(500),
    content TEXT,
    category_id INT NOT NULL,
    view_count INT DEFAULT 0,
    create_time DATETIME NOT NULL,
    update_time DATETIME NOT NULL,
    status TINYINT DEFAULT 1,
    url VARCHAR(255),
    FOREIGN KEY (category_id) REFERENCES categories(category_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建文章标签关联表
CREATE TABLE article_tags (
    article_id INT,
    tag_id BIGINT,
    PRIMARY KEY (article_id, tag_id),
    FOREIGN KEY (article_id) REFERENCES articles(article_id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(tag_id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建用户信息表
CREATE TABLE user_info (
    phone_number VARCHAR(20) PRIMARY KEY,
    user_name VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建评论表
CREATE TABLE comments (
    comment_id INT PRIMARY KEY AUTO_INCREMENT,
    content TEXT NOT NULL,
    create_time DATETIME NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    FOREIGN KEY (phone_number) REFERENCES user_info(phone_number) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 添加索引优化查询性能
ALTER TABLE articles ADD INDEX idx_category (category_id);
ALTER TABLE articles ADD INDEX idx_create_time (create_time);
ALTER TABLE articles ADD INDEX idx_url (url);
ALTER TABLE tags ADD UNIQUE INDEX idx_tag_name (tag_name);
ALTER TABLE categories ADD UNIQUE INDEX idx_category_name (category_name);
ALTER TABLE comments ADD INDEX idx_phone_number (phone_number);


-- 插入测试分类
INSERT INTO categories (category_name) VALUES 
('network'),
('algorithm'),
('architecture'),
('os'),
('linux'),
('database'),
('golang'),
('leetcode'),
('project');

```







## go后端的 type 定义



```go
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
	CreateTime  time.Time `gorm:"column:ccreate_time;type:datetime;not null;"`
	PhoneNumber string    `gorm:"column:phone_number;type:varchar(20);not null;"`
}

```













