package mysql

import (
	"Linkux/models"
)

func GetUserRank() (data []*models.ApiRankDetail, err error) {
	sqlStr := `select
	username, contribution, pic_link
	from user
	order by contribution
	desc
	`
	data = make([]*models.ApiRankDetail, 0, 10)
	err = db.Select(&data, sqlStr)
	return
}

func GetMyRank(userID string) ([]*models.ApiRankDetail, error) {
	sqlStr := `select
	username, contribution, pic_link
	from user
	where user_id = ?
	`
	me := make([]*models.ApiRankDetail, 0, 2)
	err := db.Select(&me, sqlStr, userID)
	if err != nil {
		return nil, err
	}
	return me, err
}
