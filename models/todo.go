package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID int64

	Title       string
	Description string
	Finished    sql.NullBool `gorm:"default:false"`
}
