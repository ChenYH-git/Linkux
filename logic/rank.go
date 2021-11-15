package logic

import (
	"Linkux/dao/mysql"
	"Linkux/models"

	"go.uber.org/zap"
)

func GetUserRank() (data []*models.ApiRankDetail, err error) {
	data, err = mysql.GetUserRank()
	if err != nil {
		zap.L().Error("mysql.GetUserRank() failed", zap.Error(err))
		return nil, err
	}
	return
}

func GetMyRank(userID string) (data []*models.ApiRankDetail, err error) {
	data, err = mysql.GetMyRank(userID)
	if err != nil {
		return nil, err
	}
	return
}
