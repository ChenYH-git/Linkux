package controllers

import (
	"Linkux/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetUserRankHandler 排行榜接口
// @Summary 排行榜接口
// @Description 返回排行榜数据
// @Tags 排行榜相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} _RankResponseMsg
// @Router /rank [get]
func GetUserRankHandler(c *gin.Context) {
	data, err := logic.GetUserRank()
	if err != nil {
		zap.L().Error("logic.GetUserRank() failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	userID, err := getCurrentUserID(c)

	if err != nil {
		zap.L().Error("rank: get my id failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}

	me, err := logic.GetMyRank(userID)
	if err != nil {
		zap.L().Error("get my rank failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, gin.H{
		"me":   me,
		"rank": data,
	})
}
