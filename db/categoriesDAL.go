package db

import (
	"../entity"
)

type categoriesDAL struct {

}

var CategoriesDAL = categoriesDAL{}

//通过id获取板块名字
func (this *categoriesDAL) GetNameById(id int) string {
	db := openDb()
	defer db.Close()

	var category entity.Categories
	db.Table("categories").Where("id=?", id).Select("CName").Find(&category)
	return category.CName;
}