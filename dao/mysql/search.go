package mysql

import (
	"Linkux/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

func GetPostListByIDsAndSearch(ids []string, p *models.ParamPostList) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, label_id, collect_num, viewd_num, create_time, status
	from post
	where post_id in (?)
	and status = 1
	and (title like ? or content like ?)
	order by FIND_IN_SET(post_id, ?)
	`

	query, args, err := sqlx.In(sqlStr, ids, "%"+p.Search+"%", "%"+p.Search+"%", strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)

	err = db.Select(&postList, query, args...)
	if err != nil {
		return nil, err
	}
	sqlStr = `select count(post_id) from vpost where post_id = ?`

	for i, _ := range postList {
		var count int
		err = db.Get(&count, sqlStr, postList[i].PostID)
		if err != nil {
			return nil, err
		}
		if count < 1 {
			postList[i].Qualified = false
			continue
		}
		postList[i].Qualified = true
	}
	return
}
