package logic

import (
	"Linkux/dao/redis"
	"Linkux/models"

	"go.uber.org/zap"
)

func VoteForPost(userID string, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.String("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(userID, p.PostID, float64(p.Direction))
}
