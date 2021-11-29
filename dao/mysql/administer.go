package mysql

import (
	"Linkux/models"
	"errors"
)

func GetAllUser(p *models.ParamPostList) (data []*models.ApiRankDetail, err error) {
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1

	sqlStr := `select
	username, contribution, pic_link
	from user
	order by contribution
	desc
	limit ?,?
	`
	data = make([]*models.ApiRankDetail, 0, 10)
	err = db.Select(&data, sqlStr, start, end)
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
	if count < 1 {
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
	if count < 1 {
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
