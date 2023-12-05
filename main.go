package main

import (
	"fmt"
	"log"
	"telegram-bot/functions"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error to add file .env \n", err.Error())
		panic(err)
	}

	chatId := -123456
	topicId := 1234
	text := "Bot message"
	response, _ := functions.SendTextToTelegramChat(chatId, topicId, text)
	fmt.Println(response)
}
