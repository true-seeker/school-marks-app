package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"school-marks-app/api/config"
	errorHandler "school-marks-app/api/error"
)

type GormConnection struct {
	conn gorm.DB
}

// GetDBConnection подключение к БД
func GetDBConnection() GormConnection {
	var PostgresConnectionString = fmt.Sprintf("host=%s "+
		"user=%s "+
		"password=%s "+
		"dbname=%s "+
		"port=%s "+
		"sslmode=disable TimeZone=Asia/Yekaterinburg",
		config.GetConfig().Get("database.address"),
		config.GetConfig().Get("database.user"),
		config.GetConfig().Get("database.password"),
		config.GetConfig().Get("database.dbname"),
		config.GetConfig().Get("database.port"))

	db, err := gorm.Open(postgres.Open(PostgresConnectionString), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	errorHandler.FailOnError(err, "Failed to connect to DB")
	return GormConnection{conn: *db}
}
