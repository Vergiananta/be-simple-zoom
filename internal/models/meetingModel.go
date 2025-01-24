package models

import "gorm.io/gorm"

type Meeting struct {
	gorm.Model
	ID 		   			uint     `gorm:"primaryKey" `
	Title      			string   `gorm:"not null" json:"title"`
	Description       	string   `gorm:"type:text" json:"body"`
	ActiveFrom 			string 		`gorm:"type:string" json:"activeFrom"`
	ActiveTo 			string `gorm:"type:string" json:"activeTo"`
	UserID     			uint     `gorm:"foreignkey:UserID" json:"userID"`
	User       			User     `gorm:"foreignkey:UserID"`
	CreatedAt 			string `gorm:"type:string" json:"createdAt"`
	CreatedBy 			uint `gorm:"type:string" json:"createdBy"`
	ModifiedAt 			string `gorm:"type:string" json:"modifiedAt"`
	ModifiedBy 			uint `gorm:"type:string" json:"modifiedBy"`
}
