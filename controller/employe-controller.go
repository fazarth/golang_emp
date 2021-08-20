package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/helper"
	"github.com/ydhnwb/golang_api/service"
)

//EmployeController is a ....
type EmployeController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type employeController struct {
	employeService service.EmployeService
	jwtService  service.JWTService
}

//NewEmployeController is creating anew instance of EmployeControlller
func NewEmployeController(employeService service.EmployeService, jwtService service.JWTService) EmployeController {
	return &employeController{
		employeService: employeService,
		jwtService:  jwtService,
	}
}

func (c *employeController) Update(context *gin.Context) {
	var employeUpdateDTO dto.EmployeUpdateDTO
	errDTO := context.ShouldBind(&employeUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["employe_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	employeUpdateDTO.Employe_ID = id
	u := c.employeService.Update(employeUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

func (c *employeController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["employe_id"])
	employe := c.employeService.Profile(id)
	res := helper.BuildResponse(true, "OK", employe)
	context.JSON(http.StatusOK, res)

}
