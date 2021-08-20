package repository

import (
	"github.com/fazarth/golang_emp/entity"
	"gorm.io/gorm"
)

//UserRepository is a ....
type UserRepository interface {
	InsertUser(b entity.User) entity.User
	UpdateUser(b entity.User) entity.User
	DeleteUser(b entity.User)
	AllUser() []entity.User
	FindUserByID(userID uint64) entity.User
	VerifyCredential(email string, password string) interface{}
}

type userConnection struct {
	connection *gorm.DB
}

//NewUserRepository creates an instance UserRepository
func NewUserRepository(dbConn *gorm.DB) UserRepository {
	return &userConnection{
		connection: dbConn,
	}
}

func (db *userConnection) InsertUser(b entity.User) entity.User {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *userConnection) UpdateUser(b entity.User) entity.User {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.Employe
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) DeleteUser(b entity.User) {
	db.connection.Delete(&b)
}

func (db *userConnection) FindUserByID(userID uint64) entity.User {
	var user entity.User
	db.connection.Preload("User").Find(&user, userID)
	return user
}

func (db *userConnection) AllUser() []entity.User {
	var users []entity.User
	db.connection.Preload("User").Find(&users)
	return users
}
