package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Username string
	Expiry   time.Time
	Token    uuid.UUID
}
