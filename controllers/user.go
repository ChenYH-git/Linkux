package controllers

import (
	"Linkux/logic"
	"Linkux/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoginHandler 用户登录接口
// @Summary 用户登录接口
// @Description 微信一键登录后端接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param object body models.User true "用户具体参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	p := new(models.User)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Invalid login param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}

	var token string
	var err error

	if token, err = logic.Login(p); err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, gin.H{
		"token": token,
	})
}

// GetUserContributionHandler 获取`我的贡献`接口
// @Summary 获取`我的贡献`接口
// @Description 获取`我的贡献`
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /contribution [get]
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
	userID, err := getCurrentUserID(c)
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

// AddCollectionHandler 加入收藏接口
// @Summary 加入收藏接口
// @Description 加入收藏
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body models.Trigger true "帖子id参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /collect [post]
func AddCollectionHandler(c *gin.Context) {
	p := new(models.Trigger)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("invalid collection param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	userID, err := getCurrentUserID(c)
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

	ResponseSuccess(c, nil)
}

// DeleteCollectionHandler 取消收藏接口
// @Summary 取消收藏接口
// @Description 取消收藏
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body models.Trigger true "帖子id参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /collect/delete [put]
func DeleteCollectionHandler(c *gin.Context) {
	p := new(models.Trigger)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("invalid collection param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	userID, err := getCurrentUserID(c)
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

	ResponseSuccess(c, nil)
}

// GetCollectionHandler 获取收藏列表接口
// @Summary 获取收藏列表接口
// @Description 获取收藏列表
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.ParamPostList true "分页参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /collect/get [get]
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
	userID, err := getCurrentUserID(c)
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

// AddViewHandler 观看量+1接口
// @Summary 观看量+1接口
// @Description 观看量+1
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body models.Trigger true "帖子id参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /view/add [put]
func AddViewHandler(c *gin.Context) {
	p := new(models.Trigger)
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

	ResponseSuccess(c, nil)
}

// AddFollowHandler 关注作者接口
// @Summary 关注作者接口
// @Description 关注作者
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body models.Follow true "作者id参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /follow [post]
func AddFollowHandler(c *gin.Context) {
	p := new(models.Follow)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("invalid follow param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	userID, err := getCurrentUserID(c)
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

	ResponseSuccess(c, nil)
}

// CancelFollowHandler 取消关注作者接口
// @Summary 取消关注作者接口
// @Description 取消关注作者
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body models.Follow true "作者id参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /follow/cancel [put]
func CancelFollowHandler(c *gin.Context) {
	p := new(models.Follow)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("invalid follow param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	userID, err := getCurrentUserID(c)
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

	ResponseSuccess(c, nil)
}

// GetFollowUserHandler 获取关注作者接口
// @Summary 获取关注作者接口
// @Description 获取关注作者
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseFollowList
// @Router /follow/get/follow [get]
func GetFollowUserHandler(c *gin.Context) {
	userID, err := getCurrentUserID(c)
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

// GetFollowedUserHandler 获取粉丝接口
// @Summary 获取粉丝接口
// @Description 获取粉丝
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseFollowList
// @Router /follow/get/followed [get]
func GetFollowedUserHandler(c *gin.Context) {
	userID, err := getCurrentUserID(c)
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

// GetFollowPostHandler 获取关注帖子接口
// @Summary 获取关注帖子接口
// @Description 获取关注帖子
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.ParamPostList true "分页参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /follow/get/post [get]
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

	userID, err := getCurrentUserID(c)
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
