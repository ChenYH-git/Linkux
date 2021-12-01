package mysql

import (
	"Linkux/models"
)

func GetUserRank() (data []*models.User, err error) {
	sqlStr := `select
	username, contribution, pic_link
	from user
	order by contribution
	desc
	limit 10
	`
	data = make([]*models.User, 0, 10)
	err = db.Select(&data, sqlStr)
	if err != nil {
		return nil, err
	}

	sqlStr = `select count(user_id) from vuser where user_id = ?`

	for i, _ := range data {
		var count int
		err = db.Get(&count, sqlStr, data[i].UserID)
		if err != nil {
			return nil, err
		}
		if count < 1 {
			data[i].Qualified = false
			continue
		}
		data[i].Qualified = true
	}
	return
}

func GetMyRank(userID string) ([]*models.User, error) {
	sqlStr := `select
	username, contribution, pic_link
	from user
	where user_id = ?
	`
	me := make([]*models.User, 0, 2)
	err := db.Select(&me, sqlStr, userID)
	if err != nil {
		return nil, err
	}

	var count int
	sqlStr = `select count(user_id) from vuser where user_id = ?`
	err = db.Get(&count, sqlStr, userID)
	if err != nil {
		return nil, err
	}

	if count < 1 {
		me[0].Qualified = false
	} else {
		me[0].Qualified = true
	}
	return me, err
}
