package containers

import (
	"fmt"
	"log"

	"github.com/abdur-rahman41/hall-management-backend/pkg/config"
	"github.com/abdur-rahman41/hall-management-backend/pkg/connection"
	"github.com/abdur-rahman41/hall-management-backend/pkg/controllers"
	"github.com/abdur-rahman41/hall-management-backend/pkg/repositories"
	"github.com/abdur-rahman41/hall-management-backend/pkg/routes"
	"github.com/abdur-rahman41/hall-management-backend/pkg/services"

	"github.com/labstack/echo/v4"
)

func Serve(e *echo.Echo) {

	// Config initializations
	config.SetConfig()

	//Db initialization
	connection.Connect()
	db := connection.GetDB()
	//auth  module initialization
	authRepository := repositories.AuthDBInstance(db)
	authService := services.AuthServiceInstance(authRepository)
	authController := controllers.NewAuthController(authService)
	authRoutes := routes.NewAuthRoutes(e, authController)
	authRoutes.InitAuthRoutes()

	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))

}
