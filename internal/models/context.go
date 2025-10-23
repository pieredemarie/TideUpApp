package models

type Context struct {
	ID int `json:"context" gorm:"primaryKey"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	IsHidden bool `json:"ishidden"`
	Tasks []Task `gorm:"foreignKey:ContextID"`
}



