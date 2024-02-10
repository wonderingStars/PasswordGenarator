package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	// enter the size of the string the 2nd parm is to for letters ,numbers and symobla chose 1,2,3
	// custom the number if chars in custom has to be the same or more than the lenth

	shuffle(10, 4, "")

	e := echo.New()

	e.POST("/genaratePassowrd", getPassword)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
