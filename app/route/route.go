package route

import (
	"tsheet/domain/handler"

	tele "gopkg.in/telebot.v3"
)

func MainRoute(b *tele.Bot) {
	hand := handler.NewHandler()

	//User Route
	user := b.Group()
	user.Handle("/register", hand.Register)
}
