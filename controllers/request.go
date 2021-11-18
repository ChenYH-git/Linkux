package controllers

import (
	"errors"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	ErrorChangeString = errors.New("强制转换为string失败")
	ErrorNoUser       = errors.New("无此用户")
)

const CtxUserIDKey = "userID"

var Mu sync.RWMutex

var Keys map[string]string

func getCurrentUserID(key string) (value string, err error) {
	Mu.RLock()
	value, exists := Keys[key]
	if !exists {
		err = errors.New("获取id失败")
		return "", err
	}
	Mu.RUnlock()
	return value, nil
}

func setCurrentUserID(key string, value string) {
	Mu.Lock()
	if Keys == nil {
		Keys = make(map[string]string)
	}

	Keys[key] = value
	Mu.Unlock()
}

func getPageInfo(c *gin.Context) (int64, int64) {
	pageNumStr := c.Query("page")
	sizeStr := c.Query("size")

	var (
		size    int64
		pageNum int64
		err     error
	)

	pageNum, err = strconv.ParseInt(pageNumStr, 10, 64)
	if err != nil {
		pageNum = 0
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return pageNum, size
}
