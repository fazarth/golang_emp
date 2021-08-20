package service

import (
	"fmt"
	"log"

	"github.com/fazarth/golang_emp/dto"
	"github.com/fazarth/golang_emp/entity"
	"github.com/fazarth/golang_emp/repository"
	"github.com/mashingan/smapping"
)

//UserService is a ....
type UserService interface {
	Insert(b dto.UserCreateDTO) entity.User
	Update(b dto.UserUpdateDTO) entity.User
	Delete(b entity.User)
	All() []entity.User
	FindByID(userID uint64) entity.User
	IsAllowedToEdit(employeID string, userID uint64) bool
}

type userService struct {
	userRepository repository.UserRepository
}

//NewUserService .....
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Insert(b dto.UserCreateDTO) entity.User {
	user := entity.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.userRepository.InsertUser(user)
	return res
}

func (service *userService) Update(b dto.UserUpdateDTO) entity.User {
	user := entity.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.userRepository.UpdateUser(user)
	return res
}

func (service *userService) Delete(b entity.User) {
	service.userRepository.DeleteUser(b)
}

func (service *userService) All() []entity.User {
	return service.userRepository.AllUser()
}

func (service *userService) FindByID(userID uint64) entity.User {
	return service.userRepository.FindUserByID(userID)
}

func (service *userService) IsAllowedToEdit(employeID string, userID uint64) bool {
	b := service.userRepository.FindUserByID(userID)
	id := fmt.Sprintf("%v", b.ID)
	return employeID == id
}
