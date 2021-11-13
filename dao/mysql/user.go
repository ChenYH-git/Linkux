package mysql

import (
	"Linkux/models"
)

func InsertUser(user *models.User) (err error) {
	sqlStr := `select count(user_id) from user where user_id = ?`
	var count int
	if err := db.Get(&count, sqlStr, user.UserID); err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	sqlStr = `insert into user(user_id, contribution) values(?,?)`

	_, err = db.Exec(sqlStr, user.UserID, user.Contribution)
	return
}

func GetUserByID(uid string) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select username from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	return
}
