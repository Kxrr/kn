package kn

import (
	"gopkg.in/telegram-bot-api.v4"
	log "github.com/kxrr/klog"
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
	log.Printf("Connected with bot %s", bot.Self.UserName)
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
		_, err := t.SendMdMessage(text)
		if err != nil {
			log.Errorf("Send md message error: %s", err.Error())
		}
	}()

}

func (t *TelegramMangonel) SendHtmlMessage(text string) (tgbotapi.Message, error) {
	return t.bot.Send(tgbotapi.NewMessage(t.toUserId, text))
}



func newMdMessage(chatID int64, text string) tgbotapi.MessageConfig {
	m := tgbotapi.NewMessage(chatID, text)
	m.ParseMode = "Markdown"
	return m
}
