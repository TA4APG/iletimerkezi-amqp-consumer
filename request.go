package main

import "consumer-sms/src"

var bodyBase src.IletiMerkeziSms
var footerMsg string

func RequestSetup() {
	sender := src.GetConfigString("ILETIMERKEZI_SENDER")
	iys := src.GetConfigString("ILETIMERKEZI_IYS")
	iyslist := src.GetConfig("ILETIMERKEZI_IYSLIST")
	key := src.GetConfigString("ILETIMERKEZI_KEY")
	hash := src.GetConfigString("ILETIMERKEZI_HASH")

	footerMsg = src.GetConfigString("MESSAGE_FOOTER")

	bodyBase = src.IletiMerkeziSms{Request: src.IMRequest{
		Authentication: src.IMAuthentication{
			Key:  key,
			Hash: hash,
		},
		Order: src.IMOrder{
			Sender: sender,
			Iys:    iys,
		},
	}}
	if iyslist != nil {
		bodyBase.Request.Order.IysList = iyslist.(string)
	}
}

// Http requestin bodysine parametreleri girer ve mesaja footer ekler
func MakeBody(addresses []string, message string) *src.IletiMerkeziSms {
	body := bodyBase
	body.Request.Order.Message.Receipents.Number = addresses
	body.Request.Order.Message.Text = message + footerMsg

	return &body
}
