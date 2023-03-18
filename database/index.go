package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() error {
	var err error
	dsn := fmt.Sprintf("%s&parseTime=true", os.Getenv("DSN"))

	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})

	if err == nil {
		fmt.Println(("Successfully connected to PlanetScale MySql"))
	}

	return err
}
