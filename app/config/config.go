package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	SECRET_JWT = ""
)

type AppConfig struct {
	SERVER_PORT int
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOSTNAME string
	DB_PORT     int
	DB_NAME     string
	JWT_KEY     string
	OPENAI_API_KEY string
	SMTP_SERVER string
	SMTP_PORT 	int
	SMTP_USERNAME string
	SMTP_PASSWORD string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	// inisialisasi variabel dg type struct AppConfig
	app := AppConfig{}

	godotenv.Load(".env")

	// proses mencari & membaca environment var dg key tertentu
	if val, found := os.LookupEnv("SERVERPORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.SERVER_PORT = cnv
	}
	// if val, found := os.LookupEnv("OPENAI_API_KEY"); found {
	// 	app.OPENAI_API_KEY = val
	// }
	if val, found := os.LookupEnv("JWT_KEY"); found {
		app.JWT_KEY = val
	}
	if val, found := os.LookupEnv("DBUSER"); found {
		app.DB_USERNAME = val
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		app.DB_PASSWORD = val
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DB_HOSTNAME = val
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.DB_PORT = cnv
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DB_NAME = val
	}
	// if val, found := os.LookupEnv("SMTPSERVER"); found {
	// 	app.SMTP_SERVER = val
	// }
	// if val, found := os.LookupEnv("SMTPPORT"); found {
	// 	cnv, _ := strconv.Atoi(val)
	// 	app.SMTP_PORT = cnv
	// }
	// if val, found := os.LookupEnv("SMTPUSERNAME"); found {
	// 	app.SMTP_USERNAME = val
	// }
	// if val, found := os.LookupEnv("SMTPPASSWORD"); found {
	// 	app.SMTP_PASSWORD = val
	// }

	// SECRET_JWT = app.JWT_KEY
	// //fmt.Print(app.OPENAI_API_KEY)
	return &app
}
