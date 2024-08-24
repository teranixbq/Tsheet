package types

import "tsheet/domain/dto"

type TsheetRepositoryType interface {
	Register(data dto.Register)(string, error)
}

type TsheetServiceType interface {
	Register(data dto.Register)(string, error)
}