package entity

type User_Employe struct {
	ID           uint64  `gorm:"primary_key:auto_increment" json:"user_employes_id"`
	Employe_ID   uint64  `gorm:"not null" json:"-"`
	Employe_Name Employe `gorm:"foreignkey:Employe_ID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"employe_name"`
	User_Grup_ID uint64  `gorm:"type:int(11)" json:"user_grup_id"`
	User_Grup    string  `gorm:"type:varchar(255)" json:"user_grup"`
}
