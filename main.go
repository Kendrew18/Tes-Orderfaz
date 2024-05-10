package main

import (
	"Test-orderfaz/db"
	"Test-orderfaz/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":38600"))
}
