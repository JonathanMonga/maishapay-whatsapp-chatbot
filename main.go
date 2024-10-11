package main

import (
	"log"

	chatbot "github.com/green-api/whatsapp-chatbot-golang"
	"github.com/green-api/whatsapp-demo-chatbot-golang/scenes"
	"github.com/joho/godotenv"
)

func main() {
	idInstance := "{ID_INSTANCE}"
	authToken := "{AUTH_TOKEN}"
	envFile, err := godotenv.Read("instance.env")
	if err == nil {
		if val, exists := envFile["ID_INSTANCE"]; exists && len(val) > 0 {
			idInstance = val
		}
		if val, exists := envFile["AUTH_TOKEN"]; exists && len(val) > 0 {
			authToken = val
		}
	}

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

	_, err = bot.GreenAPI.Methods().Account().SetSettings(map[string]interface{}{
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
}
