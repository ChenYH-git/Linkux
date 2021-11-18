package controllers

import (
	"Linkux/logic"
	"Linkux/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PostVoteController(c *gin.Context) {
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) // 类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans)) // 翻译并去除掉错误提示中的结构体标识
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}

	userID, err := getCurrentUserID(CtxUserIDKey)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}
