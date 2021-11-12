package models

type User struct {
	Contribution int64  `db:"contribution"`
	UserID       string `db:"user_id"`
}

type LoginResBody struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}
