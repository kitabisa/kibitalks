package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kitabisa/kibitalk/config"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
	"time"
)

type MySQLDb struct {
	Client *sql.DB
}

var MySqlDB IMySQL

type MySQLConfig struct {
	Host                 string
	Port                 int
	DriverName           string
	Username             string
	Password             string
	DBName               string
	AdditionalParameters string
	MaxOpenConns         int
	MaxIdleConns         int
	ConnMaxLifetime      time.Duration
}

func (m MySQLConfig) InitMysqlDB() (db *sqlx.DB, err error) {
	db, err = sqlx.Open(m.DriverName, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", m.Username, m.Password, m.Host, m.Port, m.DBName, m.AdditionalParameters))
	if err != nil {
		return
	}

	db.SetConnMaxLifetime(m.ConnMaxLifetime)
	db.SetMaxIdleConns(m.MaxIdleConns)
	db.SetMaxOpenConns(m.MaxOpenConns)

	err = db.Ping()
	if err != nil {
		return

	}

	return
}

// MySQLConfig builder pattern code
type MySQLConfigBuilder struct {
	mySQLConfig *MySQLConfig
}

func NewMySQLConfigBuilder() *MySQLConfigBuilder {
	mysqlPort, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))

	mySQLConfig := &MySQLConfig{
		Host:                 os.Getenv("MYSQL_HOST"),
		Port:                 mysqlPort,
		Username:             os.Getenv("MYSQL_USERNAME"),
		Password:             os.Getenv("MYSQL_PASSWORD"),
		AdditionalParameters: "charset=utf8&parseTime=True&loc=Asia%2fJakarta&time_zone=%27%2B07%3A00%27",
		MaxOpenConns:         280,
		MaxIdleConns:         0,
		ConnMaxLifetime:      5 * time.Second,
		DriverName:           "mysql",
	}
	b := &MySQLConfigBuilder{mySQLConfig: mySQLConfig}
	return b
}

func (b *MySQLConfigBuilder) WithDriverName(driver string) *MySQLConfigBuilder {
	b.mySQLConfig.DriverName = driver
	return b
}

func (b *MySQLConfigBuilder) WithHost(host string) *MySQLConfigBuilder {
	b.mySQLConfig.Host = host
	return b
}

func (b *MySQLConfigBuilder) WithPort(port int) *MySQLConfigBuilder {
	b.mySQLConfig.Port = port
	return b
}

func (b *MySQLConfigBuilder) WithUsername(username string) *MySQLConfigBuilder {
	b.mySQLConfig.Username = username
	return b
}

func (b *MySQLConfigBuilder) WithPassword(password string) *MySQLConfigBuilder {
	b.mySQLConfig.Password = password
	return b
}

func (b *MySQLConfigBuilder) WithDBName(dBName string) *MySQLConfigBuilder {
	b.mySQLConfig.DBName = dBName
	return b
}

func (b *MySQLConfigBuilder) WithAdditionalParameters(additionalParameters string) *MySQLConfigBuilder {
	b.mySQLConfig.AdditionalParameters = additionalParameters
	return b
}

func (b *MySQLConfigBuilder) WithMaxOpenConns(maxOpenConns int) *MySQLConfigBuilder {
	b.mySQLConfig.MaxOpenConns = maxOpenConns
	return b
}

func (b *MySQLConfigBuilder) WithMaxIdleConns(maxIdleConns int) *MySQLConfigBuilder {
	b.mySQLConfig.MaxIdleConns = maxIdleConns
	return b
}

func (b *MySQLConfigBuilder) WithConnMaxLifetime(connMaxLifetime time.Duration) *MySQLConfigBuilder {
	b.mySQLConfig.ConnMaxLifetime = connMaxLifetime
	return b
}

func (b *MySQLConfigBuilder) Build() *MySQLConfig {
	return b.mySQLConfig
}

func NewMysqlClient() IMySQL {
	c := config.AppCfg
	dbMySQL := NewMySQLConfigBuilder().
		WithHost(c.DB.Host).
		WithPort(c.DB.Port).
		WithUsername(c.DB.User).
		WithPassword(c.DB.Pass).
		WithDBName(c.DB.Name).
		Build()

	dbSql, err := dbMySQL.InitMysqlDB()
	if err != nil {
		log.Ctx(context.Background()).Error().Msgf("Error connection to DB | %v", err)
		os.Exit(1)
	}

	return &MySQLDb{
		Client: dbSql.DB,
	}
}

func InitMySQL() {
	mysqlClient := NewMysqlClient()
	MySqlDB = mysqlClient
}

type IMySQL interface {
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row
	Ping(ctx context.Context) error
	GetDB() *sql.DB
}

func (m MySQLDb) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return m.Client.ExecContext(ctx, query, args...)
}

func (m MySQLDb) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return m.Client.QueryRowContext(ctx, query, args...)
}

func (m MySQLDb) Ping(ctx context.Context) error {
	return m.Client.PingContext(ctx)
}

func (m MySQLDb) GetDB() *sql.DB {
	return m.Client
}
