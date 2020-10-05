package common

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DatabaseConnection stores the information of a connection
type DatabaseConnection struct {
	Host      string
	User      string
	Password  string
	Schema    string
	Port      uint64
	EnableSSL bool
}

// Database abstracts a database
type Database struct {
	*gorm.DB
}

// Handler of the Database
var handler Database

// GetDatabase returns the database
func GetDatabase() *Database {
	return &handler
}

// GetPostgresString returns the databaseconnection string for a Postgres database
func (dc *DatabaseConnection) GetPostgresString() string {
	baseFormat := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
	if dc.EnableSSL {
		baseFormat = "host=%s user=%s password=%s dbname=%s port=%d sslmode=enable"
	}

	return fmt.Sprintf(
		baseFormat,
		dc.Host,
		dc.User,
		dc.Password,
		dc.Schema,
		dc.Port,
	)
}

// GetMySQLString returns the databaseconnection string for a MySQL database
func (dc *DatabaseConnection) GetMySQLString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
		dc.User,
		dc.Password,
		dc.Host,
		dc.Port,
		dc.Schema,
	)
}

// initPostgres initializes a db handler with a pg database
func (database *Database) initPostgres(conn *DatabaseConnection) (err error) {
	database.DB, err = gorm.Open(postgres.Open(conn.GetPostgresString()), &gorm.Config{})
	return err
}

// initMySQL initializes a db handler with a mysql database
func (database *Database) initMySQL(conn *DatabaseConnection) (err error) {
	database.DB, err = gorm.Open(mysql.Open(conn.GetMySQLString()), &gorm.Config{})
	return err
}

// InitPostConn initializes the databse as postgress using a connection
func (database *Database) InitPostConn(conn *DatabaseConnection) error {
	return database.initPostgres(conn)
}

// InitEnv initializes the database with environment variables
func (database *Database) InitEnv() error {
	dbPort, err := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 32)
	if err != nil {
		return err
	}

	dbSSL, err := strconv.ParseBool(os.Getenv("DB_SSL"))
	if err != nil {
		return err
	}

	conn := &DatabaseConnection{
		Host:      os.Getenv("DB_HOST"),
		Password:  os.Getenv("DB_PASS"),
		Schema:    os.Getenv("DB_SCHEMA"),
		User:      os.Getenv("DB_USER"),
		Port:      dbPort,
		EnableSSL: dbSSL,
	}

	switch strings.ToLower(os.Getenv("DB_ENGINE")) {
	case "mysql":
	case "mariadb":
		database.initMySQL(conn)
		break
	case "postgres":
	case "postgresql":
	default:
		database.initPostgres(conn)
		break
	}

	return nil
}
