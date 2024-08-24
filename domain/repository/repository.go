package repository

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"tsheet/domain/dto"
	"tsheet/domain/types"

	"github.com/joho/godotenv"
)

type TsheetRepository struct {}

func NewRepository() types.TsheetRepositoryType {
	return &TsheetRepository{}
}

func (t *TsheetRepository) Register(data dto.Register) (string, error) {
	dataResponse := dto.ResponseRegister{}

	godotenv.Load(".env")

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(os.Getenv("URL")+"/register", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &dataResponse)
	if err != nil {
		return "", err
	}

	return dataResponse.Message, nil
}
