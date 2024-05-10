package response

type Logistic_Response struct {
	Id               int    `json:"id"`
	Logistic_name    string `json:"logistic_name"`
	Amount           int    `json:"amount"`
	Destination_name string `json:"destination_name"`
	Origin_name      string `json:"origin_name"`
	Duration         string `json:"duration"`
}
