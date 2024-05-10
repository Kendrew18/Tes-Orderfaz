package register

import (
	"Test-orderfaz/model/request"
	"Test-orderfaz/model/response"
	"Test-orderfaz/service/register"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.String(http.StatusUnauthorized, "Invalid JWT token")
		}

		token, err := jwt.ParseWithClaims(tokenString, &request.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("T3stOrd3rf4z10052024 s3cr3t K3ys"), nil
		})

		if err != nil {
			return c.String(http.StatusUnauthorized, "Invalid JWT token")
		}

		claims, ok := token.Claims.(*request.JWTClaims)
		if !ok {
			return c.String(http.StatusUnauthorized, "Invalid JWT token")
		}

		c.Set("user_id", claims.Uuid)

		fmt.Println(claims.Uuid)

		return next(c)
	}
}

func Register(c echo.Context) error {
	var Request request.Register_Request
	var result response.Response
	var err error

	err = c.Bind(&Request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if len(Request.Msisdn) < 10 || Request.Msisdn[:2] != "62" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid MSISDN"})
	}

	result, err = register.Register(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func Login(c echo.Context) error {
	var Request request.Login_Request
	var result response.Response
	var err error

	err = c.Bind(&Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err = register.Login(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func GetProfile(c echo.Context) error {
	var Request request.Get_Profile_Request
	var result response.Response
	var err error

	userID := c.Get("user_id").(string)

	fmt.Println(userID)

	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid JWT token"})
	}

	Request.Uuid = userID

	result, err = register.Get_Profile(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
