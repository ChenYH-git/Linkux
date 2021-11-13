package logic

import (
	"Linkux/dao/mysql"
	"Linkux/models"
)

func GetLabelpostList() ([]*models.Label, error) {
	return mysql.GetLabelList()
}

func GetLabelDetail(id int64) (*models.LabelDetail, error) {
	return mysql.GetLabelDetailByID(id)
}
