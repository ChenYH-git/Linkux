package mysql

import (
	"Linkux/models"
	"database/sql"

	"go.uber.org/zap"
)

func GetLabelList() (labelList []*models.Label, err error) {
	sqlStr := `select label_id, label_name from label`
	if err := db.Select(&labelList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("No label in db")
			err = nil
		}
	}
	return
}
