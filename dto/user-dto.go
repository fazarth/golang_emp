package dto

//UserUpdateDTO is used by client when PUT update profile
type UserUpdateDTO struct {
	User_ID         uint64 `json:"user_id" form:"user_id"`
	Employe_ID      uint64 `json:"employe_id" form:"employe_id"`
	User_name				string `json:"user_name" form:"user_name"`
	Password				string `json:"password" form:"password"`
	Is_Default			bool `json:"is_default" form:"is_default"`
	Comments				string `json:"comments" form:"comments"`
	Is_Usable				bool `json:"is_usable" form:"is_usable"`
	Create_User			string `gorm:"type:varchar(255)" json:"create_user"`
	Create_Date			string `gorm:"type:datetime" json:"create_date"`
	Update_User			string `gorm:"type:varchar(255)" json:"update_user"`
	User_Date				string `gorm:"type:datetime" json:"user_date"`
	Token    				string  `gorm:"-" json:"token,omitempty"`
}

type UserCreateDTO struct {
	Employe_ID      uint64 `json:"employe_id" form:"employe_id"`
	User_name				string `json:"user_name" form:"user_name"`
	Password				string `json:"password" form:"password"`
	Is_Default			bool `json:"is_default" form:"is_default"`
	Comments				string `json:"comments" form:"comments"`
	Is_Usable				bool `json:"is_usable" form:"is_usable"`
	Create_User			string `gorm:"type:varchar(255)" json:"create_user"`
	Create_Date			string `gorm:"type:datetime" json:"create_date"`
	Update_User			string `gorm:"type:varchar(255)" json:"update_user"`
	User_Date				string `gorm:"type:datetime" json:"user_date"`
	Token    				string  `gorm:"-" json:"token,omitempty"`
}
