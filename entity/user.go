package entity

import "time"

//Book struct represents books table in database
type User struct {
	ID				      uint64 `gorm:"primary_key:auto_increment" json:"user_id"`
	Employe_ID      uint64 `gorm:"not null" json:"-"`
	Employe        	Employe   `gorm:"foreignkey:Employe_ID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"employe"`
	User_name				string `gorm:"type:varchar(255)" json:"user_name"`
	Password				string `gorm:"->;<-;not null" json:"-"`
	Is_Default			bool `gorm:"type:char(1)" json:"is_default"`
	Comments				string `gorm:"type:varchar(255)" json:"comments"`
	Is_Usable				bool `gorm:"type:char(1)" json:"is_usable"`
	Create_User			string `gorm:"type:varchar(255)" json:"create_user"`
	Create_Date			time.Time `gorm:"type:datetime" json:"create_date"`
	Update_User			string `gorm:"type:varchar(255)" json:"update_user"`
	User_Date				time.Time `gorm:"type:datetime" json:"user_date"`
	Token    				string  `gorm:"-" json:"token,omitempty"`
}
