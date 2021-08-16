package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main()  {
	
	// Instanciar echo
	e := echo.New()

	// RUta
	e.GET("/", func (c echo.Context) error  {
		return c.String(http.StatusOK, "Hello World")
	})
	e.Logger.Fatal(e.Start(":1323"))

}