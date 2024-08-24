package utils

import 	tele "gopkg.in/telebot.v3"

var (
	menu     = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnHelp     = menu.Text("Help")
	BtnSettings = menu.Text("⚙ Settings")
)