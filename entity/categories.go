package entity

type Categories struct {
	BaseEntity
	CName string `gorm:"column:CName"`
}

func (Categories) TableName() string {
	return "categories"
}
