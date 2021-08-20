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
	db             						*gorm.DB                  						= config.SetupDatabaseConnection()
	employeRepository 				repository.EmployeRepository 					= repository.NewEmployeRepository(db)
	userRepository 						repository.UserRepository 						= repository.NewUserRepository(db)
	user_employeRepository 		repository.User_EmployeRepository 		= repository.NewUser_EmployeRepository(db)
	jwtService     						service.JWTService        						= service.NewJWTService()
	employeService    				service.EmployeService       					= service.NewEmployeService(employeRepository)
	userService    						service.UserService       						= service.NewUserService(userRepository)
	user_employeService				service.User_EmployeService       		= service.NewUser_EmployeService(user_employeRepository)
	authService    						service.AuthService       						= service.NewAuthService(userRepository)
	authController 						controller.AuthController 						= controller.NewAuthController(authService, jwtService)
	employeController 				controller.EmployeController 					= controller.NewEmployeController(employeService, jwtService)
	userController 						controller.UserController 						= controller.NewUserController(userService, jwtService)
	user_employeController 		controller.User_EmployeController 		= controller.NewUser_EmployeController(user_employeService, jwtService)
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
		employeRoutes.GET("/", employeController.All)
		employeRoutes.POST("/", employeController.Insert)
		employeRoutes.GET("/:id", employeController.FindByID)
		employeRoutes.PUT("/:id", employeController.Update)
		employeRoutes.DELETE("/:id", employeController.Delete)
	}

	userRoutes := r.Group("api/users", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/", userController.All)
		userRoutes.POST("/", userController.Insert)
		userRoutes.GET("/:id", userController.FindByID)
		userRoutes.PUT("/:id", userController.Update)
		userRoutes.DELETE("/:id", userController.Delete)
	}

	user_employeRoutes := r.Group("api/user_employe", middleware.AuthorizeJWT(jwtService))
	{
		user_employeRoutes.GET("/", user_employeController.All)
		user_employeRoutes.POST("/", user_employeController.Insert)
		user_employeRoutes.GET("/:id", user_employeController.FindByID)
		user_employeRoutes.PUT("/:id", user_employeController.Update)
		user_employeRoutes.DELETE("/:id", user_employeController.Delete)
	}

	r.Run()
}
