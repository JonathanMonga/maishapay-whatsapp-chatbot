package scenes

import (
	"maishapay-whatsapp-chatbot/util"

	"github.com/go-playground/validator/v10"
	chatbot "github.com/green-api/whatsapp-chatbot-golang"
)

type AskForRecipientPhoneNumberScene struct{ 
	lang string 
}

func (s AskForRecipientPhoneNumberScene) Start(bot *chatbot.Bot) {
	bot.IncomingMessageHandler(func(message *chatbot.Notification) {
		if !util.IsSessionExpired(message) {
			text, _ := message.Text()

			validate = validator.New(validator.WithRequiredStructEnabled())
			errs := validate.Var(text, "required,alphanum")

			if errs == nil {
				if text != "000" {
					s.sendIssuerPaymentMethodMenu(message, s.lang)
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

func (s AskForRecipientPhoneNumberScene) sendIssuerPaymentMethodMenu(message *chatbot.Notification, phoneNumber string) {
	message.UpdateStateData(map[string]interface{}{"phone_number": phoneNumber})

	message.SendText(util.GetString([]string{"menu", s.lang, "p2p_issuer_payment_method_menu"}))
	message.ActivateNextScene(IssuerPaymentMethodMenuScene{lang: s.lang})
}
