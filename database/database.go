package database

import (
	"adopet/config"
	"adopet/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func GetDatabase() {
	env, err := config.LoadEnv()
	if err != nil {
		log.Printf("Falha ao carregar o env  = %s", err)
	}

	host := env.Host
	port := env.Port
	user := env.User
	pass := env.Password
	db := env.Database

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, db)
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Tutor{}, &models.Abrigo{}, &models.Pet{}, &models.Login{}, &models.User{})
}
