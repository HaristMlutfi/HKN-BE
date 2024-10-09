package config

import (
	"encoding/json"
	"hkn-be/constants"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server    ServerConfig
	DB        DBConfig
	Redis     RedisServer
	JwtConfig JwtConfig
	Mailer    Mailer
}

type ServerConfig struct {
	BaseUrl         string
	Port            string
	WriteTimeout    int
	ReadTimeout     int
	GracefulTimeout int
}

type DBConfig struct {
	Host            string
	Name            string
	Username        string
	Password        string
	MaxOpenConn     int
	MaxIdleConn     int
	MaxConnLifetime int
}

type RedisServer struct {
	Host     string
	Password string
	Timeout  int
	MaxIdle  int
}

type JwtConfig struct {
	Issuer            string
	Secret            string
	TokenLifeTimeHour int
}

type Mailer struct {
	Host       string
	Port       int
	Username   string
	Password   string
	UseTls     bool
	Sender     string
	MaxAttempt int
}

func InitConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}

	jwtTokenLifeTimeHour, err := strconv.Atoi(os.Getenv(constants.JwtTokenLifeTimeHour))
	if err != nil {
		log.Fatalf("Invalid JWT Token Lifetime Hour Format. Err : %s", err.Error())
	}

	mailerPort, err := strconv.Atoi(os.Getenv(constants.MailPort))
	if err != nil {
		log.Fatalf("Invalid Mailer Port Format, Err : %s", err.Error())
	}

	mailerMaxAttemptStr := os.Getenv(constants.MailMaxAttempt)

	if mailerMaxAttemptStr == "" {
		mailerMaxAttemptStr = "3" // Set default value to 3
	}

	mailerMaxAttempt, err := strconv.Atoi(mailerMaxAttemptStr)
	if err != nil {
		log.Fatalf("Invalid Mailer Max Attempt Format, Err : %s", err.Error())
	}

	mailerUseTls, err := strconv.ParseBool(os.Getenv(constants.MailUseTls))
	if err != nil {
		log.Fatalf("Invalid Mailer Use TLS Format, Err : %s", err.Error())
	}

	tmpConfig := Config{
		Server: ServerConfig{
			BaseUrl:         os.Getenv(constants.ServerBaseUrl),
			Port:            os.Getenv(constants.ServerPort),
			WriteTimeout:    10,
			ReadTimeout:     10,
			GracefulTimeout: 10,
		},
		DB: DBConfig{
			Host:            os.Getenv(constants.DbHost),
			Name:            os.Getenv(constants.DbName),
			Username:        os.Getenv(constants.DbUserName),
			Password:        os.Getenv(constants.DbPassword),
			MaxOpenConn:     10,
			MaxIdleConn:     10,
			MaxConnLifetime: 500,
		},
		Redis: RedisServer{
			Host:     os.Getenv(constants.RedisHost),
			Password: os.Getenv(constants.RedisPassword),
			Timeout:  10,
			MaxIdle:  10,
		},
		JwtConfig: JwtConfig{
			Issuer:            os.Getenv(constants.JwtIssuer),
			Secret:            os.Getenv(constants.JwtSecret),
			TokenLifeTimeHour: jwtTokenLifeTimeHour,
		},
		Mailer: Mailer{
			Host:       os.Getenv(constants.MailHost),
			Port:       mailerPort,
			Username:   os.Getenv(constants.MailUsername),
			Password:   os.Getenv(constants.MailPassword),
			UseTls:     mailerUseTls,
			Sender:     os.Getenv(constants.MailSender),
			MaxAttempt: mailerMaxAttempt,
		},
	}

	configString, _ := json.Marshal(tmpConfig)
	log.Println("============================================")
	log.Println("--> Config String : ")
	log.Println(string(configString))
	log.Println("============================================")
	log.Println("")

	return tmpConfig
}
