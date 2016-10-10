package dto

import (
	"github.com/jinzhu/gorm"
	"database/sql"
)

type ArticleItem struct {
	Id         int                               //id
	Title      string `gorm:"column:ATitle"`      //标题
	Summary    string `gorm:"column:Summary"`    //摘要
	Cid        int    `gorm:"column:Fk_Cid"`     //类别id
	CreateTime int    `gorm:"column:CreateTime"` //发布时间
	UserId     int    `gorm:"column:Fk_userId"`  //发布者id
	NickName   string `gorm:"column:NickName"`   //发布者昵称
	HeaderImg string `gorm:"column:HeaderImg"`   //发布者头像
}

func (this *ArticleItem)Rows2Entites(db *gorm.DB, rows *sql.Rows, slice []ArticleItem, model ArticleItem) (int, error) {
	i := 0
	for rows.Next() {
		curr := model
		err := db.ScanRows(rows, &curr)
		if err != nil {
			return -1, err
		}
		slice[i] = curr
		i++
	}

	return i, nil
}