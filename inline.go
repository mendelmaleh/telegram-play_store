package main

import (
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	customsearch "google.golang.org/api/customsearch/v1"
)

func answer(svc *customsearch.Service, cx, id, query string) tgbotapi.InlineConfig {
	results := search(svc, cx, query)
	var answers []interface{}
	for i, v := range results {
		title, _ := v["name"].(string)
		desc, _ := v["description"].(string)
		url, _ := v["url"].(string)

		article := tgbotapi.InlineQueryResultArticle{
			Type:        "article",
			ID:          strconv.Itoa(i),
			Title:       title,
			Description: desc,
			InputMessageContent: tgbotapi.InputTextMessageContent{
				Text: url,
			},
		}
		answers = append(answers, article)
	}
	config := tgbotapi.InlineConfig{
		InlineQueryID: id,
		Results:       answers,
	}
	return config
}
