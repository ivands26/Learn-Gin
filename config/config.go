package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Driver   string
	Name     string
	Address  string
	Port     int
	Username string
	Password string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Cannot read configuration")
		return nil
	}

	SECRET = os.Getenv("SECRET")
	cnv, err := strconv.Atoi(os.Getenv("SERVERPORT"))
	if err != nil {
		log.Fatal("Failed to parse serverport")
		return nil
	}
	SERVERPORT = int16(cnv)
	defaultConfig.Name = os.Getenv("DBName")
	defaultConfig.Address = os.Getenv("DBHost")
	defaultConfig.Username = os.Getenv("DBUsername")
	defaultConfig.Password = os.Getenv("DBPassword")
	cnvPort, err := strconv.Atoi(os.Getenv("DBPort"))
	if err != nil {
		log.Fatal("Failed to parse DBPort")
		return nil
	}
	defaultConfig.Port = cnvPort

	return &defaultConfig

}
