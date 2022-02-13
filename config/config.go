package config

import (
	"os"
	"strconv"
)

type Config struct {
	KafkaTopic               string
	DBTransactionsConnection string
	MailDialerHost           string
	MailDialerPort           int
	MailDialerUsername       string
	MailDialerPassword       string
}

func NewConfig() *Config {
	port, _ := strconv.Atoi(os.Getenv("mailDialerPort"))

	config := &Config{
		KafkaTopic:               os.Getenv("KAFKA_TOPIC"),
		DBTransactionsConnection: os.Getenv("DB_TRANSACTIONS_CONNECTION"),
		MailDialerHost:           os.Getenv("mailDialerHost"),
		MailDialerPort:           port,
		MailDialerUsername:       os.Getenv("mailDialerUsername"),
		MailDialerPassword:       os.Getenv("mailDialerPassword"),
	}
	return config
}
