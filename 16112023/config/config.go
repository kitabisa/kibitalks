package config

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/iamolegga/enviper"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

var AppCfg AppConfig

func readViperConfig() *viper.Viper {
	// handle config path for unit test
	dirPath, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("Error get working dir: %s", err))
	}
	dirPaths := strings.Split(dirPath, "/internal")

	godotenv.Load(fmt.Sprintf("%s/params/.env", dirPaths[0]))
	godotenv.Load("./params/.env")
	godotenv.Load()

	v := viper.New()
	v.AllowEmptyEnv(true)
	v.AutomaticEnv()
	return v
}

func NewAppConfig() {
	var conf AppConfig
	e := enviper.New(readViperConfig())
	err := e.Unmarshal(&conf)
	if err != nil {
		log.Ctx(context.Background()).Fatal().Msgf("Cannot start app. Fatal error occured during SetupMySQLMaster | %v | %s", err, "exiting now..")
	}

	json.Marshal(conf)

	log.Logger = log.With().Caller().Logger()
	zerolog.TimeFieldFormat = time.RFC3339

	if conf.Log.Pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, FormatTimestamp: func(i interface{}) string { return time.Now().Format(time.RFC3339) }})
	}

	switch conf.Log.Level {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	log.Ctx(context.Background()).Debug().Msgf("Config: %+v\n", conf)

	AppCfg = conf
	return
}

type AppConfig struct {
	App         App         `mapstructure:"app"`
	DB          DB          `mapstructure:"db"`
	Log         Log         `mapstructure:"log"`
	Cache       Cache       `mapstructure:"cache"`
	RabbitMQ    RabbitMQ    `mapstructure:"rabbitmq"`
	ApiPayment  ApiPayment  `mapstructure:"api_payment"`
	ApiCampaign ApiCampaign `mapstructure:"api_campaign"`
}

type ApiPayment struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ApiCampaign struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Cache struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Log struct {
	Level  string `mapstructure:"level"`
	Pretty bool   `mapstructure:"pretty"`
}

type App struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type DB struct {
	Host     string   `mapstructure:"host"`
	Port     int      `mapstructure:"port"`
	Name     string   `mapstructure:"name"`
	User     string   `mapstructure:"user"`
	Pass     string   `mapstructure:"pass"`
	RabbitMQ RabbitMQ `mapstructure:"rabbitmq"`
}

type RabbitMQ struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
}
