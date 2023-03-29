package main

import (
	"adopet/database"
	"adopet/server"
)

func main() {
	database.GetDatabase()

	s := server.NewServer()
	s.Run()
}
