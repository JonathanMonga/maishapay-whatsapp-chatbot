package scenes

import (
	"maishapay-whatsapp-chatbot/util"

	"github.com/go-playground/validator/v10"
	chatbot "github.com/green-api/whatsapp-chatbot-golang"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

type AskForAmountToSendScene struct{ 
	lang string 
}

func (s AskForAmountToSendScene) Start(bot *chatbot.Bot) {
	bot.IncomingMessageHandler(func(message *chatbot.Notification) {
		if !util.IsSessionExpired(message) {
			text, _ := message.Text()

			validate = validator.New(validator.WithRequiredStructEnabled())
			errs := validate.Var(text, "required,number")

			if errs == nil {
				if text != "000" {
					s.asveAmountAndSendReceiverPaymentMethodMenu(message, text)
				} else {
					sendStopMenu(message, s.lang)
				}
			} else {
				message.SendText(util.GetString([]string{"p2p", s.lang, "ask_for_amount_to_send"}))
			}
		} else {
			message.ActivateNextScene(StartScene{})
		}
	})
}

func (s AskForAmountToSendScene) asveAmountAndSendReceiverPaymentMethodMenu(message *chatbot.Notification, amount string) {
	message.UpdateStateData(map[string]interface{}{"amount": amount})

	message.SendText(util.GetString([]string{"menu", s.lang, "p2p_receiver_payment_method_menu"}))
	message.ActivateNextScene(ReceiverPaymentMethodMenuScene{lang: s.lang})
}
