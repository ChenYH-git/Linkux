package redis

import (
	"Linkux/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func CreateRedisPost(postID, labelID int64) error {
	pipeline := client.TxPipeline()
	pipeline.ZAdd(getRedisKey(KeyPostTimeZset), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	pipeline.ZAdd(getRedisKey(KeyPostScoreZset), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	lKey := getRedisKey(KeyLabelSetPF + strconv.Itoa(int(labelID)))
	pipeline.SAdd(lKey, postID)
	_, err := pipeline.Exec()
	return err
}

func getIDsFromKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	return client.ZRevRange(key, start, end).Result()
}

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	key := getRedisKey(KeyPostTimeZset)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZset)
	}
	return getIDsFromKey(key, p.Page, p.Size)
}

func GetLabelPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZset)
	if orderKey == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZset)
	}

	lKey := getRedisKey(KeyLabelSetPF + strconv.Itoa(int(p.LabelID)))
	key := orderKey + strconv.Itoa(int(p.LabelID))
	if client.Exists(key).Val() < 1 {
		pipeline := client.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, lKey, orderKey)
		pipeline.Expire(key, 60*time.Second)
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	return getIDsFromKey(key, p.Page, p.Size)
}

func GetPostVoteData(ids []string) (data []int64, err error) {
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZsetPF + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(ids))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

func GetPostVoteDataSingle(id string) (data int64, err error) {
	pipeline := client.Pipeline()

	key := getRedisKey(KeyPostVotedZsetPF + id)
	pipeline.ZCount(key, "1", "1")

	cmder, err := pipeline.Exec()
	if err != nil {
		return 0, err
	}

	data = cmder[0].(*redis.IntCmd).Val()

	return
}
