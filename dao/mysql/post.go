package mysql

import (
	"Linkux/models"
	"database/sql"
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
)

const conPerPost = 10

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
	post_id, title, content, author_id, label_id, trans_id)
	values(? , ?, ?, ?, ?, ?)
	`
	_, err = db.Exec(sqlStr, p.PostID, p.Title, p.Content, p.AuthorID, p.LabelID, p.TransID)
	return
}

func AddContribution(p *models.Post) (err error) {
	sqlStr := `update user set contribution = contribution + ? where user_id = ?`
	_, err = db.Exec(sqlStr, conPerPost, p.AuthorID)
	return
}

func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, label_id, collect_num, viewd_num, create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id, ?)
	`

	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)

	err = db.Select(&postList, query, args...)
	return
}

func GetLabelDetailByID(id int64) (label *models.LabelDetail, err error) {
	label = new(models.LabelDetail)
	sqlStr := `select
				label_id, label_name, introduction, create_time
				from label
				where label_id = ?
	`
	if err = db.Get(label, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("无效标签id")
		}
	}
	return label, err
}

func GetPostByID(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select
	post_id, title, content, author_id, label_id, collect_num, viewd_num, create_time
	from post
	where post_id = ?
	`
	err = db.Get(post, sqlStr, pid)
	return
}
