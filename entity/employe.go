package entity

//User represents users table in database
type Employe struct {
	ID				      	uint64  `gorm:"primary_key:auto_increment" json:"employe_id"`
	Employe_Name     	string  `gorm:"type:varchar(255)" json:"name"`
	Dept_ID  					uint64  `gorm:"type:int(11)" json:"dept_id"`
	Dept_Name  				string  `gorm:"type:varchar(255)" json:"dept_name"`
	Div_ID  					uint64  `gorm:"type:int(11)" json:"div_id"`
	Div_Name  				string  `gorm:"type:varchar(255)" json:"div_name"`
	Position  				string  `gorm:"type:varchar(255)" json:"position"`
	Address  					string  `gorm:"type:varchar(255)" json:"address"`
	Email  						string  `gorm:"type:varchar(255)" json:"email"`
	Telephone  				string  `gorm:"type:varchar(255)" json:"telephone"`
	Token    					string  `gorm:"-" json:"token,omitempty"`

}
