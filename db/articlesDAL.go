package db

import (
	"../dto/articles"
	"strings"
	"strconv"
	"fmt"
	"math"
)

type articlesDAL struct {

}

var ArticlesDAL = articlesDAL{}

/**
	根据类别id获取帖子列表

	@param cid 模块id 如果是0 表示推荐内容(仅按照时间排序)
	@param offset 偏移值
	@param size 获取数量
	@param minId 最小id 取到的值要大于该id
	@param maxId 最大id 取到的值要小于该id
 */
func (this *articlesDAL)GetListByCategoryId(cid int, size int, minId int, maxId int) ([]dto.ArticleItem, error) {
	db := openDb()
	defer db.Close()
	sql := `SELECT
				articles.id,
				articles.ATitle,
				articles.Summary,
				articles.Fk_Cid,
				articles.Fk_userId,
				articles.CreateTime,
				u.NickName
			FROM
				articles
			RIGHT JOIN userinfos u ON articles.Fk_userId = u.id
			WHERE
				u.IsDelete = 0
			AND articles.IsDelete = 0
			AND articles.Id > ?
			AND articles.Id < ?
			{0}
			ORDER BY
				articles.Id DESC
			LIMIT 0,?`

	//cid的where条件语句
	//2147483647

	if maxId <= 0 {
		maxId = math.MaxInt32
	}

	cidWhere := ""
	//如果cid为0推荐内容 不筛选板块
	if (cid != 0) {
		cidWhere = "AND articles.Fk_Cid = " + strconv.Itoa(cid)
	}

	//替换sql中的cidwhere条件
	sql = strings.Replace(sql, "{0}", cidWhere, 1)

	rows, err := db.Raw(sql, minId, maxId, size).Rows()
	if err != nil {
		fmt.Println("获取失败", err)
		return nil, err
	}

	//rows, err := DbHelper.Query(sql)
	defer rows.Close()

	temp_articleItem := dto.ArticleItem{}
	slice_articles := make([]dto.ArticleItem, size)
	length, errResu := temp_articleItem.Rows2Entites(db, rows, slice_articles, temp_articleItem)
	if errResu != nil {
		return nil, errResu
	}

	return slice_articles[:length], nil
}

/**
获取帖子总数
 */
func (this *articlesDAL)GetCountByCid(cid int) int {
	db := openDb()
	defer db.Close()

	var count int
	if cid == 0 {
		db.Table("articles").Where("IsDelete = 0").Count(&count)
	} else {
		db.Table("articles").Where("IsDelete = 0 AND Fk_cid=?", cid).Count(&count)
	}

	return count
}