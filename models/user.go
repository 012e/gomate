package models

type User struct {
	Base         
	ID           int64  `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Email        string `json:"email,omitempty"`
	Username     string `json:"username,omitempty"`
	PasswordHash []byte `json:"password_hash,omitempty"`

	HaveGroup bool  `gorm:"default:false" json:"have_group,omitempty"`
	GroupID   int64 `json:"group_id,omitempty"`
}
