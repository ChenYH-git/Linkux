package logic

import (
	"Linkux/dao/mysql"
	"Linkux/dao/redis"
	"Linkux/models"
	"strconv"

	"go.uber.org/zap"
)

func VoteForPost(userID string, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.String("userID", userID),
		zap.Int64("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(userID, p.PostID, float64(p.Direction))
}

func GetVoteCollect(p *models.Trigger, userID string) (vflag, cflag bool, err error) {
	vflag, err = redis.GetMyVote(strconv.FormatInt(p.PostID, 10), userID)
	if err != nil {
		zap.L().Error("GetMyVote",
			zap.String("userID", userID),
			zap.Int64("postID", p.PostID),
			zap.Error(err))
		return false, false, err
	}

	cflag, err = mysql.CheckCollect(strconv.FormatInt(p.PostID, 10), userID)
	if err != nil {
		zap.L().Error("CheckCollect",
			zap.String("userID", userID),
			zap.Int64("postID", p.PostID),
			zap.Error(err))
		return false, false, err
	}
	return
}
