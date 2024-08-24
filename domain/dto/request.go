package dto

import "fmt"

type Register struct{
	TeleID     string `json:"teleID"`
	Username     string `json:"username"`
}

// Mapping
func RequestToRegisterDTO(teleID int64, username string) Register {
	return Register{
		TeleID: fmt.Sprint(teleID),
		Username: username,
	}
}

