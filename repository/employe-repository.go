package repository

import (
	"github.com/fazarth/golang_emp/entity"
	"gorm.io/gorm"
)

//EmployeRepository is a ....
type EmployeRepository interface {
	InsertEmploye(b entity.Employe) entity.Employe
	UpdateEmploye(b entity.Employe) entity.Employe
	DeleteEmploye(b entity.Employe)
	AllEmploye() []entity.Employe
	FindEmployeByID(employeID uint64) entity.Employe
	VerifyCredential(email string, password string) interface{}
}

type employeConnection struct {
	connection *gorm.DB
}

//NewEmployeRepository creates an instance EmployeRepository
func NewEmployeRepository(dbConn *gorm.DB) EmployeRepository {
	return &employeConnection{
		connection: dbConn,
	}
}

func (db *employeConnection) InsertEmploye(b entity.Employe) entity.Employe {
	db.connection.Save(&b)
	db.connection.Preload("Employe").Find(&b)
	return b
}

func (db *employeConnection) UpdateEmploye(b entity.Employe) entity.Employe {
	db.connection.Save(&b)
	db.connection.Preload("Employe").Find(&b)
	return b
}

func (db *employeConnection) VerifyCredential(email string, password string) interface{} {
	var employe entity.Employe
	res := db.connection.Where("email = ?", email).Take(&employe)
	if res.Error == nil {
		return employe
	}
	return nil
}

func (db *employeConnection) DeleteEmploye(b entity.Employe) {
	db.connection.Delete(&b)
}

func (db *employeConnection) FindEmployeByID(employeID uint64) entity.Employe {
	var employe entity.Employe
	db.connection.Preload("Employe").Find(&employe, employeID)
	return employe
}

func (db *employeConnection) AllEmploye() []entity.Employe {
	var employes []entity.Employe
	db.connection.Preload("Employe").Find(&employes)
	return employes
}
