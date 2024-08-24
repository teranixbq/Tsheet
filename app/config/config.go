package config

import (
	"os"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	TOKEN	   string
	URL 	  string
}

func InitConfig() *AppConfig {
	var res = new(AppConfig)
	res = loadConfig()

	if res == nil {
		logrus.Fatal("Config : Cannot start program, failed to load configuration")
		return nil
	}

	return res
}

func loadConfig() *AppConfig {
	var res = new(AppConfig)

	godotenv.Load(".env")

	// if val, found := os.LookupEnv("SERVERPORT"); found {
	// 	port, err := strconv.Atoi(val)
	// 	if err != nil {
	// 		logrus.Error("Config : invalid port value,", err.Error())
	// 		return nil
	// 	}
	// 	res.SERVERPORT = port
	// }

	if val, found := os.LookupEnv("TOKEN"); found {
		res.TOKEN = val
	}

	if val, found := os.LookupEnv("URL"); found {
		res.URL = val
	}

	return res
}
