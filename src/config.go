package src

import (
	"fmt"

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

	viper.SetConfigName("iletimerkezi-worker")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/run/secrets/")
	viper.AddConfigPath("/etc/iletimerkezi-worker/")
	viper.AddConfigPath("$HOME/.iletimerkezi-worker")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.ReadInConfig()
	err := viper.MergeInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	viper.SetDefault("AMQP_MANDATORY", false)
	viper.SetDefault("AMQP_IMMEDIATE", false)

	viper.SetDefault("AMQP_QUEUE_DURABLE", true)
	viper.SetDefault("AMQP_QUEUE_AUTO_DELETE", false)
	viper.SetDefault("AMQP_QUEUE_EXCLUSIVE", false)
	viper.SetDefault("AMQP_QUEUE_NO_WAIT", false)

	viper.SetDefault("REQUEST_METHODS", "POST")
	viper.SetDefault("REQUEST_ADDRESS_KEY", "address")
	viper.SetDefault("REQUEST_MESSAGE_KEY", "message")

	viper.SetDefault("MESSAGE_FOOTER", "")
	viper.SetDefault("REQUEST_URL", "https://api.iletimerkezi.com/v1/send-sms/json")

	RequireConf("AMQP_CONNECTIONSTRING")
	RequireConf("AMQP_QUEUE_NAME")
	RequireConf("MODEL_ADDRESSES")
	RequireConf("MODEL_MESSAGES")
	RequireConf("ILETIMERKEZI_SENDER")
	RequireConf("ILETIMERKEZI_IYS")
	// RequireConf("ILETIMERKEZI_IYSLIST")
	RequireConf("ILETIMERKEZI_KEY")
	RequireConf("ILETIMERKEZI_HASH")
}

func RequireConf(key string) {
	_, ok := viper.Get(key).(string)
	if !ok {
		panic(fmt.Errorf("Parameter has required: %s \n", key))
	}
}

func CreateHttp() *resty.Client {
	return resty.New().SetHeader("Content-Type", "application/json")

}
