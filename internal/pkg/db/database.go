package db

import (
	_ "database/sql"
	"github.com/Aibier/go-aml-service/internal/pkg/config"
	"github.com/Aibier/go-aml-service/internal/pkg/models/tasks"
	"github.com/Aibier/go-aml-service/internal/pkg/models/users"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var (
	DB  *gorm.DB
	err error
)

type Database struct {
	*gorm.DB
}

// SetupDB opens a database and saves the reference to `Database` struct.
func SetupDB() {
	var db = DB

	configuration := config.GetConfig()

	driver := configuration.Database.Driver
	database := configuration.Database.Dbname
	username := configuration.Database.Username
	password := configuration.Database.Password
	host := configuration.Database.Host
	port := configuration.Database.Port

	if driver == "sqlite" { // SQLITE
		db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		if err != nil {
			log.WithError(err).Println("db err: ", err)
		}
	} else if driver == "postgres" { // POSTGRES
		postgresInfo := "host="+host+" port="+port+" user="+username+" dbname="+database+"  sslmode=disable password="+password + "TimeZone=Asia/Shanghai"
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN: postgresInfo,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{})
		if err != nil {
			log.WithError(err).Println("db err: ", err)
		}
	} else if driver == "mysql" { // MYSQL
		dsn := username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.WithError(err).Println("db err: ", err)
		}
	}
	// Change this to true if you want to see SQL queries
	myDB, err := db.DB()
	if err != nil {
		log.WithError(err).Println("db err: ", err)
	}
	myDB.SetMaxIdleConns(configuration.Database.MaxIdleConns)
	myDB.SetMaxOpenConns(configuration.Database.MaxOpenConns)
	myDB.SetConnMaxLifetime(time.Duration(configuration.Database.MaxLifetime) * time.Second)
	DB = db
	migration()
}

// Auto migrate project models
func migration(){
	err := DB.AutoMigrate(&users.User{})
	if err != nil {
		log.WithError(err).Printf("failed migrate users")
	}
	err = DB.AutoMigrate(&users.UserRole{})
	if err != nil {
		log.WithError(err).Printf("failed migrate user roles")
	}
	err = DB.AutoMigrate(&tasks.Task{})
	if err != nil {
		log.WithError(err).Printf("failed migrate tasks")
	}
}

func GetDB() *gorm.DB {
	return DB
}
