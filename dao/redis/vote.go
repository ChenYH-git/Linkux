package redis

import (
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekSeconds = 7 * 24 * 3600
	scorePerVote   = 432 // 一个赞多少分
)

var (
	errVoteTimeExpire = errors.New("投票时间已过")
	errVoteRepeated   = errors.New("不允许重复投相同票")
)

func VoteForPost(userID string, postID int64, value float64) error {
	postTime := client.ZScore(getRedisKey(KeyPostTimeZset), strconv.FormatInt(postID, 10)).Val()
	if float64(time.Now().Unix())-postTime > oneWeekSeconds {
		return errVoteTimeExpire
	}

	ov := client.ZScore(getRedisKey(KeyPostVotedZsetPF+strconv.FormatInt(postID, 10)), userID).Val()
	if value == ov {
		return errVoteRepeated
	}
	var dir float64
	if value > ov {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ov - value)
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZset), dir*diff*scorePerVote, strconv.FormatInt(postID, 10))

	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZsetPF+strconv.FormatInt(postID, 10)), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZsetPF+strconv.FormatInt(postID, 10)), redis.Z{
			Score:  value,
			Member: userID,
		})
	}

	_, err := pipeline.Exec()
	return err
}
