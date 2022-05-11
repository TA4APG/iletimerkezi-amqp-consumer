package src

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

var HttpClient *resty.Client

func HttpConfigure() {
	HttpClient = createHttp()
	HttpClient.SetRootCertificate("./certs/root-ca.pem")
}
func createHttp() *resty.Client {
	return resty.New().SetHeader("Content-Type", "application/json")
}

func SendSms(req *IletiMerkeziSms) bool {
	request := HttpClient.R()
	byteArr, err := json.Marshal(req)
	if FailOnError(err, "Serializasyon hatasi") {
		return false
	}

	request.SetBody(byteArr)
	url := GetConfigString("REQUEST_URL")
	var resp *resty.Response

	resp, err = request.Post(url)

	if FailOnError(err, "Sms gonderirken hata") {
		return false
	}

	fmt.Println(string(resp.Body()))
	return true
}
