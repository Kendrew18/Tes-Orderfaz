package routes

import (
	"Test-orderfaz/controller/logistic"
	"Test-orderfaz/controller/register"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Project-NDL")
	})

	AUTH := e.Group("/AUTH")
	LOG := e.Group("/LOG")

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	//AUTH
	AUTH.POST("/register", register.Register)
	AUTH.POST("/login", register.Login)
	AUTH.GET("/get-profile", register.GetProfile, register.ValidateJWT)

	//Logistic
	LOG.GET("/logistic", logistic.ReadLogistic, register.ValidateJWT)
	LOG.POST("/logistic-filter", logistic.ReadLogisticFilter, register.ValidateJWT)

	return e
}
