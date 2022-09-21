package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	port := os.Args[1]

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("server %s", port))
	})

	if err := app.Start(":" + port); err != nil {
		panic(err)
	}
}
