package main

import "consumer-sms/src"

var bodyBase src.IletiMerkeziSms
var footerMsg string

func RequestSetup() {
	StaticParams, ok := src.GetConfig("STATIC_PARAMS").(map[string]interface{})
	footerMsg = src.GetConfigString("MESSAGE_FOOTER")
	if !ok {
		src.FailOnError(nil, "Parameters not found")
	}

	bodyBase = src.IletiMerkeziSms{Request: src.IMRequest{
		Authentication: src.IMAuthentication{
			Key:  StaticParams["key"].(string),
			Hash: StaticParams["hash"].(string),
		},
		Order: src.IMOrder{
			Sender: StaticParams["sender"].(string),
		},
	}}
	iys, ok := StaticParams["iys"]
	if ok {
		bodyBase.Request.Order.Iys = iys.(string)
	}
	iysList, ok := StaticParams["iysList"]
	if ok {
		bodyBase.Request.Order.IysList = iysList.(string)
	} else {
		bodyBase.Request.Order.IysList = ""
	}
}

func SendRequest(addresses []string, message string) *src.IletiMerkeziSms {

	body := bodyBase
	// body := bodyBase
	body.Request.Order.Message.Receipents.Number = addresses
	body.Request.Order.Message.Text = message + footerMsg

	return &body
}
