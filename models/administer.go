package models

type Administer struct {
	Name string `json:"name" binding:"required"` // 管理员名称，校验用
}

type StarUser struct {
	UserID string `json:"user_id"` // 用户id
}
