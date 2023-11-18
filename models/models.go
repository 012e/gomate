package models

import (
	"database/sql"
	"time"
)

type Base struct {
	ID        int64        `gorm:"primaryKey" json:"id,omitempty"`
	CreatedAt time.Time    `json:"created_at,omitempty"`
	UpdatedAt time.Time    `json:"updated_at,omitempty"`

	// DeletedAt shouldn't be used json encoding, since sql.NullTime doesn't implement json literals
	// and it should exist the moment users access the resource.
	DetetedAt sql.NullTime `gorm:"index" json:"-"`
}
