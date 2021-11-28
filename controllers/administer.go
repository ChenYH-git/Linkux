package controllers

import (
	"Linkux/models"
	"Linkux/pkg/jwt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var administer = []string{"chenyuhan", "pengshiyang", "guonansheng", "aizhenwen", "huangyijie", "wujianhang", "liwending", "zhangquanyu", "zhenglingrui", "chenzonhang"}

// AdministerLoginHandler 管理员登录接口
// @Summary 管理员登录接口
// @Description 管理员登录
// @Tags 管理员相关接口
// @Accept application/json
// @Produce application/json
// @Param object body models.Administer true "管理员登录接口具体参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /administer/login [post]
func AdministerLoginHandler(c *gin.Context) {
	p := new(models.Administer)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Invalid login param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}

	var flag bool
	flag = false

	for _, v := range administer {
		if p.Name == v {
			flag = true
			break
		}
	}

	if !flag {
		zap.L().Error("not administer")
		ResponseError(c, CodeUserNotExist)
		return
	}

	token, err := jwt.GenAdToken(p.Name)
	if err != nil {
		zap.L().Error("Gen token failed, err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, gin.H{
		"token": token,
	})
}
