package controllers

import (
	"Linkux/dao/mysql"
	"Linkux/dao/redis"
	"Linkux/pkg/snowflakes"
	"Linkux/settings"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
	err := mysql.Init(&dbCfg)
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

	if err := snowflakes.Init("2021-03-04", 1); err != nil {
		fmt.Println("Init snowflakes failed")
		return
	}
}

func TestLoginHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/login"
	r.POST(url, LoginHandler)

	body := `{
		"username": "xx",
		"pic_link": "xxx",
		"code": "x"
	}`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	//assert.Contains(t, w.Body.String(), "需要登录")
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, res.Code, CodeSuccess)
}

func TestGetUserContributionHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/contribution"
	r.GET(url, GetUserContributionHandler)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("size", "10")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, res.Code, CodeUserNotExist)
}

func TestAddCollectionHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/collect"
	r.POST(url, AddCollectionHandler)

	body := `{
		"post_id": 11111
	}`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	//assert.Contains(t, w.Body.String(), "需要登录")
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, res.Code, CodeUserNotExist)
}

func TestDeleteCollectionHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/collect/delete"
	r.PUT(url, DeleteCollectionHandler)

	body := `{
		"post_id": 22222
	}`

	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	//assert.Contains(t, w.Body.String(), "需要登录")
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, res.Code, CodeUserNotExist)
}

func TestGetCollectionHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/collect/get"
	r.GET(url, GetCollectionHandler)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("size", "10")
	req.URL.RawQuery = q.Encode()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, res.Code, CodeUserNotExist)
}

func TestAddViewHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/view/add"
	r.PUT(url, AddViewHandler)

	body := `{
		"post_id": 2222222
	}`

	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	//assert.Contains(t, w.Body.String(), "需要登录")
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, res.Code, CodeSuccess)
}

func TestAddFollowHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/follow"
	r.POST(url, AddFollowHandler)

	body := `{
		"follow_id": 11111
	}`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	//assert.Contains(t, w.Body.String(), "需要登录")
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, res.Code, CodeInvalidParam)
}

func TestCancelFollowHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/follow/cancel"
	r.PUT(url, CancelFollowHandler)

	body := `{
		"followed_id": "11111"
	}`

	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	//assert.Contains(t, w.Body.String(), "需要登录")
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, res.Code, CodeUserNotExist)
}

func TestGetFollowUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/follow/get/follow"
	r.GET(url, GetFollowUserHandler)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, res.Code, CodeUserNotExist)
}

func TestGetFollowedUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/follow/get/followed"
	r.GET(url, GetFollowedUserHandler)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, res.Code, CodeUserNotExist)
}

func TestGetFollowPostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/follow/get/post"
	r.GET(url, GetFollowPostHandler)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, res.Code, CodeUserNotExist)
}
