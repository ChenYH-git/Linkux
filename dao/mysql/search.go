package mysql

import (
	"Linkux/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

func GetPostListByIDsAndSearch(ids []string, p *models.ParamPostList) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, label_id, collect_num, viewd_num, create_time
	from post
	where post_id in (?)
	and title like ?
	order by FIND_IN_SET(post_id, ?)
	`

	query, args, err := sqlx.In(sqlStr, ids, "%"+p.Search+"%", strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)

	err = db.Select(&postList, query, args...)
	return
}
