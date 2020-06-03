package main

import (
	"bytes"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
	"time"
)

func sendLogsToTelegram(chatID int64, logs *bytes.Buffer) error {
	token := os.Getenv("TG_BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	logFileName := fmt.Sprintf("%d.log", time.Now().Unix())
	logFile, err := os.Create(fmt.Sprintf(logFileName))
	if err != nil {
		return err
	}

	logsBytes := logs.Next(logs.Len())

	_, err = logFile.Write(logsBytes)
	if err != nil {
		return err
	}

	logFile.Close()

	msg := tgbotapi.NewDocumentUpload(chatID, logFileName)

	_, err = bot.Send(msg)
	if err != nil {
		return err
	}

	os.Remove(logFileName)

	return nil
}
