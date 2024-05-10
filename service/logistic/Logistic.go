package logistic

import (
	"Test-orderfaz/db"
	"Test-orderfaz/model/request"
	"Test-orderfaz/model/response"
	"net/http"
)

func Read_Logistic(Request request.Read_Logistic_Request) (response.Response, error) {
	var res response.Response
	var arr_invent []response.Logistic_Response

	con := db.CreateConGorm()

	err := con.Table("logistic").Select("id", "logistic_name", "amount", "destination_name", "origin_name", "duration").Where("uuid_user = ?", Request.Uuid).Order("id ASC").Scan(&arr_invent)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		return res, err.Error
	}

	if arr_invent == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_invent
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_invent
	}

	return res, nil
}

func Read_Logistic_Filter(Request request.Read_Logistic_Filter_Request) (response.Response, error) {
	var res response.Response
	var arr_data []response.Logistic_Response

	con := db.CreateConGorm()

	statement := "uuid_user = '" + Request.Uuid + "'"

	if Request.Origin_name != "" {
		statement += " && origin_name = '" + Request.Origin_name + "'"
	}

	if Request.Destination_name != "" {
		statement += " && destination_name = '" + Request.Destination_name + "'"
	}

	err := con.Table("logistic").Select("id", "logistic_name", "amount", "destination_name", "origin_name", "duration").Where(statement).Order("id ASC").Scan(&arr_data)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		return res, err.Error
	}

	if arr_data == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_data
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_data
	}

	return res, nil
}
