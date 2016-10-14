package db

import (
	"../entity"
	"../dto/categories"
	"fmt"
)

type categoriesDAL struct {

}

const categoryTableName = "categories"

var CategoriesDAL = categoriesDAL{}

//通过id获取板块名字
func (this *categoriesDAL) GetNameById(id int) string {
	db := openDb()
	defer db.Close()

	var category entity.Categories
	db.Table(categoryTableName).Where("id=?", id).Select("CName").Find(&category)
	return category.CName;
}

func (this *categoriesDAL) GetListName() []categories.CategoriesName {
	db := openDb()
	defer db.Close()

	sql := `SELECT
				c.id,
				c.CName,
				count(a.id) as ArticleCount
			FROM
				categories c
			LEFT JOIN (select id,Fk_Cid from articles where IsDelete = 0) a ON a.Fk_Cid = c.id
			WHERE
				c.IsDelete = 0
			GROUP BY
				c.id
			order by c.id DESC `
	rows, _ := db.Raw(sql).Rows()
	defer rows.Close()

	listRet := make([]categories.CategoriesName,0,20)
	for rows.Next(){
		curr := categories.CategoriesName{}
		err := db.ScanRows(rows, &curr)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		listRet = append(listRet,curr)
	}

	return listRet;
}
