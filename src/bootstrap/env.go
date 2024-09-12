package bootstrap

import (
	"log"

	"github.com/spf13/viper"

	"github.com/wizard-corp/api-gateway/src/domain"
)

type Env struct {
	AppEnv                 int    `mapstructure:"APP_ENV"`
	SystemUId              string `mapstructure:"SYSTEM_UID"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	RateLimit              int    `mapstructure:"RATE_LIMIT"`
	Burts                  int    `mapstructure:"BURTS"`
	TTL                    int64  `mapstructure:"TTL"`
	MongoHost              string `mapstructure:"MONGO_HOST"`
	MongoPort              int    `mapstructure:"MONGO_PORT"`
	MongoUser              string `mapstructure:"MONGO_USER"`
	MongoPassword          string `mapstructure:"MONGO_PASSWORD"`
	MongoDatabase          string `mapstructure:"MONGO_DATABASE"`
	RabbitmqHost           string `mapstructure:"RABBITMQ_HOST"`
	RabbitmqPort           int    `mapstructure:"RABBITMQ_PORT"`
	RabbitmqUser           string `mapstructure:"RABBITMQ_USER"`
	RabbitmqPassword       string `mapstructure:"RABBITMQ_PASSWORD"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
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

	if env.AppEnv == 0 {
		log.Println(domain.DEV_ENVIRONMENT)
	}

	return &env
}
