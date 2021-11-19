package logic

import (
	"Linkux/dao/mysql"
	"Linkux/dao/redis"
	"Linkux/models"
	"strconv"

	"go.uber.org/zap"
)

func GetSearchRes(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder return 0 id")
		return
	}

	posts, err := mysql.GetPostListByIDsAndSearch(ids, p)
	if err != nil {
		return
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
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
		voteData, err := redis.GetPostVoteDataSingle(strconv.FormatInt(post.PostID, 10))
		if err != nil {
			return nil, err
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:  user.Username,
			PicLink:     user.PicLink,
			VoteNum:     voteData,
			Post:        post,
			LabelDetail: label,
		}
		data = append(data, postDetail)
	}
	return
}
