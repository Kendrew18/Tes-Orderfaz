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

##  Login

Link: localhost:38600/AUTH/get-profile

Method: GET

Header: 

keys | value
------------ | -------------
Authorization | <token>

Contoh Raw Json: -
