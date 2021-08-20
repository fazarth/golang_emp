package service

import (
	"fmt"
	"log"

	"github.com/fazarth/golang_emp/dto"
	"github.com/fazarth/golang_emp/entity"
	"github.com/fazarth/golang_emp/repository"
	"github.com/mashingan/smapping"
)

//User_EmployeService is a ....
type User_EmployeService interface {
	Insert(b dto.User_EmployeCreateDTO) entity.User_Employe
	Update(b dto.User_EmployeUpdateDTO) entity.User_Employe
	Delete(b entity.User_Employe)
	All() []entity.User_Employe
	FindByID(user_employeID uint64) entity.User_Employe
	IsAllowedToEdit(employeID string, user_employeID uint64) bool
}

type user_employeService struct {
	user_employeRepository repository.User_EmployeRepository
}

//NewUser_EmployeService .....
func NewUser_EmployeService(user_employeRepo repository.User_EmployeRepository) User_EmployeService {
	return &user_employeService{
		user_employeRepository: user_employeRepo,
	}
}

func (service *user_employeService) Insert(b dto.User_EmployeCreateDTO) entity.User_Employe {
	user_employe := entity.User_Employe{}
	err := smapping.FillStruct(&user_employe, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.user_employeRepository.InsertUser_Employe(user_employe)
	return res
}

func (service *user_employeService) Update(b dto.User_EmployeUpdateDTO) entity.User_Employe {
	user_employe := entity.User_Employe{}
	err := smapping.FillStruct(&user_employe, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.user_employeRepository.UpdateUser_Employe(user_employe)
	return res
}

func (service *user_employeService) Delete(b entity.User_Employe) {
	service.user_employeRepository.DeleteUser_Employe(b)
}

func (service *user_employeService) All() []entity.User_Employe {
	return service.user_employeRepository.AllUser_Employe()
}

func (service *user_employeService) FindByID(user_employeID uint64) entity.User_Employe {
	return service.user_employeRepository.FindUser_EmployeByID(user_employeID)
}

func (service *user_employeService) IsAllowedToEdit(employeID string, user_employeID uint64) bool {
	b := service.user_employeRepository.FindUser_EmployeByID(user_employeID)
	id := fmt.Sprintf("%v", b.ID)
	return employeID == id
}
