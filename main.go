package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os/exec"
	"fmt"
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
    command := "exec.sh"
    cmd := exec.Command("/bin/bash", command)
    _, err := cmd.Output()
    if err != nil {
        fmt.Printf("Execute Shell:%s failed with error:%s\n", command, err.Error())
        return err
    }
    return c.JSON(http.StatusOK, map[string]interface{}{"code":0,"message":"succeed"})
}
