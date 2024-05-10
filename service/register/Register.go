package register

import (
	"Test-orderfaz/db"
	"Test-orderfaz/model/request"
	"Test-orderfaz/model/response"
	"Test-orderfaz/tools/decrypt"
	"Test-orderfaz/tools/encrypt"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func Register(Request request.Register_Request) (response.Response, error) {

	var res response.Response

	con := db.CreateConGorm()

	uuid_user := ""

	err := con.Table("user").Select("uuid").Where("msisdn=? || username=?", Request.Msisdn, Request.Username).Scan(&uuid_user)

	if err.Error != nil || uuid_user != "" {
		res.Status = http.StatusInternalServerError
		res.Message = "msidns dan atau username sudah ada"
		res.Data = Request
		return res, err.Error
	}

	uuid := uuid.NewString()

	Request.Uuid = uuid

	key := []byte("T3stOrd3rf4z10052024 s3cr3t K3ys")

	encryptedPassword, err2 := encrypt.EncryptPassword(Request.Password, key)
	if err2 != nil {
		fmt.Println(err2)
	}

	Request.Password = encryptedPassword

	err = con.Table("user").Select("uuid", "name", "msisdn", "username", "password").Create(&Request)

	if err.Error != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "msidns dan atau username sudah ada"
		res.Data = Request
		return res, err.Error
	}

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = map[string]int64{
			"rows": err.RowsAffected,
		}
	}

	return res, nil
}

func Login(Request request.Login_Request) (response.Response, error) {
	var res response.Response
	var data response.Login_Response

	con := db.CreateConGorm()

	password := ""

	key := []byte("T3stOrd3rf4z10052024 s3cr3t K3ys")

	err := con.Table("user").Select("password").Where("msisdn = ?", Request.Msisdn).Scan(&password)

	fmt.Println(err.Error)

	if err.Error != nil {
		fmt.Println("masuk")
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	fmt.Println(password)

	decryptedPassword, err2 := decrypt.DecryptPassword(password, key)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(decryptedPassword)

	if decryptedPassword == Request.Password {

		uuid := ""

		err = con.Table("user").Select("uuid").Where("msisdn = ?", Request.Msisdn).Scan(&uuid)

		if err.Error != nil {
			fmt.Println("masuk")
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, request.JWTClaims{
			Uuid: uuid,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			},
		})

		//_, err = token.SignedString([]byte("T3stOrd3rf4z10052024 s3cr3t K3ys"))

		Token, err2 := token.SignedString([]byte("T3stOrd3rf4z10052024 s3cr3t K3ys"))

		if err2 != nil {
			res.Status = http.StatusInternalServerError
			res.Message = "Token Bermasalah"
			res.Data = Request
			return res, err2
		}

		data.Token = Token

		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = data

	} else {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, nil
	}

	return res, nil
}

func Get_Profile(Request request.Get_Profile_Request) (response.Response, error) {
	var res response.Response

	var data response.Get_Profile_Response

	con := db.CreateConGorm()

	err := con.Table("user").Select("uuid", "name", "msisdn").Where("uuid = ?", Request.Uuid).Scan(&data)

	if err.Error != nil {
		fmt.Println("masuk")
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = data
	}

	return res, nil
}
