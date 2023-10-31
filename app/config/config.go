package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
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
	JWT_KEY      string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	// inisialisasi variabel dg type struct AppConfig
	app := AppConfig{}
	isRead := true

	// proses mencari & membaca environment var dg key tertentu
	if val, found := os.LookupEnv("SERVERPORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.SERVER_PORT = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("JWT_KEY"); found {
		app.JWT_KEY = val
		isRead = false
	}
	
	if val, found := os.LookupEnv("DBUSER"); found {
		app.DB_USERNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		app.DB_PASSWORD = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DB_HOSTNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.DB_PORT = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DB_NAME = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}
		
		app.SERVER_PORT, _ = strconv.Atoi(viper.Get("SERVERPORT").(string))
		app.JWT_KEY = viper.Get("JWT_KEY").(string)
		app.DB_USERNAME = viper.Get("DBUSER").(string)
		app.DB_PASSWORD = viper.Get("DBPASS").(string)
		app.DB_HOSTNAME = viper.Get("DBHOST").(string)
		app.DB_PORT, _ = strconv.Atoi(viper.Get("DBPORT").(string))
		app.DB_NAME = viper.Get("DBNAME").(string)
	}

	SECRET_JWT = app.JWT_KEY
	fmt.Println("check", app.JWT_KEY)
	return &app
}
