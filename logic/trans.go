package logic

import (
	"Linkux/dao/mysql"
	"Linkux/models"
	"Linkux/pkg/snowflakes"

	"go.uber.org/zap"
)

func CreateTrans(p *models.Trans) (err error) {
	p.TransID = snowflakes.GenID()

	err = mysql.CreateTrans(p)
	if err != nil {
		return err
	}

	return
}

func GetTransTask() (data []*models.Trans, err error) {
	data, err = mysql.GetTransTask()
	if err != nil {
		return nil, err
	}
	return
}

func GetTransExist(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	data, err = mysql.GetTransExist(p)
	if err != nil {
		zap.L().Error("mysql.GetTransExist(p) failed",
			zap.Int64("trans_id", p.TransID),
			zap.Error(err))
		return nil, err
	}

	return data, nil
}
