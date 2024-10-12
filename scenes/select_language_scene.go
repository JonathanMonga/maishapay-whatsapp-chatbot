package scenes

import (
	"maishapay-whatsapp-chatbot/util"

	chatbot "github.com/green-api/whatsapp-chatbot-golang"
)

type SelectLanguageScene struct{}

func (s SelectLanguageScene) Start(bot *chatbot.Bot) {
	bot.IncomingMessageHandler(func(message *chatbot.Notification) {
		if !util.IsSessionExpired(message) {
			text, _ := message.Text()
			switch text {
			case "1":
				s.saveLanguageAndSendMainMenu(message, "en")
			case "2":
				s.saveLanguageAndSendMainMenu(message, "fr")
			case "3":
				s.saveLanguageAndSendMainMenu(message, "sw")
			default:
				message.SendText(util.GetString([]string{"specify_language"}))
			}
		} else {
			message.ActivateNextScene(SelectLanguageScene{})
			message.SendText(util.GetString([]string{"select_language"}))
		}
	})
}

func (s SelectLanguageScene) saveLanguageAndSendMainMenu(message *chatbot.Notification, lang string) {
	message.UpdateStateData(map[string]interface{}{"lang": lang})
	sendMainMenu(message, lang)
}
