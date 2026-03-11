package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/joho/godotenv"
)

func GetAdmins() []int64 {
	strings := strings.Split(os.Getenv("TG_ADMINS"), ",")
	ints := make([]int64, 0, len(strings))
	for _, str := range strings {
		num, err := strconv.ParseInt(str, 10, 0)
		if err != nil {
			log.Panic(err)
		}
		ints = append(ints, num)
	}
	return ints
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, relying on system environment")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	admins := GetAdmins()

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !slices.Contains(admins, update.Message.Chat.ID) {
			continue
		}

		response := ""

		msg_args := []string{}
		if update.Message.CommandArguments() != "" {
			msg_args = strings.Split(update.Message.CommandArguments(), " ")
		}

		if !update.Message.IsCommand() {
			response = "Provide a command."
		} else if update.Message.Command() == "status" {
			if len(msg_args) > 0 {
				if msg_args[0] == "servers" {
					response = "servers..?"
				}
			} else {
				response = "Bot is up and running."
			}
		} else {
			response = "Unknown command."
		}

		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, response))
	}
}
