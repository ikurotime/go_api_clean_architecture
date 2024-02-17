package main

import (
	"ikurotime/go_api_hexagonal_architecture/src/domain"
	"ikurotime/go_api_hexagonal_architecture/src/infrastructure"
	"log"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
    db  *gorm.DB
    err error
    dsn = "root:password@tcp(127.0.0.1:3306)/go_sample?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
    dbinit()
    e := echo.New()
    infrastructure.Init(e)
    e.Logger.Fatal(e.Start(":1323"))
}

func dbinit() {
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }
    db.Migrator().CreateTable(domain.User{})
}
