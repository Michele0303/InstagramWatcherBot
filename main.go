package InstagramSpyBot

import (
	"log"
)

func main() {

	// load configuration from file.env
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// setup telegram bot
	setupTelegramBot(config.TokenBot, config.ChatID)

	// setup client for request
	setupHttpClient(config.SessionID, config.AppID, config.UserAgent)

	// list of users to be monitored
	usernames := map[string]string{"fedez": "", "chiaraferragni": ""}

	// run bot
	runBot(usernames)
}
