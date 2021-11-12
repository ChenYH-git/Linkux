package logic

import (
	"Linkux/dao/mysql"
	"Linkux/dao/redis"
	"Linkux/models"
	"Linkux/pkg/snowflakes"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	p.ID = snowflakes.GenID()

	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreateRedisPost(p.ID, p.LabelID)
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

		label, err :=
	}
}

func GetLabelList(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {

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
