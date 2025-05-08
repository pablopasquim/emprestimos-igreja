package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectionDataBase() {
	var dbConectionString string = "host=localhost user=root password=root dbname=root port=5432"

	DB, err = gorm.Open(postgres.Open(dbConectionString))

	if err != nil {
		fmt.Errorf("Erro ao conectar com a base de dados: %w", err)
	}
}
