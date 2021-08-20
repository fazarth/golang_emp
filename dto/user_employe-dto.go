package dto

//UserUpdateDTO is used by client when PUT update profile
type User_EmployeUpdateDTO struct {
	ID         		uint64 `json:"user_id" form:"user_id"`
	Employe_Name  string `json:"employe_name" form:"employe_name"`
	User_Grup_ID	uint64 `json:"user_grup_id" form:"user_grup_id"`
	User_Grup			string `json:"user_grup" form:"user_grup"`
}

type User_EmployeCreateDTO struct {
	Employe_Name  uint64 `json:"employe_name" form:"employe_name"`
	User_Grup_ID	uint64 `json:"user_grup_id" form:"user_grup_id"`
	User_Grup			string `json:"user_grup" form:"user_grup"`
}
