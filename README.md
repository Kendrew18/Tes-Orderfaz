# How TO Use API 
__________
##  Register

Link: localhost:38600/AUTH/register

Method: POST

Header: -

Contoh Raw Json:
```
{
  "name" : "Alex",
  "msisdn" : "6281332226828",
  "username" : "densu",
  "password" : "denz123"
}
```

##  Login

Link: localhost:38600/AUTH/login

Method: POST

Header: -

Contoh Raw Json:
```
{
  "msisdn":"6281332226828",
  "password":"denz123"
}
```

##  Get Profile

Link: localhost:38600/AUTH/get-profile

Method: GET

Header: 

keys | value
------------ | -------------
Authorization | value of jwt token from login

Contoh Raw Json: -

##  Logistic

Link: localhost:38600/LOG/logistic

Method: GET

Header: 

keys | value
------------ | -------------
Authorization | value of jwt token from login

Contoh Raw Json: -

##  Logistic_flter

Link: localhost:38600/LOG/logistic-filter

Method: POST

Header: 

keys | value
------------ | -------------
Authorization | value of jwt token from login

Contoh Raw Json: 
```
{
  "origin_name":"Bandung",
  "destination_name":"Surabaya"
}
```
--------
NB: saya menyertakan Export sql yang saya gunakan pada folder sql-database

##  mohon maaf karena goswagger yang terinstal tidak dapat di init. Saya memberikan dokumentasi berupa readme.md. Atas perhatiannya saya ucapkan terima kasih.
