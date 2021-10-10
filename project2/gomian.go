package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"eamil" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()
	e.GET("/getUser", getUser)

	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	u := new(User)
	u.Name = "sao"
	u.Email = "111.com"
	return c.JSON(http.StatusOK, u)
}
