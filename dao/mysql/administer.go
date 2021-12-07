package mysql

import (
	"Linkux/models"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

func GetAllUser(p *models.ParamPostList) (data []*models.User, err error) {
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1

	sqlStr := `select
	username, user_id, contribution, pic_link
	from user
	order by contribution
	desc
	limit ?,?
	`
	data = make([]*models.User, 0, 10)
	err = db.Select(&data, sqlStr, start, end)
	if err != nil {
		return
	}

	ids := make([]string, 0, len(data))
	for _, v := range data {
		ids = append(ids, v.UserID)
	}

	sqlStr = `select count(user_id) from vuser where user_id = ?`

	for i, _ := range data {
		var count int
		err = db.Get(&count, sqlStr, data[i].UserID)
		if err != nil {
			return nil, err
		}
		if count == 0 {
			data[i].Qualified = false
			continue
		}
		data[i].Qualified = true
	}
	return
}

func StarPosts(p *models.Trigger) (err error) {
	sqlStr := `select count(post_id) from vpost where post_id = ?`
	var count int
	if err := db.Get(&count, sqlStr, p.PostID); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("帖子已加精")
	}

	sqlStr = `insert into vpost(post_id) values(?)`

	_, err = db.Exec(sqlStr, p.PostID)
	return err
}

func CancelStarPosts(p *models.Trigger) (err error) {
	sqlStr := `select count(post_id) from vpost where post_id = ?`
	var count int
	if err := db.Get(&count, sqlStr, p.PostID); err != nil {
		return err
	}
	if count == 0 {
		return errors.New("帖子未加精，无法取消")
	}

	sqlStr = `delete from vpost where post_id = ?`

	_, err = db.Exec(sqlStr, p.PostID)
	return err
}

func DeletePosts(p *models.Trigger) (err error) {
	sqlStr := `delete from post where post_id = ?`
	_, err = db.Exec(sqlStr, p.PostID)
	sqlStr = `delete from vpost where post_id = ?`
	_, err = db.Exec(sqlStr, p.PostID)
	return err
}

func CheckPost(p *models.Trigger) (qualified bool, err error) {
	sqlStr := `select count(post_id) from vpost where post_id = ?`
	var count int
	if err := db.Get(&count, sqlStr, p.PostID); err != nil {
		return false, err
	}
	if count > 0 {
		qualified = true
	} else {
		qualified = false
	}
	return
}

func StarUser(p *models.StarUser) (err error) {
	sqlStr := `select count(user_id) from vuser where user_id = ?`
	var count int
	if err := db.Get(&count, sqlStr, p.UserID); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已加v")
	}
	sqlStr = `insert into vuser(user_id) values(?)`

	_, err = db.Exec(sqlStr, p.UserID)
	return err
}

func CancelStarUser(p *models.StarUser) (err error) {
	sqlStr := `select count(user_id) from vuser where user_id = ?`
	var count int
	if err := db.Get(&count, sqlStr, p.UserID); err != nil {
		return err
	}
	if count == 0 {
		return errors.New("用户未加v，无法取消")
	}

	sqlStr = `delete from vuser where user_id = ?`

	_, err = db.Exec(sqlStr, p.UserID)
	return err
}

func CheckUser(p *models.StarUser) (qualified bool, err error) {
	sqlStr := `select count(user_id) from vuser where user_id = ?`
	var count int
	if err := db.Get(&count, sqlStr, p.UserID); err != nil {
		return false, err
	}
	if count > 0 {
		qualified = true
	} else {
		qualified = false
	}
	return
}

func GetWaitingTransTask(p *models.ParamPostList) (data []*models.Trans, err error) {
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1

	sqlStr := `select trans_id, title, content, create_time from trans where status = 0 limit ?,?`
	if err = db.Select(&data, sqlStr, start, end); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("No trans_task waiting\n")
			err = nil
		}
	}
	return
}

func GetWaitingPostsTask(p *models.ParamPostList) (data []*models.Post, err error) {
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1

	sqlStr := `select post_id, label_id, title, content, create_time from post where status = 0 limit ?,?`
	if err = db.Select(&data, sqlStr, start, end); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("No posts waiting\n")
			err = nil
		}
	}

	return
}

func GetAuthorByPostID(pid int64) (AuthorID string, err error) {
	sqlStr := `select author_id from post where post_id = ?`
	err = db.Get(&AuthorID, sqlStr, pid)
	return
}

func JudgePass(p *models.Judge) (err error) {
	if p.TransID == 0 {
		sqlStr := `update post set status = 1 where post_id = ?`
		_, err = db.Exec(sqlStr, p.PostID)
		return
	}

	sqlStr := `update trans set status = 1 where trans_id = ?`
	_, err = db.Exec(sqlStr, p.TransID)
	return
}

func DeleteTrans(p *models.Task) (err error) {
	sqlStr := `delete from trans where trans_id = ?`
	_, err = db.Exec(sqlStr, p.TransID)
	return err
}
