package repository

import (
	"log"

	"github.com/ydhnwb/golang_api/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//EmployeRepository is contract what employRepository can do to db
type EmployeRepository interface {
	InsertEmploye(employ entity.Employe) entity.Employe
	UpdateEmploye(employ entity.Employe) entity.Employe
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.Employe
	ProfileEmploye(employID string) entity.Employe
}

type employConnection struct {
	connection *gorm.DB
}

//NewEmployeRepository is creates a new instance of EmployeRepository
func NewEmployeRepository(db *gorm.DB) EmployeRepository {
	return &employConnection{
		connection: db,
	}
}

func (db *employConnection) InsertEmploye(employ entity.Employe) entity.Employe {
	db.connection.Save(&employ)
	return employ
}

func (db *employConnection) UpdateEmploye(employ entity.Employe) entity.Employe {
	db.connection.Save(&employ)
	return employ
}

func (db *employConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var employ entity.Employe
	return db.connection.Where("email = ?", email).Take(&employ)
}

func (db *employConnection) FindByEmail(email string) entity.Employe {
	var employ entity.Employe
	db.connection.Where("email = ?", email).Take(&employ)
	return employ
}

func (db *employConnection) ProfileEmploye(employID string) entity.Employe {
	var employ entity.Employe
	db.connection.Preload("Books").Preload("Books.Employe").Find(&employ, employID)
	return employ
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
