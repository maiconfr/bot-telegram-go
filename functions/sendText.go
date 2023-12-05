package functions

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// Pass token and sensible APIs through environment variables
const telegramApiBaseUrl string = "https://api.telegram.org/bot"
const telegramApiSendMessage string = "/sendMessage"
const telegramTokenEnv string = "TELEGRAM_BOT_TOKEN"

// sendTextToTelegramChat sends a text message to the Telegram chat identified by its chat Id
func SendTextToTelegramChat(chatId int, topicId int, text string) (string, error) {
	var telegramApi string = telegramApiBaseUrl + os.Getenv(telegramTokenEnv) + telegramApiSendMessage
	log.Printf("Sending %s to chat_id: %d", text, chatId)
	log.Printf("telegramApi: %s", telegramApi)
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id":             {strconv.Itoa(chatId)},
			"reply_to_message_id": {strconv.Itoa(topicId)},
			"text":                {text},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = io.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)

	return bodyString, nil
}
