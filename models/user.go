package models

type User struct {
	Contribution int64  `db:"contribution"`             // 用户贡献度
	UserID       string `db:"user_id"`                  // 用户id
	Username     string `json:"username" db:"username"` // 用户名
	PicLink      string `json:"pic_link" db:"pic_link"` // 头像链接
	Code         string `json:"code"`                   // 微信返回的 code
	Token        string `json:"token"`                  // 后端返回的用户 token
}

type LoginResBody struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	Errmsg     string `json:"errmsg"`
	Errcode    int    `json:"errcode"`
}

type Trigger struct {
	PostID int64 `json:"post_id,string" db:"post_id"` // 对应帖子id
}

type Follow struct {
	FollowID   string `json:"follow_id" db:"follow_id"`     // 关注的作者id
	FollowedID string `json:"followed_id" db:"followed_id"` // 粉丝id
}
