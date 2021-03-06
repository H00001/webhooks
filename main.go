package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/merge", merge)
	e.POST("/merge", merge)

	// Start server
	path := ConfigPath()
	fmt.Println("path: ",path)
	Init(path)
	e.Logger.Fatal(e.Start(strings.Trim(string(confs["port"]),"\"")))
}

// Handler
func merge(c echo.Context) error {
	go func(){
		command := "exec.sh"
		cmd := exec.Command("/bin/bash", command)
		_, err := cmd.Output()
		if err != nil {
			fmt.Printf("Execute Shell:%s failed with error:%s\n", command, err.Error())
		}
	}()
        return c.JSON(http.StatusOK, map[string]interface{}{"code":0,"message":"succeed"})
}
