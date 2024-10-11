package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"maishapay-whatsapp-chatbot/scenes"

	chatbot "github.com/green-api/whatsapp-chatbot-golang"
)

func main() {
	idInstance := os.Getenv("ID_INSTANCE")
	authToken := os.Getenv("AUTH_TOKEN")
	port := os.Getenv("PORT")

	if idInstance == "{ID_INSTANCE}" || authToken == "{AUTH_TOKEN}" {
		log.Fatal("No idInstance or authToken set")
	}

	bot := chatbot.NewBot(idInstance, authToken)

	go func() {
		for err := range bot.ErrorChannel {
			if err != nil {
				log.Println(err)
			}
		}
	}()

	_, err := bot.GreenAPI.Methods().Account().SetSettings(map[string]interface{}{
		"incomingWebhook":            "yes",
		"outgoingMessageWebhook":     "yes",
		"outgoingAPIMessageWebhook":  "yes",
		"pollMessageWebhook":         "yes",
		"markIncomingMessagesReaded": "yes",
	})

	if err != nil {
		bot.ErrorChannel <- err
	}

	bot.SetStartScene(scenes.StartScene{})

	bot.StartReceivingNotifications()

	// listen and serve
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
