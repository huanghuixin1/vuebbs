package test

import (
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
	"encoding/json"
)


//查询测试
func testQuery() {
	u := &UserInfo{}
	u.Id = 1
	db, _ := openDb()
	defer db.Close()
	db.First(u)

	buffer, _ := json.Marshal(u)
	fmt.Println(string(buffer))
}

//带条件的查询
func testQueryWidthWhere() {
	db, _ := openDb()
	defer db.Close()
	articles := &[]Article{}
	db.Where("fk_userid=? and IsDelete=?", 1, 0).Order("id desc").Limit(3).Find(articles)

	fmt.Println(articles)
}

//链接查询
func testQueryJoin() {
	db, _ := openDb()
	defer db.Close()

	articleItems := &[]ArticleItem{}
	db.Table("articles").Select(`articles.id as id,
		articles.ATitle as Title,
			articles.Summary as Summary,
			articles.Fk_Cid,
			articles.Fk_userId,
			articles.CreateTime,userinfos.NickName`).Joins("right join userinfos on articles.fk_userid=userinfos.id").Where("userinfos.IsDelete=0 and articles.IsDelete=0").Find(articleItems)

	fmt.Println(articleItems)
}

//直接执行sql语句查询
func querySql() {
	sql := `SELECT
				articles.id,
				articles.ATitle,
				articles.Summary,
				articles.Fk_Cid,
				articles.Fk_userId,
				articles.CreateTime,
				u.nickName
			FROM
				articles
			RIGHT JOIN userinfos u ON articles.Fk_userId = u.id
			WHERE
				u.IsDelete = 0
			AND articles.IsDelete = 0
			ORDER BY
				articles.CreateTime DESC
			LIMIT ?,?`
	db,_ := openDb()
	defer db.Close()
	rows,err := db.Raw(sql,0,10).Rows()
	if err != nil{
		fmt.Println("查询出错:", err)
	}

	articlesArr := make([]ArticleItem,10)
	i:=0
	for rows.Next(){
		articles := &ArticleItem{}
		db.ScanRows(rows, articles)
		articlesArr[i] = *articles
		i++
	}

	ret := articlesArr[:i]
	fmt.Println(len(ret), ret)
}

func testAdd() {
	u := UserInfo{}

	u.NickName = "hhx223";
	u.LoginName = "hhx223";
	u.LoginPwd = "123";
	u.EmailAddr = "123@23.123";
	u.IsDelete = false;
	u.Level = 999;
	u.CreateTime = int(time.Now().Unix());

	db, _ := openDb()
	defer db.Close()

	db.Save(&u)
}

func openDb() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:root123@tcp(localhost:3306)/note_online?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("打开数据库出错:", err)
	}

	return db, err
}

type UserInfo struct {
	//gorm.Model
	BaseEntity
	NickName  string `gorm:"column:NickName"`
	LoginName string  `gorm:"column:LoginName"`
	LoginPwd  string `gorm:"column:LoginPwd"`
	EmailAddr string `gorm:"column:EmailAddr"`
	Level     int `gorm:"column:ULevel"`
}

func (UserInfo) TableName() string {
	return "userinfos"
}

type Article struct {
	BaseEntity
	Title     string `gorm:"column:ATitle"`  //标题
	Summary   string `gorm:"column:Summary"` //摘要
	Content   string `gorm:"column:Content"` //内容
	Fk_Cid    int  `gorm:"column:Fk_Cid"`    //模块id
	Fk_UserId int  `gorm:"column:Fk_UserId"`
}

func (this *Article) Article() string {
	return "articles"
}

//鸡肋entity
type BaseEntity struct {
	Id         int `gorm:"primary_key;column:id"`
	CreateTime int `gorm:"column:CreateTime"`
	IsDelete   bool `gorm:"column:IsDelete"`
}

type ArticleItem struct {
	Id         int    //id
	Title      string `gorm:"column:Title"` //标题
	Summary    string `gorm:"column:Summary"` //摘要
	Cid        int    `gorm:"column:Fk_Cid"` //类别id
	CreateTime int    `gorm:"column:CreateTime"` //发布时间
	UserId     int    `gorm:"column:Fk_userId"` //发布者id
	NickName   string `gorm:"column:NickName"` //发布者昵称
}




