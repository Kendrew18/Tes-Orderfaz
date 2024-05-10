package request

import (
	"github.com/golang-jwt/jwt"
)

// swagger:model User
type Register_Request struct {
	Uuid     string `json:"uuid"`
	Msisdn   string `json:"msisdn"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// swagger:model User
type Login_Request struct {
	Msisdn   string `json:"msisdn"`
	Password string `json:"password"`
}

type JWTClaims struct {
	Uuid string `json:"id"`
	jwt.StandardClaims
}

// swagger:model User
type Get_Profile_Request struct {
	Uuid string `json:"uuid"`
}
