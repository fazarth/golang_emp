package repository

import (
	"github.com/fazarth/golang_emp/entity"
	"gorm.io/gorm"
)

//User_EmployeRepository is a ....
type User_EmployeRepository interface {
	InsertUser_Employe(b entity.User_Employe) entity.User_Employe
	UpdateUser_Employe(b entity.User_Employe) entity.User_Employe
	DeleteUser_Employe(b entity.User_Employe)
	AllUser_Employe() []entity.User_Employe
	FindUser_EmployeByID(user_employeID uint64) entity.User_Employe
	VerifyCredential(email string, password string) interface{}
}

type user_employeConnection struct {
	connection *gorm.DB
}

//NewUser_EmployeRepository creates an instance User_EmployeRepository
func NewUser_EmployeRepository(dbConn *gorm.DB) User_EmployeRepository {
	return &user_employeConnection{
		connection: dbConn,
	}
}

func (db *user_employeConnection) InsertUser_Employe(b entity.User_Employe) entity.User_Employe {
	db.connection.Save(&b)
	db.connection.Preload("User_Employe").Find(&b)
	return b
}

func (db *user_employeConnection) UpdateUser_Employe(b entity.User_Employe) entity.User_Employe {
	db.connection.Save(&b)
	db.connection.Preload("User_Employe").Find(&b)
	return b
}

func (db *user_employeConnection) VerifyCredential(email string, password string) interface{} {
	var user_employe entity.Employe
	res := db.connection.Where("email = ?", email).Take(&user_employe)
	if res.Error == nil {
		return user_employe
	}
	return nil
}

func (db *user_employeConnection) DeleteUser_Employe(b entity.User_Employe) {
	db.connection.Delete(&b)
}

func (db *user_employeConnection) FindUser_EmployeByID(user_employeID uint64) entity.User_Employe {
	var user_employe entity.User_Employe
	db.connection.Preload("User_Employe").Find(&user_employe, user_employeID)
	return user_employe
}

func (db *user_employeConnection) AllUser_Employe() []entity.User_Employe {
	var user_employes []entity.User_Employe
	db.connection.Preload("User_Employe").Find(&user_employes)
	return user_employes
}
