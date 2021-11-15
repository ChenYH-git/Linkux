package models

type User struct {
	Contribution int64  `db:"contribution"`
	UserID       string `db:"user_id"`
	Username     string `json:"username" db:"username"`
	PicLink      string `json:"pic_link" db:"pic_link"`
	Code         string `json:"code"`
}

type LoginResBody struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	Errmsg     string `json:"errmsg"`
	Errcode    int    `json:"errcode"`
}

type Collection struct {
	PostID int64 `json:"post_id" db:"post_id"`
}
