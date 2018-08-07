package kn

import (
	"gopkg.in/telegram-bot-api.v4"
)

type TelegramMangonel struct {
	toUserId int64
	bot      *tgbotapi.BotAPI
}

func NewTelegramMangonel(toUserId int64, botToken string, debug bool) (*TelegramMangonel, error) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}
	if debug {
		bot.Debug = true
	}
	return &TelegramMangonel{toUserId, bot}, nil

}

/**
Supported Md syntax: https://core.telegram.org/bots/api#markdown-style
 */
func (t *TelegramMangonel) SendMdMessage(text string) (tgbotapi.Message, error) {
	return t.bot.Send(newMdMessage(t.toUserId, text))
}

func (t *TelegramMangonel) SendMdMessageNoWait(text string) {
	go func() {
		t.SendMdMessage(text)
	}()

}
func newMdMessage(chatID int64, text string) tgbotapi.MessageConfig {
	m := tgbotapi.NewMessage(chatID, text)
	m.ParseMode = "Markdown"
	return m
}
