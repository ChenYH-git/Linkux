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
	sqlStr := `select trans_id, title, content, create_time from trans`

	err = db.Select(&data, sqlStr)
	return
}

func GetTransExist(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1

	sqlStr := `select post_id, trans_id, title, content, author_id, label_id, collect_num, viewd_num, create_time
	from post
	where trans_id = ?
	limit ?,?
	`
	err = db.Select(&data, sqlStr, p.TransID, start, end)

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
		data[idx].LabelDetail = label
	}
	return
}
