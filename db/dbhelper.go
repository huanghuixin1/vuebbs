package db

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)

type dbhelper struct{}

var DbHelper = dbhelper{}

//打开数据库
func openDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/note_online?charset=utf8mb4")
	if err != nil {
		fmt.Println("数据库打开出错", err)
	}
	return db
}

//
///**
//执行查询 返回受影响行数
// */
//func (this *dbhelper)Excute(sql string, args ...interface{}) (int64, error) {
//	db := openDb()
//	defer db.Close()
//	ret, err := db.Exec(sql, args...)
//	affected, _ := ret.RowsAffected()
//	return affected, err
//}
//
///**
//执行查询 返回结果集
// */
//func (this *dbhelper)Query(sql string, args ...interface{}) (*sql.Rows, error) {
//	db := openDb()
//	defer db.Close()
//
//	rows, error := db.Query(sql, args...)
//	return rows, error
//}
//
//func (this *dbhelper)QueryRow(sql string, args ...interface{}) (*sql.Row) {
//	db := openDb()
//	defer db.Close()
//
//	return db.QueryRow(sql, args...)
//}
//
////一堆的测试++++++++++++++++++++++++++++++++++++++
//
////测试实体
//type article struct {
//	id      int
//	title   string
//	summary string
//}
//
////测试插入
//func TestExcute(title string, summary string, content string, cid int) {
//	createTime := time.Now().Unix()
//
//	sql := "insert into Articles(ATitle,Summary,AContent,Fk_Cid,CreateTime) values(?,?,?,?,?)"
//
//	rows, error := DbHelper.Excute(sql, title, summary, content, cid, createTime)
//
//	if error != nil {
//		fmt.Println(error)
//	}
//	fmt.Println(rows)
//}
//
//func TestQuery() {
//	rows, err := DbHelper.Query("select id,ATitle,Summary from Articles limit 0 ,3")
//	if err != nil {
//		fmt.Println("查询错误", err)
//	}
//
//	var arrArticles [5]article
//
//	i := 0
//	for rows.Next() {
//		curra := article{}
//		rows.Scan(&curra.id, &curra.title, &curra.summary)
//		arrArticles[i] = curra
//		i++
//	}
//	defer rows.Close()
//
//	ret := arrArticles[:i]
//
//	fmt.Println(ret, len(ret))
//}