package main

import (
	"adopet/database"
	"adopet/server"
)

// @title AdoPet
// @version 0.0.1
// @description API para o aplicativo AdoPet

// @contact.name API AdoPet
// @contact.url https://www.google.com/
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080/

func main() {
	database.GetDatabase()

	s := server.NewServer()
	s.Run()
}
