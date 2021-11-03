package main

import "github.com/hjunor/api-rest-golang.git/server"

func main() {
	server := server.NewServer()
	server.Run()
}
