package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"os"
)

var (
	rabbitHost = os.Getenv("RABBIT_HOST")
	rabbitPort = os.Getenv("RABBIT_PORT")
	rabbitUser = os.Getenv("RABBIT_USER")
	rabbitPassword = os.Getenv("RABBIT_PASSWORD")
)

func main() {
	router := httprouter.New()

	router.POST("/publish/:message", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		submit(w,r,p)
	})

	fmt.Println("Running ...")
	log.Fatal(http.ListenAndServe(":80", router))

}

func submit(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	message := p.ByName("message")

	fmt.Println("Received  message: "+ message)

	conn, err := amqp.Dial("amqp://"+ rabbitUser + ":" + rabbitPassword + "@" + rabbitHost + ":" + rabbitPort + "/")

	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to a channel", err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"publisher",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing {
			ContentType: "text/plain",
			Body: 	[]byte(message),
		})

	if err != nil {
		log.Fatalf("%s: %s", "Failed to publish a message", err)
	}

	log.Println("Success Publishing.")
}
