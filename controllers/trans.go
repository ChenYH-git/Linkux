package controllers

import (
	"Linkux/logic"
	"Linkux/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateTransTaskHandler 创建翻译任务接口
// @Summary 创建翻译任务接口
// @Description 按输入内容创建翻译任务
// @Tags 翻译相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body models.Trans true "创建翻译任务参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /trans [post]
func CreateTransTaskHandler(c *gin.Context) {
	p := new(models.Trans)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON failed", zap.Any("Err", err))
		zap.L().Error("create trans failed")
		ResponseError(c, CodeInvalidParam)
		return
	}

	if err := logic.CreateTrans(p); err != nil {
		zap.L().Error("logic.CreateTrans(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// GetTransTaskHandler 获取翻译任务接口
// @Summary 获取翻译任务接口
// @Description 获取翻译任务
// @Tags 翻译相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseTransList
// @Router /trans/get/task [get]
func GetTransTaskHandler(c *gin.Context) {
	data, err := logic.GetTransTask()
	if err != nil {
		zap.L().Error("logic.GetTransTask() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

// GetTransExistHandler 获取翻译任务对应的当前翻译文章接口
// @Summary 获取翻译任务对应的当前翻译文章接口
// @Description 获取翻译任务对应的当前翻译文章
// @Tags 翻译相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.ParamPostList true "翻译任务id参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /trans/get/exist [get]
func GetTransExistHandler(c *gin.Context) {
	p := &models.ParamPostList{
		Page: 1,
		Size: 10,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetTransExistHandler get query err: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetTransExist(p)
	if err != nil {
		zap.L().Error("logic.GetTransExist(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}
