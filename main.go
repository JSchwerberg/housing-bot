package main

import (
	"fmt"
	"os"
	//"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	api_token := os.Getenv("TG_HOUSING_BOT_API_TOKEN")

	if len(api_token) < 1 {
		fmt.Println("TG_HOUSING_BOT_API_TOKEN must be set in environment.")
		os.Exit(1)
	}
}
