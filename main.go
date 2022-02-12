package main

import (
	"github.com/alexsuriano/goRestAPI/database"
	"github.com/alexsuriano/goRestAPI/server"
)

func main() {

	database.StartDB()

	server := server.NewSever()

	server.Run()
}
