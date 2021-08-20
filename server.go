package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_api/config"
	"github.com/ydhnwb/golang_api/controller"
	"github.com/ydhnwb/golang_api/middleware"
	"github.com/ydhnwb/golang_api/repository"
	"github.com/ydhnwb/golang_api/service"
	"gorm.io/gorm"
)

var (
	db             		*gorm.DB                  		= config.SetupDatabaseConnection()
	employeRepository repository.EmployeRepository 	= repository.NewEmployeRepository(db)
	userRepository 		repository.UserRepository 		= repository.NewUserRepository(db)
	jwtService     		service.JWTService        		= service.NewJWTService()
	employeService    service.EmployeService       	= service.NewEmployeService(employeRepository)
	userService    		service.UserService       		= service.NewUserService(userRepository)
	authService    		service.AuthService       		= service.NewAuthService(userRepository)
	authController 		controller.AuthController 		= controller.NewAuthController(authService, jwtService)
	employeController controller.EmployeController 	= controller.NewEmployeController(employeService, jwtService)
	userController 		controller.UserController 		= controller.NewUserController(userService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	employeRoutes := r.Group("api/employe", middleware.AuthorizeJWT(jwtService))
	{
		employeRoutes.GET("/profile", employeController.Profile)
		employeRoutes.PUT("/profile", employeController.Update)
	}

	userRoutes := r.Group("api/users", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/", userController.All)
		userRoutes.POST("/", userController.Insert)
		userRoutes.GET("/:id", userController.FindByID)
		userRoutes.PUT("/:id", userController.Update)
		userRoutes.DELETE("/:id", userController.Delete)
	}

	r.Run()
}
