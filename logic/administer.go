package logic

import (
	"Linkux/dao/mysql"
	"Linkux/models"

	"go.uber.org/zap"
)

func GetAllUser(p *models.ParamPostList) (data []*models.ApiRankDetail, err error) {
	data, err = mysql.GetAllUser(p)
	if err != nil {
		zap.L().Error("mysql.GetAllUser() failed", zap.Error(err))
		return nil, err
	}
	return
}

func StarPosts(p *models.Trigger) (err error) {
	if err = mysql.StarPosts(p); err != nil {
		zap.L().Error("mysql.StarPosts(p) failed", zap.Error(err))
		return err
	}
	return
}

func CancelStarPosts(p *models.Trigger) (err error) {
	if err = mysql.CancelStarPosts(p); err != nil {
		zap.L().Error("mysql.CancelStarPosts(p) failed", zap.Error(err))
		return err
	}
	return
}

func DeletePosts(p *models.Trigger) (err error) {
	if err = mysql.DeletePosts(p); err != nil {
		zap.L().Error("mysql.DeletePosts(p) failed", zap.Error(err))
		return err
	}
	return
}

func GetPostStatus(p *models.Trigger) (qualified bool, err error) {
	qualified, err = mysql.CheckPost(p)
	if err != nil {
		zap.L().Error("CheckPost",
			zap.Int64("postID", p.PostID),
			zap.Error(err))
		return false, err
	}
	return
}

func StarUser(p *models.StarUser) (err error) {
	if err = mysql.StarUser(p); err != nil {
		zap.L().Error("mysql.StarUser(p) failed", zap.Error(err))
		return err
	}
	return
}

func CancelStarUser(p *models.StarUser) (err error) {
	if err = mysql.CancelStarUser(p); err != nil {
		zap.L().Error("mysql.CancelStarUser(p) failed", zap.Error(err))
		return err
	}
	return
}

func GetUserStatus(p *models.StarUser) (qualified bool, err error) {
	qualified, err = mysql.CheckUser(p)
	if err != nil {
		zap.L().Error("CheckUser",
			zap.String("userID", p.UserID),
			zap.Error(err))
		return false, err
	}
	return
}
