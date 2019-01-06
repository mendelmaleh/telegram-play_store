package main

import (
	"fmt"

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

	msg := tgbotapi.NewMessage(config.Updates, "Starting...")
	msg.DisableNotification = true
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

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

	/*
		r := search(svc, config.Cx, "niagara launcher")
		fmt.Println(r)
		fmt.Println(len(r))
	*/
}
