package main

import (
	"github.com/hjunor/api-rest-golang.git/database"
	"github.com/hjunor/api-rest-golang.git/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()
	server.Run()
}
