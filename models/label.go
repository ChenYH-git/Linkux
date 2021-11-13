package models

import "time"

type Label struct {
	ID   int64  `json:"id" db:"label_id"`
	Name string `json:"name" db:"label_name"`
}

type LabelDetail struct {
	ID           int64     `json:"id" db:"label_id"`
	Name         string    `json:"name" db:"label_name"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
}
