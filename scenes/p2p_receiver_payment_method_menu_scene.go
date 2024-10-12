package scenes

import (
	"maishapay-whatsapp-chatbot/util"

	chatbot "github.com/green-api/whatsapp-chatbot-golang"
)

type ReceiverPaymentMethodMenuScene struct {
	lang string
}

func (s ReceiverPaymentMethodMenuScene) Start(bot *chatbot.Bot) {
	bot.IncomingMessageHandler(func(message *chatbot.Notification) {
		if !util.IsSessionExpired(message) {
			text, _ := message.Text()
			switch text {
			case "1":
				s.saveMobileOperatorAndAskRecipientPhoneNumber(message, "AIRTEL")
			case "2":
				s.saveMobileOperatorAndAskRecipientPhoneNumber(message, "VODACOM")
			case "3":
				s.saveMobileOperatorAndAskRecipientPhoneNumber(message, "ORANGE")
			case "4":
				s.saveMobileOperatorAndAskRecipientPhoneNumber(message, "CREDIT_CARD")
			case "000":
				sendStopMenu(message, s.lang)
			default:
				message.SendText(util.GetString([]string{"menu", s.lang, "p2p_receiver_payment_method_menu"}))
			}
		} else {
			message.ActivateNextScene(SelectLanguageScene{})
			message.SendText(util.GetString([]string{"select_language"}))
		}
	})
}

func (s ReceiverPaymentMethodMenuScene) saveMobileOperatorAndAskRecipientPhoneNumber(message *chatbot.Notification, paymentMethod string) {
	message.UpdateStateData(map[string]interface{}{"receiver_payment_method": paymentMethod})

	message.SendText(util.GetString([]string{"p2p", s.lang, "ask_for_recipient_phone_number"}))
	message.ActivateNextScene(AskForRecipientPhoneNumberScene{lang: s.lang})
}
