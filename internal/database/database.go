package database

import (
	"fmt"
	"os"
	"vatansoft-case/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DbConnection struct {
	db *gorm.DB
}

func NewDbConnection() *DbConnection {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	DB = db

	return &DbConnection{
		db: db,
	}
}

func (conn *DbConnection) CloseDbConnection() {
	sqlDB, err := conn.db.DB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get SQL database: %v\n", err)
		return
	}
	sqlDB.Close()
}

func (conn *DbConnection) AutoMigrate() error {
	if err := conn.db.AutoMigrate(&model.User{}); err != nil {
		return err
	}
	return nil
}
