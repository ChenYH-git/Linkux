package redis

import (
	"Linkux/models"
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

func TestCreateRedisPost(t *testing.T) {
	pid := int64(123678)
	label := int64(1)

	err := CreateRedisPost(pid, label)
	if err != nil {
		t.Fatalf("CreateRedisPost insert record into mysql failed, err:%v\n", err)
	}
	t.Logf("CreateRedisPost insert record into mysql suceess\n")
}

func TestGetPostIDsInOrder(t *testing.T) {
	p := models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: "score",
	}

	data, err := GetPostIDsInOrder(&p)
	if err != nil {
		t.Fatalf("GetPostIDsInOrder failed, err:%v\n", err)
	}
	t.Logf("GetPostIDsInOrder suceess\ndata:%v\n", data[0])
}

func TestGetLabelPostIDsInOrder(t *testing.T) {
	p := models.ParamPostList{
		Page:    1,
		Size:    10,
		LabelID: 1,
		Order:   "score",
	}

	data, err := GetLabelPostIDsInOrder(&p)
	if err != nil {
		t.Fatalf("GetLabelPostIDsInOrder failed, err:%v\n", err)
	}
	t.Logf("GetLabelPostIDsInOrder suceess\ndata:%v\n", data[0])
}

func TestGetPostVoteData(t *testing.T) {
	ids := []string{"94353058070269952"}

	data, err := GetPostVoteData(ids)
	if err != nil {
		t.Fatalf("GetPostVoteData failed, err:%v\n", err)
	}
	t.Logf("GetPostVoteData suceess\ndata:%v\n", data[0])
}

func TestGetPostVoteDataSingle(t *testing.T) {
	ids := "94353058070269952"

	data, err := GetPostVoteDataSingle(ids)
	if err != nil {
		t.Fatalf("GetPostVoteDataSingle failed, err:%v\n", err)
	}
	t.Logf("GetPostVoteDataSingle suceess\ndata:%v\n", data)
}
