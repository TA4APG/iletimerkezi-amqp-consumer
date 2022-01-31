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

func SendSms(req *IletiMerkeziSms) {
	request := HttpClient.R()
	byteArr, err := json.Marshal(req)
	request.SetBody(byteArr)
	url := GetConfigString("REQUEST_URL")
	resp, err := request.Post(url)
	FailOnError(err, "Sms gonderirken hata")
	// if err != nil {
	// 	log.Print(err)
	// 	os.Exit(1)
	// }

	fmt.Println(string(resp.Body()))
}
