package InstagramWatcherBot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

var bot *tgbotapi.BotAPI
var chatID int64

func sendPhoto(mediaUrl string, username string) {
	msg := tgbotapi.NewPhotoShare(chatID, mediaUrl)
	msg.Caption = username
	bot.Send(msg)
}

func sendVideo(mediaUrl string, username string) {
	msg := tgbotapi.NewVideoShare(chatID, mediaUrl)
	msg.Caption = username
	bot.Send(msg)
}

func setupTelegramBot(tokenBot string, id string) {
	// parse chatID string to int64
	chatIDTemp, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	chatID = chatIDTemp

	// run telegram bot
	botTemp, err := tgbotapi.NewBotAPI(tokenBot)
	if err != nil {
		log.Fatal(err)
	}
	bot = botTemp

}
