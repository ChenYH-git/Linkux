package controllers

import (
	"Linkux/logic"
	"Linkux/models"
	"Linkux/pkg/jwt"

	"github.com/go-playground/validator/v10"

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
		ResponseError(c, CodeNotAdminister)
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

// ExamineGetTaskHandler 管理员获取待审核翻译任务接口
// @Summary 管理员获取待审核翻译任务接口
// @Description 管理员获取待审核翻译任务
// @Tags 管理员相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.ParamPostList false "分页参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /administer/examine/gettask [get]
func ExamineGetTaskHandler(c *gin.Context) {
	p := &models.ParamPostList{
		Page: 1,
		Size: 10,
	}

	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("ExamineGetTaskHandler get query err: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	transdata, err := logic.GetWaitingTransTask(p)
	if err != nil {
		zap.L().Error("logic.GetWaitingTransTask() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, transdata)
}

// ExamineGetPostsHandler 管理员获取待审核帖子接口
// @Summary 管理员获取待审核帖子接口
// @Description 管理员获取待审核帖子
// @Tags 管理员相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.ParamPostList false "分页参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /administer/examine/getposts [get]
func ExamineGetPostsHandler(c *gin.Context) {
	p := &models.ParamPostList{
		Page: 1,
		Size: 10,
	}

	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("ExamineGetPostsHandler get query err: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	postsdata, err := logic.GetWaitingPosts(p)
	if err != nil {
		zap.L().Error("logic.GetWaitingPosts() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, postsdata)
}

// ExaminePutChangesHandler 管理员审核通过接口
// @Summary 管理员审核通过接口
// @Description 此接口仅用于通过帖子或翻译任务审核
// @Tags 管理员相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body models.Judge true "（帖子id、社区id）或（翻译任务id）参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /administer/examine/put [put]
func ExaminePutChangesHandler(c *gin.Context) {
	p := new(models.Judge)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Invalid ExaminePutChangesHandler param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}

	if err := logic.JudgePass(p); err != nil {
		zap.L().Error("logic.JudgePass() failed, err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// GetPostsExistsHandler 管理员获取现有帖子列表接口
// @Summary 管理员获取现有帖子列表接口
// @Description 管理员获取现有帖子
// @Tags 管理员相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.ParamPostList false "分页排序参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /administer/posts/get [get]
func GetPostsExistsHandler(c *gin.Context) {
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderScore,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostsExistsHandler get query err: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetPostList(p)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// StarPostsHandler 管理员帖子加精接口
// @Summary 管理员帖子加精接口
// @Description 管理员帖子加精
// @Tags 管理员相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body models.Trigger true "帖子id参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /administer/posts/star [put]
func StarPostsHandler(c *gin.Context) {
	p := new(models.Trigger)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}

	if err := logic.StarPosts(p); err != nil {
		zap.L().Error("logic.StarPosts(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// CancelStarPostsHandler 管理员帖子取消加精接口
// @Summary 管理员帖子取消加精接口
// @Description 管理员帖子取消加精
// @Tags 管理员相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body models.Trigger true "帖子id参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /administer/posts/star/cancel [put]
func CancelStarPostsHandler(c *gin.Context) {
	p := new(models.Trigger)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}

	if err := logic.CancelStarPosts(p); err != nil {
		zap.L().Error("logic.CancelStarPosts(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// DeletePostsHandler 管理员删除帖子接口
// @Summary 管理员删除帖子接口
// @Description 管理员删除帖子
// @Tags 管理员相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body models.Trigger true "帖子id参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /administer/posts/delete [delete]
func DeletePostsHandler(c *gin.Context) {
	p := new(models.Trigger)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}

	if err := logic.DeletePosts(p); err != nil {
		zap.L().Error("logic.DeletePosts(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// DeleteTransHandler 管理员删除翻译任务接口
// @Summary 管理员删除翻译任务接口
// @Description 管理员删除翻译任务
// @Tags 管理员相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body models.Task true "翻译任务id参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /administer/trans/delete [delete]
func DeleteTransHandler(c *gin.Context) {
	p := new(models.Task)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}

	if err := logic.DeleteTrans(p); err != nil {
		zap.L().Error("logic.DeleteTrans(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// GetPostStatusHandler 获取当前帖子是否加精的接口
// @Summary 获取当前帖子是否加精的接口
// @Description 获取当前帖子是否加精
// @Tags 管理员相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body models.Trigger true "帖子id参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /administer/getp [post]
func GetPostStatusHandler(c *gin.Context) {
	p := new(models.Trigger)
	var err error
	if err = c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}

	vc := new(models.PUorNot)

	vc.Qualified, err = logic.GetPostStatus(p)
	if err != nil {
		zap.L().Error("logic.GetPostStatus() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, vc)
}

// GetUserExistsHandler 管理员获取用户接口
// @Summary 管理员获取用户接口
// @Description 管理员获取用户
// @Tags 管理员相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.ParamPostList false "分页参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /administer/user/get [get]
func GetUserExistsHandler(c *gin.Context) {
	p := &models.ParamPostList{
		Page: 1,
		Size: 10,
	}

	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetUserExistHandler get query err: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetAllUser(p)
	if err != nil {
		zap.L().Error("logic.GetAllUser(p) failed, err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

// StarUserHandler 管理员用户加v接口
// @Summary 管理员用户加v接口
// @Description 管理员用户加v
// @Tags 管理员相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body models.StarUser true "用户id参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /administer/user/star [put]
func StarUserHandler(c *gin.Context) {
	p := new(models.StarUser)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}

	if err := logic.StarUser(p); err != nil {
		zap.L().Error("logic.StarUser(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// CancelStarUserHandler 管理员用户取消加v接口
// @Summary 管理员用户取消加v接口
// @Description 管理员用户取消加v
// @Tags 管理员相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body models.StarUser true "用户id参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /administer/user/star/cancel [put]
func CancelStarUserHandler(c *gin.Context) {
	p := new(models.StarUser)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}

	if err := logic.CancelStarUser(p); err != nil {
		zap.L().Error("logic.CancelStarUser(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// GetUserStatusHandler 获取当前用户是否加v的接口
// @Summary 获取当前用户是否加v的接口
// @Description 获取当前用户是否加v
// @Tags 管理员相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body models.StarUser true "用户id参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsg
// @Router /administer/getu [post]
func GetUserStatusHandler(c *gin.Context) {
	p := new(models.StarUser)
	var err error
	if err = c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}

	vc := new(models.PUorNot)

	vc.Qualified, err = logic.GetUserStatus(p)
	if err != nil {
		zap.L().Error("logic.GetUserStatus() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, vc)
}
