package main

import (
	"consumer-sms/src"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	. "github.com/ahmetb/go-linq/v3"
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

	if src.FailOnError(err, "Failed registration to amqp") {
		os.Exit(1)
	}

	mAdrs := src.GetConfigString("MODEL_ADDRESSES")
	mMsgs := src.GetConfigString("MODEL_MESSAGES")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			var m map[string]interface{}
			json.Unmarshal(d.Body, &m)

			//Single
			valAddress, err := src.NestedMapLookup(m, strings.Split(mAdrs, ".")...)
			src.FailOnError(err, "Hata")

			valMessage, err := src.NestedMapLookup(m, strings.Split(mMsgs, ".")...)

			src.FailOnError(err, "Hata")

			var targetAddresses []string
			From(valAddress).ToSlice(&targetAddresses)

			for _, v := range valMessage.([]interface{}) {
				message := v.(string)
				body := MakeBody(targetAddresses, message)
				result := src.SendSms(body)

				logStr := fmt.Sprintf("Message: %s, target: %s", message, strings.Join(targetAddresses, ", "))
				if result == true {
					fmt.Printf("Success: %s", logStr)
				}
			}
			src.Channel.Ack(d.DeliveryTag, false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
	os.Exit(0)
}
