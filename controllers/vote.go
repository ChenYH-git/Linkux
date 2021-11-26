package controllers

import (
	"Linkux/logic"
	"Linkux/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// PostVoteController 点赞投票接口
// @Summary 点赞投票接口
// @Description 根据帖子id和投票信息进行点赞、踩
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body models.ParamVoteData true "翻译任务id参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /vote [post]
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

	userID, err := getCurrentUserID(c)
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

// GetVoteCollectHandler 获取当前用户对帖子是否点赞收藏的接口
// @Summary 获取当前用户对帖子是否点赞收藏的接口
// @Description 获取当前用户对帖子是否点赞收藏
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body models.Trigger true "两个 bool 参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /getvc [post]
func GetVoteCollectHandler(c *gin.Context) {
	p := new(models.Trigger)
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

	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	vc := new(models.VCorNot)

	vc.Voted, vc.Collected, err = logic.GetVoteCollect(p, userID)
	if err != nil {
		zap.L().Error("logic.GetVoteCollect() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, vc)
}
