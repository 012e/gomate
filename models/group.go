package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	ID   int64
	Name string
	Description string
}

// users of a group
type GroupUser struct {
	gorm.Model
	ID int64

	Group   Group
	GroupID int64

	User   User
	UserID int64
}

// link to join groups
type GroupJoinCode struct {
	gorm.Model
	ID int64

	Group   Group
	GroupID int64

	Code string
}
