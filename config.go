package InstagramSpyBot

import "github.com/spf13/viper"

type Config struct {
	TokenBot  string `mapstructure:"TOKEN_BOT"`
	ChatID    string `mapstructure:"CHAT_ID"`
	SessionID string `mapstructure:"SESSION_ID"`
	AppID     string `mapstructure:"APP_ID"`
	UserAgent string `mapstructure:"USER_AGENT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("configuration")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
