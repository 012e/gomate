package models

type Todo struct {
	Base        
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Finished    bool   `gorm:"default:false" json:"finished,omitempty"`
}
