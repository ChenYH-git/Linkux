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

func TestCreatePost(t *testing.T) {
	post := models.Post{
		PostID:   123456,
		LabelID:  1,
		AuthorID: "123",
		Title:    "TestCreatePost",
		Content:  "TestCreatePost",
	}

	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("CreatePost insert record into mysql failed, err:%v\n", err)
	}
	t.Logf("CreatePost insert record into mysql suceess\n")
}

func TestAddContribution(t *testing.T) {
	post := models.Post{
		AuthorID: "5",
	}

	err := AddContribution(&post)
	if err != nil {
		t.Fatalf("AddContribution insert record into mysql failed, err:%v\n", err)
	}
	t.Logf("AddContribution insert record into mysql suceess\n")
}

func TestGetLabelDetailByID(t *testing.T) {
	id := int64(1)

	label, err := GetLabelDetailByID(id)
	if err != nil {
		t.Fatalf("GetLabelDetailByID failed, err:%v\n", err)
	}
	t.Logf("GetLabelDetailByID suceess\nlabel:%v\n", label)
}

func TestGetPostListByIDs(t *testing.T) {
	ids := []string{"1", "2", "123"}

	posts, err := GetPostListByIDs(ids)
	if err != nil {
		t.Fatalf("GetPostListByIDs failed, err:%v\n", err)
	}
	t.Logf("GetPostListByIDs suceess\nposts: %v\n", posts[0])
}

func TestGetPostByID(t *testing.T) {
	pid := int64(1111111)

	post, err := GetPostByID(pid)
	if err != nil {
		t.Fatalf("GetPostByID failed, err:%v\n", err)
	}
	t.Logf("GetPostByID suceess\npost:%v\n", post)
}
