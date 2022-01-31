package src

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

func GetConfig(key string) interface{} {
	return viper.Get(key)
}
func GetConfigString(key string) string {
	return viper.Get(key).(string)
}
func GetConfigBool(key string) bool {
	return viper.Get(key).(bool)
}

func ViperConfigure() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	viper.SetDefault("AMQP_MANDATORY", false)
	viper.SetDefault("AMQP_IMMEDIATE", false)

	viper.SetDefault("AMQP_QUEUE_NAME", "sms")
	viper.SetDefault("AMQP_QUEUE_DURABLE", true)
	viper.SetDefault("AMQP_QUEUE_AUTO_DELETE", false)
	viper.SetDefault("AMQP_QUEUE_EXCLUSIVE", false)
	viper.SetDefault("AMQP_QUEUE_NO_WAIT", false)

	viper.SetDefault("AMQP_CONNECTIONSTRING", "amqp://guest:guest@127.0.0.1")

	viper.SetDefault("REQUEST_URL", "http://127.0.0.1")
	viper.SetDefault("REQUEST_ADDRESS_KEY", "address")
	viper.SetDefault("REQUEST_MESSAGE_KEY", "message")

	viper.SetDefault("MESSAGE_FOOTER", "")

	_, ok := viper.Get("REQUEST_URL").(string)
	if !ok {
		log.Fatalf("REQUEST_URL is required")
	}
	_, ok = viper.Get("AMQP_QUEUE_NAME").(string)
	if !ok {
		log.Fatalf("AMQP_QUEUE_NAME is required")
	}
}

func CreateHttp() *resty.Client {
	return resty.New().SetHeader("Content-Type", "application/json")

}
