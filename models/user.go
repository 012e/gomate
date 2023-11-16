package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           int64
	Name         string
	Email        string
	Username     string
	PasswordHash []byte

	HaveGroup sql.NullBool `gorm:"default:false"`
	GroupID   int64
}
