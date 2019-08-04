package stromberg_bot

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"os"
)

const configFilePath = "config.json"

type Config struct {
	BotToken	string `json:"bot_token"`
	BaseUrl		string `json:"base_url"`
}

var Conf *Config

func readConfigFile() {
	file, err := os.Open(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(Conf)
	if err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}


func main() {
	bot, err := tgbotapi.NewBotAPI(Conf.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(Conf.BaseUrl+bot.Token))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback last failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe("0.0.0.0:8080", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
		handleUpdate(bot, update)
	}
}

func handleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	if update.Message.IsCommand() {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		switch update.Message.Command() {
		case "help":
			msg.Text = "Ich geb dir gleich HILFE!"
		case "start":
			msg.Text = "Na, auch am Faulenzen?"
		case "status":
			msg.Text = "Lass das mal den Papa machen!"
		default:
			msg.Text = "Du, selbst der Papa kennt diesen Befehl nicht!"
		}
		bot.Send(msg)
	}
}