package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/guatom999/TicketShop/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	Database interface {
		GetDb() *gorm.DB
	}

	postgresDb struct {
		Db *gorm.DB
	}
)

func NewPostgresDatabase(cfg *config.Config) Database {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Db.Host,
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.DbName,
		cfg.Db.Port,
		cfg.Db.SslMode,
		cfg.Db.TimeZone,
	)

	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	return &postgresDb{Db: db}

}

func (p *postgresDb) GetDb() *gorm.DB {
	return p.Db
}
