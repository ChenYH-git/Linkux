package controllers

import (
	"Linkux/logic"
	"Linkux/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SearchHandler 帖子搜索接口
// @Summary 帖子搜索接口
// @Description 按输入内容检索帖子
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param object query models.ParamPostList true "搜索参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /search [get]
func SearchHandler(c *gin.Context) {
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderScore,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("SearchHandler get query err: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetSearchRes(p)
	if err != nil {
		zap.L().Error("logic.GetSearchRes(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
