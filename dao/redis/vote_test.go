package redis

import (
	"Linkux/settings"
	"testing"
)

func init() {
	redisCfg := settings.RedisConfig{
		Host:     "127.0.0.1",
		Password: "",
		Port:     6379,
		DB:       0,
		PoolSize: 100,
	}

	err := Init(&redisCfg)
	if err != nil {
		panic(err)
	}
}

func TestVoteForPost(t *testing.T) {
	uid := "123"
	pid := int64(123)
	value := float64(1)

	err := VoteForPost(uid, pid, value)
	if err != nil {
		t.Fatalf("VoteForPost insert record into mysql failed, err:%v\n", err)
	}
	t.Logf("VoteForPost insert record into mysql suceess\n")
}
