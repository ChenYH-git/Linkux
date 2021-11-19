package models

type User struct {
	Contribution int64  `db:"contribution"`
	UserID       string `db:"user_id"`
	Username     string `json:"username" db:"username"`
	PicLink      string `json:"pic_link" db:"pic_link"`
	Code         string `json:"code"`
	Token        string `json:"token"`
}

type LoginResBody struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	Errmsg     string `json:"errmsg"`
	Errcode    int    `json:"errcode"`
}

type Triger struct {
	PostID int64 `json:"post_id" db:"post_id"`
}

type Follow struct {
	FollowID   string `json:"follow_id" db:"follow_id"`
	FollowedID string `json:"followed_id" db:"followed_id"`
}
