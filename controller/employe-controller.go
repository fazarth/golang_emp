package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/fazarth/golang_emp/dto"
	"github.com/fazarth/golang_emp/entity"
	"github.com/fazarth/golang_emp/helper"
	"github.com/fazarth/golang_emp/service"
	"github.com/gin-gonic/gin"
)

//EmployeController is a ...
type EmployeController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type employeController struct {
	employeService service.EmployeService
	jwtService  service.JWTService
}

//NewEmployeController create a new instances of BoookController
func NewEmployeController(employeServ service.EmployeService, jwtServ service.JWTService) EmployeController {
	return &employeController{
		employeService: employeServ,
		jwtService:  jwtServ,
	}
}

func (c *employeController) All(context *gin.Context) {
	var employes []entity.Employe = c.employeService.All()
	res := helper.BuildResponse(true, "OK", employes)
	context.JSON(http.StatusOK, res)
}

func (c *employeController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var employe entity.Employe = c.employeService.FindByID(id)
	if (employe == entity.Employe{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", employe)
		context.JSON(http.StatusOK, res)
	}
}

func (c *employeController) Insert(context *gin.Context) {
	var employeCreateDTO dto.EmployeCreateDTO
	errDTO := context.ShouldBind(&employeCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		employeID := c.getEmployeIDByToken(authHeader)
		convertedEmployeID, err := strconv.ParseUint(employeID, 10, 64)
		if err == nil {
			employeCreateDTO.Employe_ID = convertedEmployeID
		}
		result := c.employeService.Insert(employeCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *employeController) Update(context *gin.Context) {
	var employeUpdateDTO dto.EmployeUpdateDTO
	errDTO := context.ShouldBind(&employeUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	employeID := fmt.Sprintf("%v", claims["employe_id"])
	if c.employeService.IsAllowedToEdit(employeID, employeUpdateDTO.Employe_ID) {
		id, errID := strconv.ParseUint(employeID, 10, 64)
		if errID == nil {
			employeUpdateDTO.Employe_ID = id
		}
		result := c.employeService.Update(employeUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *employeController) Delete(context *gin.Context) {
	var employe entity.Employe
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	employe.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	employeID := fmt.Sprintf("%v", claims["employe_id"])
	if c.employeService.IsAllowedToEdit(employeID, employe.ID) {
		c.employeService.Delete(employe)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *employeController) getEmployeIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["employe_id"])
	return id
}
