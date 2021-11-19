package controllers

import (
	"Linkux/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

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
		"msg":  "success",
	})
}
