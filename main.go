package main

import (
	"crypto/tls"
	"encoding/json"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gomail/config"
	"github.com/gomail/src/domain"
	"github.com/gomail/src/email/services"
	"github.com/gomail/src/infra/database/gateway/gorm"
	"github.com/gomail/src/infra/database/gateway/gorm/repositories"
	"github.com/gomail/src/infra/kafka"
	goMail "gopkg.in/mail.v2"
	"log"
)

func main() {

	config := config.NewConfig()

	db, err := gorm.NewGormDB(config.DBTransactionsConnection)
	if err != nil {
		log.Fatal(err)
	}

	var emailCh = make(chan domain.Email)
	var msgchan = make(chan *ckafka.Message)

	d := goMail.NewDialer(
		config.MailDialerHost,
		config.MailDialerPort,
		config.MailDialerUsername,
		config.MailDialerPassword,
	)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	mailRepo := repositories.NewEmailRepo(db)
	es := services.NewMailSender(mailRepo)
	es.From = "rhiad.ciccoli@gmail.com"
	es.Dialer = d

	go es.Send(emailCh)

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9094",
		"client.id":         "emailapp",
		"group.id":          "emailapp",
	}

	topics := []string{config.KafkaTopic}

	consumer := kafka.NewConsumer(configMap, topics)

	go consumer.Consume(msgchan)

	for msg := range msgchan {
		var input domain.Email
		json.Unmarshal(msg.Value, &input)
		emailCh <- input
	}
}
