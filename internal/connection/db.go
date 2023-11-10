package connection

import (
	"fmt"
	"log"

	"github.com/mokhlesurr031/event-management-svc/domain"
	"github.com/mokhlesurr031/event-management-svc/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB holds the database instance
var db *gorm.DB

// Ping tests if db connection is alive
func Ping() error {
	return db.Exec("SELECT 'DBD::Pg ping test';").Error
}

// Connect sets the db client of database using configuration cfg
func Connect(cfg *config.Database) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	// Open a database connection using gorm ORM with the MySQL driver
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db = d

	makeMigration := Migrate()

	if makeMigration {
		if err := db.AutoMigrate(&domain.Reservation{}, &domain.Event{}, &domain.Workshop{}); err != nil {
			log.Fatalln(err)
		}
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	log.Println(sqlDB)
	return nil
}

// DefaultDB returns default db
func DefaultDB() *gorm.DB {
	return db.Debug()
}

// ConnectDB sets the db client of database using default configuration file
func ConnectDB() error {
	cfg := config.DB()
	return Connect(cfg)
}
