package database

import (
	"fmt"

	model "skypia/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB
var MyConn *gorm.DB

func Connect() {
	DatabasePostgre()
	MigratePostgre()
	DatabaseMySQL()
	MigrateMySQL()

}

func DatabasePostgre() {
	dsn := "host=postgres user=postgres password=393406 dbname=skypia port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	Conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func MigratePostgre() {
	fmt.Println("Migrating...")
	Conn.AutoMigrate(
		&model.User{},
	)
}

func DatabaseMySQL() {
	dsn := "user:password@tcp(mysql:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	MyConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func MigrateMySQL() {
	fmt.Println("Migrating...")
	MyConn.AutoMigrate(
		&model.User{},
	)
}
