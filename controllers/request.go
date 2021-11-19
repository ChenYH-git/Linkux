package controllers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ErrorChangeString = errors.New("强制转换为string失败")
	ErrorNoUser       = errors.New("无此用户")
)

const CtxUserIDKey = "userID"

func getCurrentUserID(c *gin.Context) (userID string, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorNoUser
		return
	}
	userID, ok = uid.(string)
	if !ok {
		err = ErrorChangeString
		return
	}
	return
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
