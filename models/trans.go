package models

import "time"

type Trans struct {
	TransID    int64     `json:"trans_id,string" db:"trans_id"`
	Status     int32     `json:"status" db:"status"`
	Title      string    `json:"title" db:"title" binding:"required"`
	Content    string    `json:"content" db:"content" binding:"required"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}
