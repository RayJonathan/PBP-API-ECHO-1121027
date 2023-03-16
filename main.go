package main

import (
	"example.com/m/db"
	"example.com/m/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
