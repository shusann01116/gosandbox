package main

import (
	"database/sql"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	user := User{Name: "Tarou", Age: 19, Birthday: time.Now()}

	db.Create(&user)

	users := []*User{
		{Name: "Jiro", Age: 20, Birthday: time.Now()},
		{Name: "Yome", Age: 22, Birthday: time.Now()},
	}

	db.Create(users)
}

type User struct {
	gorm.Model
	Name         string
	Email        *string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}
