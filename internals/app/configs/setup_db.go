package configs

import (
	"log"

	"github.com/Erodot0/gym-memeber-management/internals/app/domains/enteties"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitializeSQLite() (*gorm.DB, error) {
	var err error
	if DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	}); err != nil {
		return nil, err
	}

	// Auto Migrate will create the tables
	if err = DB.AutoMigrate(
		&enteties.Member{},
		&enteties.Contacts{},
		&enteties.Address{},
		&enteties.Subscription{},
	); err != nil {
		log.Println(err)
		return nil, err
	}

	return DB, nil
}
