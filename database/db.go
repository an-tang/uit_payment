package database

import (
	"fmt"
	"log"

	"uit_payment/lib/env"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

type database struct{}

var PG database

func (this *database) GetInstance() *gorm.DB {
	if db == nil {
		db = this.initDatabase()
		log.Printf("Database connected")
	}
	return db
}

func (this *database) initDatabase() *gorm.DB {
	connectionStr := getConnectionString()
	db, err := gorm.Open("postgres", connectionStr)
	if err != nil {
		panic(err.Error())
	}

	db.DB()
	db.DB().Ping()
	db.DB().SetMaxIdleConns(env.GetMaxOpenConns())
	db.DB().SetMaxOpenConns(env.GetMaxIdleConns())
	db.SingularTable(false)
	db.LogMode(true)
	return db
}

func getConnectionString() string {
	host := env.GetDBHost()
	userName := env.GetDBUserName()
	password := env.GetDBPassword()
	sslMode := env.GetDBSSLMode()
	port := env.GetDBPort()
	databaseName := env.GetDBName()

	connectionStringTemplate := "host=%s port=%s sslmode=%s user=%s password='%s' dbname=%s "
	return fmt.Sprintf(connectionStringTemplate,
		host, port, sslMode,
		userName, password, databaseName)
}
