package models

import "time"

type Post struct {
	PostID     int64     `json:"post_id,string" db:"post_id"`               // 帖子id，由后端生成，无需填入
	TransID    int64     `json:"trans_id,string" db:"trans_id"`             // 翻译任务id，不为0说明帖子是翻译帖，对应某个翻译任务，按情况填入
	LabelID    int64     `json:"label_id" db:"label_id" binding:"required"` // 社区标签id，必填
	ViewdNum   int64     `json:"viewd_num" db:"viewd_num"`                  // 观看量，无需填入
	CollectNum int64     `json:"collect_num" db:"collect_num"`              // 收藏量，无需填入
	Status     int8      `json:"status" db:"status"`                        // 帖子状态，管理审核时使用，无需填入
	AuthorID   string    `json:"author_id" db:"author_id"`                  // 作者id，无需填入
	Title      string    `json:"title" db:"title" binding:"required"`       // 标题，必填
	Content    string    `json:"content" db:"content" binding:"required"`   // 内容，必填
	Qualified  bool      `json:"qualified"`                                 // 是否加精
	CreateTime time.Time `json:"create_time" db:"create_time"`              // 创建时间，无需填入
}

type ApiPostDetail struct {
	AuthorName      string         `json:"author_name"`        // 作者姓名
	AuthorQualified bool           `json:"author_qualified"`   // 作者是否加v
	VoteNum         int64          `json:"vote_num"`           // 点赞数
	PicLink         string         `json:"pic_link,omitempty"` // 作者头像链接
	*Post                          // 帖子详情
	*LabelDetail    `json:"label"` // 社区标签详情
}
