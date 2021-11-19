package controllers

import (
	"Linkux/logic"
	"Linkux/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

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
