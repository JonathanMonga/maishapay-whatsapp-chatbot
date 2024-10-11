package scenes

import (
	"maishapay-whatsapp-chatbot/util"

	chatbot "github.com/green-api/whatsapp-chatbot-golang"
)

type StartScene struct{}

func (s StartScene) Start(bot *chatbot.Bot) {
	bot.IncomingMessageHandler(func(message *chatbot.Notification) {
		util.IsSessionExpired(message)

		message.SendText(util.GetString([]string{"welcome_message"}))
		message.SendText(util.GetString([]string{"select_language"}))
		message.ActivateNextScene(MainMenuScene{})
	})
}
