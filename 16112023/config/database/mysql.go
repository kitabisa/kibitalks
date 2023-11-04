package database

import (
	"context"
	"database/sql"
	"github.com/kitabisa/kibitalk/config"
	"github.com/kitabisa/perkakas/database/mysql"
	plog "github.com/kitabisa/perkakas/log"
	"os"
)

type MySQLDb struct {
	Client *sql.DB
}

var MySqlDB IMySQL

type IMySQL interface {
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row
	Ping(ctx context.Context) error
	GetDB() *sql.DB
}

func NewMysqlClient() IMySQL {
	c := config.AppCfg
	dbMySQL, err := mysql.NewMySQLConfigBuilder().
		WithHost(c.DB.Host).
		WithPort(c.DB.Port).
		WithUsername(c.DB.User).
		WithPassword(c.DB.Pass).
		WithDBName(c.DB.Name).
		Build()

	if err != nil {
		plog.Zlogger(context.Background()).Error().Msgf("Error connection to DB | %v", err)
		os.Exit(1)
	}

	dbSql, err := dbMySQL.InitMysqlDB()
	if err != nil {
		plog.Zlogger(context.Background()).Error().Msgf("Error connection to DB | %v", err)
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
