package models

import "time"

type Post struct {
	ID         int64     `json:"id,string" db:"post_id"`
	LabelID    int64     `json:"label_id" db:"label_id" binding:"required"`
	Status     int32     `json:"status" db:"status"`
	AuthorID   string    `json:"author_id" db:"author_id"`
	Title      string    `json:"title" db:"title" binding:"required"`
	Content    string    `json:"content" db:"content" binding:"required"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}

type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	VoteNum    int64  `json:"vote_num"`
	PicLink    string `json:"pic_link,omitempty"`
	*Post
	*LabelDetail `json:"label"`
}
