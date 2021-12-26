package main

import (
	"golang_test/module/amqp/util"
	"log"
	"os"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
)

/**test case
./emit_log_topic "kern.critical" "kern.critical: worker message 01......" && \
./emit_log_topic "j.yellow" "j.yellow: worker message 02......" && \
./emit_log_topic "kern.critical" "kern.critical: worker message 03......" && \
./emit_log_topic "green.critical" "green.critical: worker message 04......" && \
./emit_log_topic "kern.hello" "kern.hello: worker message 05......" && \
./emit_log_topic "kern.critical" "kern.critical: worker message 06......"
*/

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	util.FailOnError(err, "Failed to declare an exchange")

	body := bodyFrom(os.Args)
	err = ch.Publish(
		"logs_topic",          // exchange
		severityFrom(os.Args), // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	util.FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "anonymous.info"
	} else {
		s = os.Args[1]
	}
	return s
}
