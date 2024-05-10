package request

// swagger:model Logistic
type Read_Logistic_Request struct {
	Uuid string `json:"uuid"`
}

// swagger:model Logistic
type Read_Logistic_Filter_Request struct {
	Uuid             string `json:"uuid"`
	Origin_name      string `json:"origin_name"`
	Destination_name string `json:"destination_name"`
}
