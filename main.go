package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {

	// draw banner
	banner()

	// load configuration from file.env
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// setup telegram bot
	setupTelegramBot(config.TokenBot, config.ChatID)
	fmt.Println("[*] TelegramBot successfully configured")

	// setup client for request
	setupHttpClient(config.SessionID, config.AppID, config.UserAgent)
	fmt.Println("[*] HttpClient successfully configured")

	// list of users to be monitored
	usernames := strings.Split(config.WatchList, ";")
	watchlist := make(map[string]string)
	for _, username := range usernames {
		watchlist[username] = ""
	}
	fmt.Println("[*] WatchList ->", usernames)

	// run bot
	runBot(watchlist)
}
