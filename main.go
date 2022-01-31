package main

import (
	"consumer-sms/src"
	"encoding/json"
	"log"
	"os"
)

func main() {
	src.ViperConfigure()
	src.AmqpConfigure()
	src.HttpConfigure()
	RequestSetup()

	src.Channel.Qos(5, 0, false)
	msgs, err := src.Channel.Consume(
		src.GetConfigString("AMQP_QUEUE_NAME"), // queue
		"@pytz0ne",                             // consumer
		false,                                  // auto-ack
		false,                                  // exclusive
		false,                                  // no-local
		false,                                  // no-wait
		nil,                                    // args
	)
	src.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var m = src.Message{}
			json.Unmarshal(d.Body, &m)
			for _, msg := range m.Messages {
				req := SendRequest(m.Addresses, msg)
				src.SendSms(req)
			}

			src.Channel.Ack(d.DeliveryTag, false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
	os.Exit(0)
}
