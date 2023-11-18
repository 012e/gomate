package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Base     
	Username string   
	Expiry   time.Time
	Token    uuid.UUID
}

