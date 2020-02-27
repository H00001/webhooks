package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os/exec"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/merge", merge)

	// Start server
	e.Logger.Fatal(e.Start(":5654"))
}

// Handler
func merge(c echo.Context) error {
	exec.Command("bash","exec.sh")
	return c.JSON(http.StatusOK, map[string]interface{}{"code":0,"message":"succeed"})
}