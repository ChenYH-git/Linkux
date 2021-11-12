package controllers

import (
	"Linkux/logic"
	"Linkux/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoginHandler(c *gin.Context) {
	code := c.Query("code")
	var err error
	user := new(models.User)

	if user, err = logic.Login(code); err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	c.Set(CtxUserIDKey, user.UserID)
	ResponseSuccess(c, gin.H{
		"openid": user.UserID,
	})
}
