package dto

//BookUpdateDTO is a model that client use when updating a book
type EmployeUpdateDTO struct {
	Employe_ID  			int32 	`json:"employe_id" form:"id" binding:"required"`
	Employe_Name     	string  `json:"employe_name" form:"employe_name" binding:"required"`
	Dept_ID  					int32  	`json:"dept_id" form:"dept_id" binding:"required"`
	Dept_Name  				string  `json:"dept_name" form:"dept_name" binding:"required"`
	Div_ID  					int32  	`json:"div_id" form:"div_id" binding:"required"`
	Div_Name  				string  `json:"div_name" form:"div_name" binding:"required"`
	Position  				string 	`json:"position" form:"position" binding:"required"`
	Address  					string  `json:"address" form:"address" binding:"required"`
	Email  						string  `json:"adderess" form:"adderess" binding:"required"`
	Telephone  				string  `json:"telephone" form:"telephone" binding:"required"`
}

//EmployeCreateDTO is is a model that clinet use when create a new book
type EmployeCreateDTO struct {
	Employe_Name     	string  `json:"employe_name" form:"employe_name" binding:"required"`
	Dept_ID  					int32  	`json:"dept_id" form:"dept_id" binding:"required"`
	Dept_Name  				string  `json:"dept_name" form:"dept_name" binding:"required"`
	Div_ID  					int32  	`json:"div_id" form:"div_id" binding:"required"`
	Div_Name  				string  `json:"div_name" form:"div_name" binding:"required"`
	Position  				string 	`json:"position" form:"position" binding:"required"`
	Address  					string  `json:"address" form:"address" binding:"required"`
	Email  						string  `json:"adderess" form:"adderess" binding:"required"`
	Telephone  				string  `json:"telephone" form:"telephone" binding:"required"`
}
