package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/entity"
	"github.com/ydhnwb/golang_api/repository"
)

//EmployeService is a contract.....
type EmployeService interface {
	Update(employe dto.EmployeUpdateDTO) entity.Employe
	Profile(employeID string) entity.Employe
}

type employeService struct {
	employeRepository repository.EmployeRepository
}

//NewEmployeService creates a new instance of EmployeService
func NewEmployeService(employeRepo repository.EmployeRepository) EmployeService {
	return &employeService{
		employeRepository: employeRepo,
	}
}

func (service *employeService) Update(employe dto.EmployeUpdateDTO) entity.Employe {
	employeToUpdate := entity.Employe{}
	err := smapping.FillStruct(&employeToUpdate, smapping.MapFields(&employe))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedEmploye := service.employeRepository.UpdateEmploye(employeToUpdate)
	return updatedEmploye
}

func (service *employeService) Profile(employeID string) entity.Employe {
	return service.employeRepository.ProfileEmploye(employeID)
}
