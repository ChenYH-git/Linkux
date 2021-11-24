package models

type ApiRankDetail struct {
	Username     string `json:"username" db:"username"`         // 用户名
	Contribution int64  `json:"contribution" db:"contribution"` // 贡献度
	PicLink      string `json:"pic_link" db:"pic_link"`         // 头像链接
}
