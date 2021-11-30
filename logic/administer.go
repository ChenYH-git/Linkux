package logic

import (
	"Linkux/dao/mysql"
	"Linkux/dao/redis"
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

	if err = redis.DeletePosts(p.PostID); err != nil {
		zap.L().Error("redis.DeletePosts(p) failed", zap.Error(err))
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

func GetWaitingTransTask(p *models.ParamPostList) (data []*models.Trans, err error) {
	data, err = mysql.GetWaitingTransTask(p)
	if err != nil {
		zap.L().Error("mysql.GetWaitingTransTask err:", zap.Error(err))
		return
	}
	return
}

func GetWaitingPosts(p *models.ParamPostList) (data []*models.Post, err error) {
	data, err = mysql.GetWaitingPostsTask(p)
	if err != nil {
		zap.L().Error("mysql.GetWaitingPostsTask err:", zap.Error(err))
		return
	}
	return
}

func JudgePass(p *models.Judge) (err error) {
	err = mysql.JudgePass(p)
	if err != nil {
		zap.L().Error("mysql.JudgePass(p) err:", zap.Error(err))
		return
	}

	if p.TransID != 0 || p.PostID == 0 {
		return err
	}
	err = redis.CreateRedisPost(p.PostID, p.LabelID)
	if err != nil {
		return err
	}

	AuthorID, err := mysql.GetAuthorByPostID(p.PostID)
	if err != nil {
		zap.L().Error("mysql.GetAuthorByPostID err:", zap.Error(err))
		return
	}
	err = mysql.AddContribution(AuthorID)
	if err != nil {
		return err
	}
	return
}

func DeleteTrans(p *models.Task) (err error) {
	if err = mysql.DeleteTrans(p); err != nil {
		zap.L().Error("mysql.DeleteTrans(p) failed", zap.Error(err))
		return err
	}
	return
}
