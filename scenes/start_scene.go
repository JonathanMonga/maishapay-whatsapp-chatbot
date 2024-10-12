package scenes

import (
	"maishapay-whatsapp-chatbot/util"

	chatbot "github.com/green-api/whatsapp-chatbot-golang"
)

type StartScene struct{}

func (s StartScene) Start(bot *chatbot.Bot) {
	bot.IncomingMessageHandler(func(message *chatbot.Notification) {
		util.IsSessionExpired(message)

		message.SendText(util.GetString([]string{"select_language"}))
		message.ActivateNextScene(SelectLanguageScene{})
	})
}

func sendStopMenu(message *chatbot.Notification, lang string) {
	message.SendText(util.GetString([]string{"stop_message", lang})+"*"+message.Body["senderData"].(map[string]interface{})["senderName"].(string)+"*!", "false")
	message.ActivateNextScene(StartScene{})
}

func sendMainMenu(message *chatbot.Notification, lang string) {
	message.SendUrlFile("https://raw.githubusercontent.com/JonathanMonga/maishapay-whatsapp-chatbot/refs/heads/main/assets/Mascot.jpg",
		"Mascot.jpg",
		util.GetString([]string{"salutation_message", lang})+"*"+message.Body["senderData"].(map[string]interface{})["senderName"].(string)+"*!"+"\n"+util.GetString([]string{"welcome_message", lang}))
	message.SendText(util.GetString([]string{"menu", lang, "main_menu"}))

	message.ActivateNextScene(SelectMainMenuScene{lang: lang})
}
