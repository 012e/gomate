package models

type Group struct {
	Base
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// users of a group
type GroupUser struct {
	Base

	Group   Group
	GroupID int64 `gorm:"not null"`

	User   User
	UserID int64 `gorm:"not null"`
}

// link to join groups
type GroupJoinCode struct {
	Base
	ID int64

	Group   Group
	GroupID int64

	Code string
}
