package mysql

import (
	"Linkux/models"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

func InsertUser(user *models.User) (err error) {
	sqlStr := `select count(user_id) from user where user_id = ?`
	var count int
	if err := db.Get(&count, sqlStr, user.UserID); err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	sqlStr = `insert into user(user_id, contribution) values(?,?)`

	_, err = db.Exec(sqlStr, user.UserID, user.Contribution)
	return
}

func GetUserByID(uid string) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select username,pic_link from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	return
}

func GetPostListOfMy(ids []string, userID string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, label_id, create_time
	from post
	where post_id in (?)
	and author_id = ?
	order by FIND_IN_SET(post_id, ?)
	`

	query, args, err := sqlx.In(sqlStr, ids, userID, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)

	err = db.Select(&postList, query, args...)
	return

}

func AddCollection(p *models.Triger, userID string) (err error) {
	sqlStr := `insert into collection(user_id, post_id) values(?, ?)`
	_, err = db.Exec(sqlStr, userID, p.PostID)
	return
}

func AddCollectNum(p *models.Triger) (err error) {
	sqlStr := `update post set collect_num = collect_num + 1 where post_id = ?`
	_, err = db.Exec(sqlStr, p.PostID)
	return
}

func DeleteCollection(p *models.Triger, userID string) (err error) {
	sqlStr := `delete from collection where user_id = ? and post_id = ?`
	_, err = db.Exec(sqlStr, userID, p.PostID)
	return
}

func DeleteCollectNum(p *models.Triger) (err error) {
	sqlStr := `update post set collect_num = collect_num - 1 where post_id = ?`
	_, err = db.Exec(sqlStr, p.PostID)
	return
}

func GetCollectionIDs(p *models.ParamPostList, userID string) (ids []string, err error) {
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	sqlStr := `select post_id
	from collection
	where user_id = ?
	limit ?, ?`
	var postList []*models.Post
	err = db.Select(&postList, sqlStr, userID, start, end)
	ids = make([]string, 0, len(postList))
	for _, v := range postList {
		ids = append(ids, strconv.FormatInt(v.ID, 10))
	}
	return
}

func AddViewNum(p *models.Triger) (err error) {
	sqlStr := `update post set viewd_num = viewd_num + 1 where post_id = ?`
	_, err = db.Exec(sqlStr, p.PostID)
	return
}
