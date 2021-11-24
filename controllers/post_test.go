package controllers

import (
	"Linkux/dao/mysql"
	"Linkux/dao/redis"
	"Linkux/settings"
	"bytes"
	"encoding/json"
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
}

func TestCreatePostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/post"
	r.POST(url, CreatePostHandler)

	body := `{
		"label_id": 4,
		"content": "TestCreatePostHandler"
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

func TestIndexHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/index"
	r.GET(url, IndexHandler)

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
	assert.Equal(t, res.Code, CodeSuccess)
}
