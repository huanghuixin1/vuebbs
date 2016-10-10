package entity


//鸡肋entity
type BaseEntity struct {
	Id         int `gorm:"primary_key;column:id"`
	CreateTime int `gorm:"column:CreateTime"`
	IsDelete   bool `gorm:"column:IsDelete"`
}
