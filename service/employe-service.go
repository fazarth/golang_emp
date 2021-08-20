package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/entity"
	"github.com/ydhnwb/golang_api/repository"
)

//EmployeService is a ....
type EmployeService interface {
	Insert(b dto.EmployeCreateDTO) entity.Employe
	Update(b dto.EmployeUpdateDTO) entity.Employe
	Delete(b entity.Employe)
	All() []entity.Employe
	FindByID(employeID uint64) entity.Employe
	IsAllowedToEdit(employeID string, userID uint64) bool
}

type employeService struct {
	employeRepository repository.EmployeRepository
}

//NewEmployeService .....
func NewEmployeService(employeRepo repository.EmployeRepository) EmployeService {
	return &employeService{
		employeRepository: employeRepo,
	}
}

func (service *employeService) Insert(b dto.EmployeCreateDTO) entity.Employe {
	employe := entity.Employe{}
	err := smapping.FillStruct(&employe, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.employeRepository.InsertEmploye(employe)
	return res
}

func (service *employeService) Update(b dto.EmployeUpdateDTO) entity.Employe {
	employe := entity.Employe{}
	err := smapping.FillStruct(&employe, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.employeRepository.UpdateEmploye(employe)
	return res
}

func (service *employeService) Delete(b entity.Employe) {
	service.employeRepository.DeleteEmploye(b)
}

func (service *employeService) All() []entity.Employe {
	return service.employeRepository.AllEmploye()
}

func (service *employeService) FindByID(employeID uint64) entity.Employe {
	return service.employeRepository.FindEmployeByID(employeID)
}

func (service *employeService) IsAllowedToEdit(employeID string, userID uint64) bool {
	b := service.employeRepository.FindEmployeByID(userID)
	id := fmt.Sprintf("%v", b.ID)
	return employeID == id
}
