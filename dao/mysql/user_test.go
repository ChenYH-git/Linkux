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

func TestInsertUser(t *testing.T) {
	user := models.User{
		UserID:   "12345",
		Username: "name",
		PicLink:  "pic",
	}

	err := InsertUser(&user)
	if err != nil {
		t.Fatalf("InsertUser insert record into mysql failed, err:%v\n", err)
	}
	t.Logf("InsertUser insert record into mysql suceess\n")
}

func TestGetUserByID(t *testing.T) {
	uid := "1234"

	data, err := GetUserByID(uid)
	if err != nil {
		t.Fatalf("GetUserByID failed, err:%v\n", err)
	}
	t.Logf("GetUserByID suceess\nuser:%v\n", data)
}

func TestGetFollowUserByIDs(t *testing.T) {
	ids := make([]*models.Follow, 0, 10)

	ids = append(ids, &models.Follow{
		FollowID: "o49Zd5yF1xxHGnHRl4Jgtb2ZuGaU",
	}, &models.Follow{
		FollowID: "o49Zd58gLLr2YjYgaRTfsLynkKVc",
	})
	data, err := GetFollowUserByIDs(ids)
	if err != nil {
		t.Fatalf("GetFollowUserByIDs failed, err:%v\n", err)
	}
	t.Logf("GetFollowUserByIDs suceess\nuser:%v\n", data[0])
}

func TestGetFollowedUserByIDs(t *testing.T) {
	ids := make([]*models.Follow, 0, 10)

	ids = append(ids, &models.Follow{
		FollowedID: "o49Zd5yF1xxHGnHRl4Jgtb2ZuGaU",
	}, &models.Follow{
		FollowedID: "o49Zd58gLLr2YjYgaRTfsLynkKVc",
	})
	data, err := GetFollowedUserByIDs(ids)
	if err != nil {
		t.Fatalf("GetFollowedUserByIDs failed, err:%v\n", err)
	}
	t.Logf("GetFollowedUserByIDs suceess\nuser:%v\n", data[0])
}

func TestGetPostListOfMy(t *testing.T) {
	ids := []string{"1", "2", "123"}
	userID := "123"

	posts, err := GetPostListOfMy(ids, userID)
	if err != nil {
		t.Fatalf("GetPostListOfMy failed, err:%v\n", err)
	}
	t.Logf("GetPostListOfMy suceess\nposts: %v\n", posts[0])
}

func TestAddCollection(t *testing.T) {
	uid := "123"
	p := models.Trigger{PostID: 123}
	err := AddCollection(&p, uid)
	if err != nil {
		t.Fatalf("AddCollection failed, err:%v\n", err)
	}
	t.Logf("AddCollection suceess\n")
}

func TestAddCollectNum(t *testing.T) {
	p := models.Trigger{PostID: 123}
	err := AddCollectNum(&p)
	if err != nil {
		t.Fatalf("AddCollectNum failed, err:%v\n", err)
	}
	t.Logf("AddCollectNum suceess\n")
}

func TestDeleteCollection(t *testing.T) {
	uid := "123"
	p := models.Trigger{PostID: 123}
	err := DeleteCollection(&p, uid)
	if err != nil {
		t.Fatalf("DeleteCollection failed, err:%v\n", err)
	}
	t.Logf("DeleteCollection suceess\n")
}

func TestGetCollectionIDs(t *testing.T) {
	uid := "1234"
	p := models.ParamPostList{
		Page: 1,
		Size: 10,
	}
	ids, err := GetCollectionIDs(&p, uid)
	if err != nil {
		t.Fatalf("GetCollectionIDs failed, err:%v\n", err)
	}
	t.Logf("GetCollectionIDs suceess\nids:%v\n", ids[0])
}

func TestAddViewNum(t *testing.T) {
	p := models.Trigger{PostID: 123}
	err := AddViewNum(&p)
	if err != nil {
		t.Fatalf("AddViewNum failed, err:%v\n", err)
	}
	t.Logf("AddViewNum suceess\n")
}

func TestAddFollow(t *testing.T) {
	follow := models.Follow{
		FollowID: "1",
	}
	uid := "1234"

	err := AddFollow(&follow, uid)
	if err != nil {
		t.Fatalf("AddFollow insert record into mysql failed, err:%v\n", err)
	}
	t.Logf("AddFollow insert record into mysql suceess\n")
}

func TestCancelFollow(t *testing.T) {
	follow := models.Follow{
		FollowID: "1",
	}
	uid := "1234"

	err := CancelFollow(&follow, uid)
	if err != nil {
		t.Fatalf("CancelFollow insert record into mysql failed, err:%v\n", err)
	}
	t.Logf("CancelFollow insert record into mysql suceess\n")
}

func TestGetFollowUser(t *testing.T) {
	uid := "3"
	data, err := GetFollowUser(uid)
	if err != nil {
		t.Fatalf("GetFollowUser failed, err:%v\n", err)
	}
	t.Logf("GetFollowUser suceess\ndata:%v\n", data[0])
}

func TestGetFollowedUser(t *testing.T) {
	uid := "3"
	data, err := GetFollowedUser(uid)
	if err != nil {
		t.Fatalf("GetFollowedUser failed, err:%v\n", err)
	}
	t.Logf("GetFollowedUser suceess\ndata:%v\n", data[0])
}

func TestGetFollowPostByIDs(t *testing.T) {
	p := models.ParamPostList{
		Page: 1,
		Size: 10,
	}
	ids := []string{"1", "123", "1234"}

	posts, err := GetFollowPostByIDs(&p, ids)
	if err != nil {
		t.Fatalf("GetFollowPostByIDs failed, err:%v\n", err)
	}
	t.Logf("GetFollowPostByIDs suceess\nposts: %v\n", posts[0])
}
