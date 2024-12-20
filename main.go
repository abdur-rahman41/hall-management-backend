package main

import (
	"github.com/abdur-rahman41/hall-management-backend/pkg/containers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	containers.Serve(e)
}
