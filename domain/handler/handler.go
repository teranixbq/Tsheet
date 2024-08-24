package handler

import (
	"tsheet/domain/dto"
	"tsheet/domain/service"

	tele "gopkg.in/telebot.v3"
)

type TsheetHandler struct {
	ts service.TsheetService
}

func NewHandler() *TsheetHandler {
	return &TsheetHandler{}
}

func (t *TsheetHandler) Register(c tele.Context) error {

	Request := c.Sender()
	telegramID := Request.ID
	username := Request.Username

	
	input := dto.RequestToRegisterDTO(telegramID, username)

	result, err := t.ts.Register(input)
	if err != nil {
		return err
	}

	return c.Send(result)

}