package service

import (
	"tsheet/domain/dto"
	"tsheet/domain/repository"
	"tsheet/domain/types"
)

type TsheetService struct {
	ts repository.TsheetRepository
}


func NewService() types.TsheetServiceType {
	return &TsheetService{}
}

func (t *TsheetService) Register(data dto.Register) (string, error) {
	response, err := t.ts.Register(data)
	if err != nil {
		return "", err
	}

	return response, nil
}