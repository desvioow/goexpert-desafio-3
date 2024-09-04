package main

import (
	"database/sql"
	"fmt"
	"github.com/desvioow/goexpert-desafio-3/configs"
	"github.com/desvioow/goexpert-desafio-3/internal/events/handler"
	"github.com/desvioow/goexpert-desafio-3/pkg/events"
	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()
	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})
}
func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
