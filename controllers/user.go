package controllers

import (
	"Linkux/logic"
	"Linkux/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoginHandler(c *gin.Context) {
	p := new(models.User)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Invalid login param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}

	var userID string
	var err error

	if userID, err = logic.Login(p); err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	c.Set(CtxUserIDKey, userID)
	ResponseSuccess(c, gin.H{
		"openid": userID,
	})
}

func GetUserContributionHandler(c *gin.Context) {
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderScore,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("IndexHandler get query err: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//userID, err := getCurrentUserID(c)
	//if err != nil {
	//	zap.L().Error("getCurrentUserID failed", zap.Error(err))
	//	ResponseError(c, CodeUserNotExist)
	//	return
	//}

	userID := "0"

	data, err := logic.GetUserConByID(p, userID)
	if err != nil {
		zap.L().Error("logic.GetUserConByID() failed,", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

func AddCollectionHandler(c *gin.Context) {
	p := new(models.Collection)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("invalid collection param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//userID, err := getCurrentUserID(c)
	//if err != nil {
	//	zap.L().Error("getCurrentUserID failed", zap.Error(err))
	//	ResponseError(c, CodeUserNotExist)
	//	return
	//}

	userID := "0"
	if err := logic.AddCollection(p, userID); err != nil {
		zap.L().Error("logic.AddCollection(userID) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, gin.H{
		"msg": "add collection success",
	})
}

func GetCollectionHandler(c *gin.Context) {
	p := &models.ParamPostList{
		Page: 1,
		Size: 10,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("IndexHandler get query err: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//userID, err := getCurrentUserID(c)
	//if err != nil {
	//	zap.L().Error("getCurrentUserID failed", zap.Error(err))
	//	ResponseError(c, CodeUserNotExist)
	//	return
	//}

	userID := "0"
	data, err := logic.GetCollection(p, userID)
	if err != nil {
		zap.L().Error("logic.GetCollection(p, userID) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}
