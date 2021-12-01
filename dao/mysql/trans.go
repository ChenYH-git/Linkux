package mysql

import (
	"Linkux/dao/redis"
	"Linkux/models"
	"strconv"

	"go.uber.org/zap"
)

func CreateTrans(p *models.Trans) (err error) {
	sqlStr := `insert into trans(
	trans_id, title, content)
	values(? , ?, ?)
	`
	_, err = db.Exec(sqlStr, p.TransID, p.Title, p.Content)
	return
}

func GetTransTask() (data []*models.Trans, err error) {
	sqlStr := `select trans_id, title, content, create_time from trans where status = 1`

	err = db.Select(&data, sqlStr)
	return
}

func GetTransExist(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1

	sqlStr := `select post_id, trans_id, title, content, author_id, label_id, collect_num, viewd_num, create_time, status
	from post
	where trans_id = ?
	and status = 1
	limit ?,?
	`
	err = db.Select(&data, sqlStr, p.TransID, start, end)
	if err != nil {
		return nil, err
	}

	sqlStr = `select count(post_id) from vpost where post_id = ?`

	for i, _ := range data {
		var count int
		err = db.Get(&count, sqlStr, data[i].PostID)
		if err != nil {
			return nil, err
		}
		if count < 1 {
			data[i].Qualified = false
			continue
		}
		data[i].Qualified = true
	}

	for idx, post := range data {
		user, err := GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed",
				zap.String("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}

		label, err := GetLabelDetailByID(post.LabelID)
		if err != nil {
			zap.L().Error("mysql.GetLabelDetailByID() failed",
				zap.Int64("label_id", post.LabelID),
				zap.Error(err))
			continue
		}
		voteData, err := redis.GetPostVoteDataSingle(strconv.FormatInt(post.PostID, 10))
		if err != nil {
			return nil, err
		}
		data[idx].VoteNum = voteData
		data[idx].AuthorName = user.Username
		data[idx].PicLink = user.PicLink
		data[idx].AuthorQualified = user.Qualified
		data[idx].LabelDetail = label
	}
	return
}
