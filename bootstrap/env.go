package bootstrap

import (
	"log"

	"github.com/spf13/viper"

	"github.com/wizard-corp/api-gateway/domain"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	MongoHost              string `mapstructure:"MONGO_HOST"`
	MongoPort              int    `mapstructure:"MONGO_PORT"`
	MongoUser              string `mapstructure:"MONGO_USER"`
	MongoPassword          string `mapstructure:"MONGO_PASSWORD"`
	MongoDatabase          string `mapstructure:"MONGO_DATABASE"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(domain.INVALID_PATH, err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal(domain.LOAD_FAIL, err)
	}

	if env.AppEnv == "development" {
		log.Println(domain.DEV_ENVIRONMENT)
	}

	return &env
}
