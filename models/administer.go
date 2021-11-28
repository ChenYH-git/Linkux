package models

type Administer struct {
	Name string `json:"name" binding:"required"`
}
