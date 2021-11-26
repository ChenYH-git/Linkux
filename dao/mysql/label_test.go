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

func TestGetLabelList(t *testing.T) {
	labels, err := GetLabelList()
	if err != nil {
		t.Fatalf("GetLabelList failed, err:%v\n", err)
	}
	t.Logf("GetLabelList suceess\npost:%v\n", labels[0])
}
