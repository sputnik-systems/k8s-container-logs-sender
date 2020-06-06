package main

import (
	"bytes"
	// "errors"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
	"time"
)

func sendLogsToTelegram(chatID int64, logs *bytes.Buffer, prefix string) error {
	token := os.Getenv("TG_BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return fmt.Errorf("[sendLogsToTelegram] failed create tg bot api connection: %s", err)
	}

	logFileName := fmt.Sprintf("%s_%d.log", prefix, time.Now().Unix())
	logFile, err := os.Create(logFileName)
	if err != nil {
		return fmt.Errorf("[sendLogsToTelegram] failed create log file %s: %s", logFileName, err)
	}

	logsBytes := logs.Next(logs.Len())

	_, err = logFile.Write(logsBytes)
	if err != nil {
		return fmt.Errorf("[sendLogsToTelegram] failed write bytes to file: %s", err)
	}

	logFile.Close()

	msg := tgbotapi.NewDocumentUpload(chatID, logFileName)

	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("[sendLogsToTelegram] failed send message to tg: %s", err)
	}

	err = os.Remove(logFileName)
	if err != nil {
		return fmt.Errorf("[sendLogsToTelegram] remove file %s: %s", logFileName, err)
	}

	return nil
}
