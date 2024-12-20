package routes

import (
	"github.com/abdur-rahman41/hall-management-backend/pkg/controllers"
	"github.com/labstack/echo/v4"
)

// AuthRoutes stores controller and echo instance for authentication.
type AuthRoutes struct {
	echo    *echo.Echo
	authCtr controllers.AuthController
}

// NewAuthRoutes returns a new instance of the AuthRoutes struct.
func NewAuthRoutes(echo *echo.Echo, authCtr controllers.AuthController) *AuthRoutes {
	return &AuthRoutes{
		echo:    echo,
		authCtr: authCtr,
	}
}

// InitAuthRoutes initializes the authentication routes.
func (routes *AuthRoutes) InitAuthRoutes() {
	e := routes.echo
	v1 := e.Group("/v1")
	v1.POST("/auth/sign-up", routes.authCtr.Signup)
	//v1.POST("/auth/login", routes.authCtr.Login)
}
