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

	setCurrentUserID(CtxUserIDKey+userID, userID)

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
		zap.L().Error("GetUserContributionHandler get query err: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	userID, err := getCurrentUserID(CtxUserIDKey)
	if err != nil {
		zap.L().Error("getCurrentUserID failed", zap.Error(err))
		ResponseError(c, CodeUserNotExist)
		return
	}

	data, err := logic.GetUserConByID(p, userID)
	if err != nil {
		zap.L().Error("logic.GetUserConByID() failed,", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

func AddCollectionHandler(c *gin.Context) {
	p := new(models.Triger)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("invalid collection param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	userID, err := getCurrentUserID(CtxUserIDKey)
	if err != nil {
		zap.L().Error("getCurrentUserID failed", zap.Error(err))
		ResponseError(c, CodeUserNotExist)
		return
	}

	if err := logic.AddCollection(p, userID); err != nil {
		zap.L().Error("logic.AddCollection(userID) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, gin.H{
		"msg": "add collection success",
	})
}

func DeleteCollectionHandler(c *gin.Context) {
	p := new(models.Triger)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("invalid collection param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	userID, err := getCurrentUserID(CtxUserIDKey)
	if err != nil {
		zap.L().Error("getCurrentUserID failed", zap.Error(err))
		ResponseError(c, CodeUserNotExist)
		return
	}

	if err := logic.DeleteCollection(p, userID); err != nil {
		zap.L().Error("logic.AddCollection(userID) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, gin.H{
		"msg": "delete collection success",
	})
}

func GetCollectionHandler(c *gin.Context) {
	p := &models.ParamPostList{
		Page: 1,
		Size: 10,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetCollectionHandler get query err: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	userID, err := getCurrentUserID(CtxUserIDKey)
	if err != nil {
		zap.L().Error("getCurrentUserID failed", zap.Error(err))
		ResponseError(c, CodeUserNotExist)
		return
	}

	data, err := logic.GetCollection(p, userID)
	if err != nil {
		zap.L().Error("logic.GetCollection(p, userID) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

func AddViewHandler(c *gin.Context) {
	p := new(models.Triger)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("invalid view param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	if err := logic.AddViewNum(p); err != nil {
		zap.L().Error("logic.AddViewNum(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, gin.H{
		"msg": "add viewNum success",
	})
}

func AddFollowHandler(c *gin.Context) {
	p := new(models.Follow)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("invalid follow param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	userID, err := getCurrentUserID(CtxUserIDKey)
	if err != nil {
		zap.L().Error("getCurrentUserID failed", zap.Error(err))
		ResponseError(c, CodeUserNotExist)
		return
	}

	if err := logic.AddFollow(p, userID); err != nil {
		zap.L().Error("logic.AddFollow(p, userID) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, gin.H{
		"msg": "add follow success",
	})
}

func CancelFollowHandler(c *gin.Context) {
	p := new(models.Follow)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("invalid follow param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	userID, err := getCurrentUserID(CtxUserIDKey)
	if err != nil {
		zap.L().Error("getCurrentUserID failed", zap.Error(err))
		ResponseError(c, CodeUserNotExist)
		return
	}

	if err := logic.CancelFollow(p, userID); err != nil {
		zap.L().Error("logic.CancelFollow(p, userID) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, gin.H{
		"msg": "cancel follow success",
	})
}

func GetFollowUserHandler(c *gin.Context) {
	userID, err := getCurrentUserID(CtxUserIDKey)
	if err != nil {
		zap.L().Error("getCurrentUserID failed", zap.Error(err))
		ResponseError(c, CodeUserNotExist)
		return
	}

	data, err := logic.GetFollowUser(userID)
	if err != nil {
		zap.L().Error("logic.GetFollowUser(userID) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

func GetFollowedUserHandler(c *gin.Context) {
	userID, err := getCurrentUserID(CtxUserIDKey)
	if err != nil {
		zap.L().Error("getCurrentUserID failed", zap.Error(err))
		ResponseError(c, CodeUserNotExist)
		return
	}

	data, err := logic.GetFollowedUser(userID)
	if err != nil {
		zap.L().Error("logic.GetFollowedUser(userID) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

func GetFollowPostHandler(c *gin.Context) {
	p := &models.ParamPostList{
		Page: 1,
		Size: 10,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetFollowPostHandler get query err: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	userID, err := getCurrentUserID(CtxUserIDKey)
	if err != nil {
		zap.L().Error("getCurrentUserID failed", zap.Error(err))
		ResponseError(c, CodeUserNotExist)
		return
	}

	data, err := logic.GetFollowPost(p, userID)
	if err != nil {
		zap.L().Error("logic.GetFollowedUser(userID) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}
