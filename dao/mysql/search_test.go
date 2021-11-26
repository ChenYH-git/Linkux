package mysql

import (
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
}

func TestGetPostListByIDsAndSearch(t *testing.T) {
	ids := []string{"1", "94354211461926912", "123"}
	p := &models.ParamPostList{
		Search: "linux",
	}

	posts, err := GetPostListByIDsAndSearch(ids, p)
	if err != nil {
		t.Fatalf("GetPostListByIDsAndSearch failed, err:%v\n", err)
	}
	t.Logf("GetPostListByIDsAndSearch suceess\nres: %v\n", posts[0])
}
