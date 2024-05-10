package response

type Login_Response struct {
	Token string `json:"token"`
}

type Get_Profile_Response struct {
	Uuid   string `json:"uuid"`
	Msisdn string `json:"msisdn"`
	Name   string `json:"name"`
}
