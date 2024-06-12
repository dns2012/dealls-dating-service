package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func gormDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			ENV("DB_USER", "default"),
			ENV("DB_PASS", "secret"),
			ENV("DB_HOST", "mysql"),
			ENV("DB_PORT", "3306"),
			ENV("DB_NAME", "default"),
		)
}

func gormConfig() *gorm.Config {
	gormLog :=  logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:              time.Second,
			LogLevel:                   logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			Colorful:                  false,
		},
	)
	return &gorm.Config{Logger: gormLog}
}

func GormDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(gormDSN()), gormConfig())
	if err != nil {
		log.Fatalf("Failed to open gorm connection: %v", err)
	}
	return db
}
