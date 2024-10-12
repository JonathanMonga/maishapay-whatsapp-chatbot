package scenes

import (
	"maishapay-whatsapp-chatbot/util"

	chatbot "github.com/green-api/whatsapp-chatbot-golang"
)

type SelectMainMenuScene struct {
	lang string
}

func (s SelectMainMenuScene) Start(bot *chatbot.Bot) {

	bot.IncomingMessageHandler(func(message *chatbot.Notification) {
		if !util.IsSessionExpired(message) {
			text, _ := message.Text()

			switch text {
			case "1":
				s.sendPaymentMenu(message, s.lang)
			// case "2":
			// 	s.sendResubscriptionMenu(message, lang)
			// case "3":
			// 	s.sendCanalboxMenu(message, lang)
			// case "4":
			// 	s.sendAirtimesMenu(message, lang)
			// case "5":
			// 	s.sendHelpMenu(message, lang)
			case "stop", "Stop", "0":
				sendStopMenu(message, s.lang)
			case "menu":
				sendMainMenu(message, s.lang)
			default:
				message.SendText(util.GetString([]string{"not_recognized_message", s.lang}), "false")
			}
		} else {
			message.ActivateNextScene(StartScene{})
		}
	})
}

func (s SelectMainMenuScene) sendPaymentMenu(message *chatbot.Notification, lang string) {
	message.SendText(util.GetString([]string{"p2p", lang, "ask_for_amount_to_send"}))
	message.ActivateNextScene(AskForAmountToSendScene{lang: lang})
}

// func (s SelectMainMenuScene) sendResubscriptionMenu(message *chatbot.Notification, lang string) {
// 	message.SendText(util.GetString([]string{"main_menu", message.GetStateData()["lang"].(string)}))
// 	message.ActivateNextScene(SelectResubscriptionMenuScene{})
// }

// func (s SelectMainMenuScene) sendCanalboxMenu(message *chatbot.Notification, lang string) {
// 	message.SendText(util.GetString([]string{"main_menu", lang}))
// 	message.ActivateNextScene(SelectCanalboxMenuScene{})
// }

// func (s SelectMainMenuScene) sendAirtimesMenu(message *chatbot.Notification, lang string) {
// 	message.SendText(util.GetString([]string{"main_menu", lang}))
// 	message.ActivateNextScene(SelectAirtimesMenuScene{})
// }

// func (s SelectMainMenuScene) sendHelpMenu(message *chatbot.Notification, lang string) {
// 	message.SendText(util.GetString([]string{"main_menu", lang}))
// 	message.ActivateNextScene(SelectHelpMenuScene{})
// }
