package categories

type CategoriesName struct {
	Id           int                                 //id
	CName        string `gorm:"column:CName"`        //板块名字
	ArticleCount string `gorm:"column:ArticleCount"` //板块帖子数量
}
