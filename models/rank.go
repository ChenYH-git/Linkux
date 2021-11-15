package models

type ApiRankDetail struct {
	Username     string `json:"username" db:"username"`
	Contribution int64  `json:"contribution" db:"contribution"`
	PicLink      string `json:"pic_link" db:"pic_link"`
}
