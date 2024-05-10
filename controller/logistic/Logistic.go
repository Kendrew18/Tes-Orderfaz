package logistic

import (
	"Test-orderfaz/model/request"
	"Test-orderfaz/model/response"
	"Test-orderfaz/service/logistic"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ReadLogistic(c echo.Context) error {
	var Request request.Read_Logistic_Request
	var result response.Response
	var err error

	userID := c.Get("user_id").(string)

	fmt.Println(userID)

	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid JWT token"})
	}

	Request.Uuid = userID

	result, err = logistic.Read_Logistic(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}

func ReadLogisticFilter(c echo.Context) error {
	var Request request.Read_Logistic_Filter_Request
	var result response.Response
	var err error

	err = c.Bind(&Request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	userID := c.Get("user_id").(string)

	fmt.Println(userID)

	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid JWT token"})
	}

	Request.Uuid = userID

	result, err = logistic.Read_Logistic_Filter(Request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(result.Status, result)
}
