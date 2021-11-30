package models

type Administer struct {
	Name string `json:"name" binding:"required"` // 管理员名称，校验用
}

type StarUser struct {
	UserID string `json:"user_id"` // 用户id
}

type Judge struct {
	PostID  int64 `json:"post_id,string" db:"post_id"`      // 如果审核的是帖子，那么这项要填，否则为0
	TransID int64 `json:"trans_id,string" db:"trans_id"`    // 如果审核的是任务，那么这项要填，否则为0
	LabelID int64 `json:"label_id,omitempty" db:"label_id"` // 如果审核的是帖子，那么这项要填，否则为0
}

type Task struct {
	TransID int64 `json:"trans_id,string"` // 翻译任务id
}
