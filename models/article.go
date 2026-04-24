package models

type Article struct {
	Id     int
	Title  string
	CateId int
	State  int
	// ArticleCateA ArticleCate `gorm:"foreignKey:CateId;references:Id"`
}

// 表示配置操作数据库的表名称
func (Article) TableName() string {
	return "article"
}
