package models

import "time"

type Trans struct {
	TransID    int64     `json:"trans_id,string" db:"trans_id"`           // 翻译任务id，不用填
	Status     int32     `json:"status" db:"status"`                      // 任务帖状态，审核用
	Title      string    `json:"title" db:"title" binding:"required"`     // 标题，必填
	Content    string    `json:"content" db:"content" binding:"required"` // 内容，必填
	CreateTime time.Time `json:"create_time" db:"create_time"`            // 创建时间，不用填
}
