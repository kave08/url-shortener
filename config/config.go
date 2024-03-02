package config

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var Cfg config

type config struct {
	Mysql    Mysql    `yml:"Mysql"`
	Redis    Redis    `yml:"Redis"`
	ApiToken ApiToken `yml:"ApiToken"`
}

type Mysql struct {
	Username           string `yml:"Username"`
	Password           string `yml:"Password"`
	Host               string `yml:"Host"`
	Port               string `yml:"Port"`
	DBName             string `yml:"DB_Name"`
	MaxOpenConnections int    `yml:"Max_Open_Connections"`
	MaxIdleConnections int    `yml:"Max_Idle_Connections"`
}

type Redis struct {
	Address  string        `yml:"Address"`
	Password string        `yml:"Password"`
	TTL      time.Duration `yml:"TTL"`
}

type ApiToken struct {
	Token string `yml:"Token"`
}

type SetupResult struct {
	MysqlConnection *sql.DB
	RedisConnection *redis.Client
}

func LoadConfig(configPath string) *SetupResult {

	viper.SetEnvPrefix("url-shortener")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigFile(configPath)
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.MergeInConfig()
	if err != nil {
		fmt.Println("Error in reading config")
		panic(err)
	}

	err = viper.Unmarshal(&Cfg, func(config *mapstructure.DecoderConfig) {
		config.TagName = "yml"
	})
	if err != nil {
		fmt.Println("Error in unmarshaling config")
		panic(err)
	}

	fmt.Printf("%v", Cfg)

	mdb, err := initializeMySQL()
	if err != nil {
		panic(err)
	}

	rdb, err := initializeRedisConnection()
	if err != nil {
		panic(err)
	}

	return &SetupResult{
		MysqlConnection: mdb,
		RedisConnection: rdb,
	}
}
