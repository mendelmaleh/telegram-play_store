package main

import (
	"fmt"
	"net/http"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	// PlayURL is the prefix that goes before the package name
	PlayURL = "https://play.google.com/store/apps/details?id="
)

func main() {
	config := getConfig("bot.json")
	svc := service(config.Key)

	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		panic(err)
	}

	bot.Debug = false
	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert(
		"https://" + config.Domain + ":8443/" + bot.Token, "cert.pem"))
	if err != nil {
		panic(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)

	msg := tgbotapi.NewMessage(config.Updates, "Starting...")
	msg.DisableNotification = true
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}

	for update := range updates {
		if update.InlineQuery != nil {
			if len(update.InlineQuery.Query) > 2 {
				ans := answer(svc, config.Cx, update.InlineQuery.ID, update.InlineQuery.Query)
				bot.AnswerInlineQuery(ans)
			}
		}

		if update.ChosenInlineResult != nil {
			bot.Send(tgbotapi.NewMessage(config.Updates,
				"@"+update.ChosenInlineResult.From.String()+" searched "+update.ChosenInlineResult.Query))
		}
	}
}
