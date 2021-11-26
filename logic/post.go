package logic

import (
	"Linkux/dao/mysql"
	"Linkux/dao/redis"
	"Linkux/models"
	"Linkux/pkg/snowflakes"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	p.PostID = snowflakes.GenID()

	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreateRedisPost(p.PostID, p.LabelID)
	if err != nil {
		return err
	}

	err = mysql.AddContribution(p)
	if err != nil {
		return err
	}
	return
}

func GetPostByID(pid int64) (data *models.ApiPostDetail, err error) {
	post, err := mysql.GetPostByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(pid) failed", zap.Int64("pid", pid), zap.Error(err))
		return
	}
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID(post.AuthorID) failed",
			zap.String("author_id", post.AuthorID),
			zap.Error(err))
		return
	}

	label, err := mysql.GetLabelDetailByID(post.LabelID)
	if err != nil {
		zap.L().Error("mysql.GetLabelDetailByID(post.LabelID)",
			zap.Int64("label_id", post.LabelID),
			zap.Error(err))
		return
	}

	data = &models.ApiPostDetail{
		AuthorName:  user.Username,
		Post:        post,
		LabelDetail: label,
	}
	return
}

func GetNoLabelList(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder return 0 id")
		return
	}

	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}
	for idx, post := range posts {

		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed",
				zap.String("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}

		label, err := mysql.GetLabelDetailByID(post.LabelID)
		if err != nil {
			zap.L().Error("mysql.GetLabelDetailByID() failed",
				zap.Int64("label_id", post.LabelID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:  user.Username,
			PicLink:     user.PicLink,
			VoteNum:     voteData[idx],
			Post:        post,
			LabelDetail: label,
		}
		data = append(data, postDetail)
	}
	return
}

func GetLabelList(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	ids, err := redis.GetLabelPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetLabelPostIDsInOrder(p) return 0 id")
		return
	}

	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}
	for idx, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed",
				zap.String("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}

		label, err := mysql.GetLabelDetailByID(post.LabelID)
		if err != nil {
			zap.L().Error("mysql.GetLabelDetailByID() failed",
				zap.Int64("label_id", post.LabelID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:  user.Username,
			PicLink:     user.PicLink,
			VoteNum:     voteData[idx],
			Post:        post,
			LabelDetail: label,
		}
		data = append(data, postDetail)
	}
	return
}

func GetPostList(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	if p.LabelID == 0 {
		data, err = GetNoLabelList(p)
	} else {
		data, err = GetLabelList(p)
	}
	if err != nil {
		zap.L().Error("GetPostList failed, err: ", zap.Error(err))
		return nil, err
	}
	return
}
