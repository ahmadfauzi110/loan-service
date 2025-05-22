package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	investmentRepository "github.com/ahmadfauzi110/loan-service/internal/infrastructure/repository/investment"
	loanRepository "github.com/ahmadfauzi110/loan-service/internal/infrastructure/repository/loan"
	userRepository "github.com/ahmadfauzi110/loan-service/internal/infrastructure/repository/user"
)

func SetupDatabase(dbConfig *DB) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.Database, dbConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(userRepository.Model{})
	db.AutoMigrate(loanRepository.Model{})
	db.AutoMigrate(investmentRepository.Model{})
	return db
}
