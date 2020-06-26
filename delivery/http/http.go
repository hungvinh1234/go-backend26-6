package http

import (
	"go-template/delivery/http/account"
	"go-template/delivery/http/authorization"
	"go-template/delivery/http/hometown"
	"go-template/delivery/http/university"
	"go-template/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewHTTPHandler .
func NewHTTPHandler(ucase *usecase.Usecase) *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Static("/static", "static/images")

	//have to except SignIn va SignUp

	// j := e.Group("/account")

	account.Init(e.Group("/account"), ucase.Account)
	authorization.Init(e.Group(""), ucase.Account)
	hometown.Init(e.Group(""), ucase.Hometown)
	university.Init(e.Group(""), ucase.University)

	return e
}
