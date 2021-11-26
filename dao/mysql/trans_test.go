package mysql

import (
	"Linkux/dao/redis"
	"Linkux/models"
	"Linkux/settings"
	"testing"
)

func init() {
	dbCfg := settings.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "chenyuhan123000",
		DbName:       "linkux",
		Port:         3306,
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	err := Init(&dbCfg)
	if err != nil {
		panic(err)
	}

	redisCfg := settings.RedisConfig{
		Host:     "127.0.0.1",
		Password: "",
		Port:     6379,
		DB:       0,
		PoolSize: 100,
	}

	err = redis.Init(&redisCfg)
	if err != nil {
		panic(err)
	}
}

func TestCreateTrans(t *testing.T) {
	p := &models.Trans{
		TransID: 3,
		Title:   "title",
		Content: "content",
	}

	err := CreateTrans(p)
	if err != nil {
		t.Fatalf("CreateTrans insert record into mysql failed, err:%v\n", err)
	}
	t.Logf("CreateTrans insert record into mysql suceess\n")
}

func TestGetTransTask(t *testing.T) {
	data, err := GetTransTask()
	if err != nil {
		t.Fatalf("GetTransTask failed, err:%v\n", err)
	}
	t.Logf("GetTransTask suceess\npost:%v\n", data[0])
}

func TestGetTransExist(t *testing.T) {
	p := models.ParamPostList{
		LabelID: 1,
		Page:    1,
		Size:    10,
		TransID: 93982895189790720,
		Order:   "score",
	}

	data, err := GetTransExist(&p)
	if err != nil {
		t.Fatalf("GetTransExist failed, err:%v\n", err)
	}
	t.Logf("GetTransExist suceess\npost:%v\n", data[0])
}
