package entity

type Article struct {
	BaseEntity
	Title     string `gorm:"column:ATitle"`  //标题
	Summary   string `gorm:"column:Summary"` //摘要
	Content   string `gorm:"column:Content"` //内容
	Fk_Cid    int  `gorm:"column:Fk_Cid"`    //模块id
	Fk_UserId int  `gorm:"column:Fk_UserId"`
}

func (Article) TableName() string {
	return "articles"
}