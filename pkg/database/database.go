package database

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/RoberPlaza/rehabilitea-webapp/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connection stores the information of a connection
type Connection struct {
	Host      string
	User      string
	Password  string
	Schema    string
	Port      uint64
	EnableSSL bool
}

// Handler of the Database
var Handler *gorm.DB

// GetPostgresString returns the connection string for a Postgres database
func (dc *Connection) GetPostgresString() string {
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

// GetMySQLString returns the connection string for a MySQL database
func (dc *Connection) GetMySQLString() string {
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
func initPostgres(conn *Connection) {
	var err error
	Handler, err = gorm.Open(postgres.Open(conn.GetPostgresString()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

// initMySQL initializes a db handler with a mysql database
func initMySQL(conn *Connection) {
	var err error
	Handler, err = gorm.Open(mysql.Open(conn.GetMySQLString()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

// Migrate migrates all the models
func Migrate() (err error) {
	models := []interface{}{
		model.Profile{},
		model.Game{},
		model.Difficulty{},
		model.Event{},
		model.Progression{},
		model.Score{},
	}

	for _, model := range models {
		if err = Handler.AutoMigrate(model); err != nil {
			return err
		}
	}

	return err
}

// InitPostConn initializes the databse as postgress using a connection
func InitPostConn(conn *Connection) error {
	initPostgres(conn)
	return Migrate()
}

// InitEnv initializes the database with environment variables
func InitEnv() error {
	dbPort, err := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 32)
	if err != nil {
		return err
	}

	dbSSL, err := strconv.ParseBool(os.Getenv("DB_SSL"))
	if err != nil {
		return err
	}

	conn := &Connection{
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
		initMySQL(conn)
		break
	case "postgres":
	case "postgresql":
	default:
		initPostgres(conn)
		break
	}

	return Migrate()
}
