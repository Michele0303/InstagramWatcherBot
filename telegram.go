package InstagramSpyBot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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
	chatID, _ = strconv.ParseInt(id, 10, 64)
	// run telegram bot
	bot, _ = tgbotapi.NewBotAPI(tokenBot)
}
