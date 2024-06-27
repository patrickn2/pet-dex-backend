package config

import "github.com/spf13/viper"

var env *envconfig

type envconfig struct {
	DBUrl                      string `mapstructure:"DATABASE_URL"`
	DBUrl_Migration            string `mapstructure:"MIGRATION_DATABASE_URL"`
	PORT                       string `mapstructure:"PORT"`
	ENV                        string `mapstructure:"ENVIRONMENT"`
	MIGRATIONS_PATH            string `mapstructure:"MIGRATIONS_PATH"`
	JWT_SECRET                 string `mapstructure:"JWT_SECRET"`
	GOOGLE_OAUTH_CLIENT_ID     string `mapstructure:"GOOGLE_OAUTH_CLIENT_ID"`
	GOOGLE_OAUTH_CLIENT_SECRET string `mapstructure:"GOOGLE_OAUTH_CLIENT_SECRET"`
	GOOGLE_REDIRECT_URL        string `mapstructure:"GOOGLE_REDIRECT_URL"`
}

func GetEnvConfig() *envconfig {
	return env
}

func LoadEnv(path string) (*envconfig, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		panic(err)
	}

	return env, nil
}
