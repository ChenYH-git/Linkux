package mysql

import (
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

func TestGetUserRank(t *testing.T) {
	data, err := GetUserRank()
	if err != nil {
		t.Fatalf("GetUserRank failed, err:%v\n", err)
	}
	t.Logf("GetUserRank suceess\ndata:%v\n", data[0])
}

func TestGetMyRank(t *testing.T) {
	id := "1"

	data, err := GetMyRank(id)
	if err != nil {
		t.Fatalf("GetMyRank failed, err:%v\n", err)
	}
	t.Logf("GetMyRank suceess\ndata:%v\n", data[0])
}
