package configs

import (
	"log"

	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitializeSQLite() (*gorm.DB, error) {
	log.Println("Connecting to SQLite database")
	var err error
	if DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	}); err != nil {
		return nil, err
	}

	// Auto Migrate will create the tables
	if err = DB.AutoMigrate(
		&entities.User{},
		&entities.Member{},
		&entities.Contacts{},
		&entities.Address{},
		&entities.Subscription{},
		&entities.Roles{},
		&entities.Permissions{},
	); err != nil {
		log.Println(err)
		return nil, err
	}

	return DB, nil
}
