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

//User_EmployeController is a ...
type User_EmployeController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type user_employeController struct {
	user_employeService service.User_EmployeService
	jwtService  service.JWTService
}

//NewUser_EmployeController create a new instances of BoookController
func NewUser_EmployeController(user_employeServ service.User_EmployeService, jwtServ service.JWTService) User_EmployeController {
	return &user_employeController{
		user_employeService: user_employeServ,
		jwtService:  jwtServ,
	}
}

func (c *user_employeController) All(context *gin.Context) {
	var user_employes []entity.User_Employe = c.user_employeService.All()
	res := helper.BuildResponse(true, "OK", user_employes)
	context.JSON(http.StatusOK, res)
}

func (c *user_employeController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var user_employe entity.User_Employe = c.user_employeService.FindByID(id)
	if (user_employe == entity.User_Employe{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", user_employe)
		context.JSON(http.StatusOK, res)
	}
}

func (c *user_employeController) Insert(context *gin.Context) {
	var user_employeCreateDTO dto.User_EmployeCreateDTO
	errDTO := context.ShouldBind(&user_employeCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		user_employeID := c.getUser_EmployeIDByToken(authHeader)
		convertedUser_EmployeID, err := strconv.ParseUint(user_employeID, 10, 64)
		if err == nil {
			user_employeCreateDTO.ID = convertedUser_EmployeID
		}
		result := c.user_employeService.Insert(user_employeCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *user_employeController) Update(context *gin.Context) {
	var user_employeUpdateDTO dto.User_EmployeUpdateDTO
	errDTO := context.ShouldBind(&user_employeUpdateDTO)
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
	user_employeID := fmt.Sprintf("%v", claims["user_employe_id"])
	if c.user_employeService.IsAllowedToEdit(user_employeID, user_employeUpdateDTO.ID) {
		id, errID := strconv.ParseUint(user_employeID, 10, 64)
		if errID == nil {
			user_employeUpdateDTO.ID = id
		}
		result := c.user_employeService.Update(user_employeUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *user_employeController) Delete(context *gin.Context) {
	var user_employe entity.User_Employe
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	user_employe.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	user_employeID := fmt.Sprintf("%v", claims["user_employe_id"])
	if c.user_employeService.IsAllowedToEdit(user_employeID, user_employe.ID) {
		c.user_employeService.Delete(user_employe)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *user_employeController) getUser_EmployeIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_employe_id"])
	return id
}
