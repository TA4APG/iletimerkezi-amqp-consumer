# Iletimerkezi Sms Sender - Amqp Consumer 

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

## Clone the project

```
$ git clone https://github.com/TA4APG/iletimerkezi-amqp-consumer
$ cd iletimerkezi-amqp-consumer
```
## Docker build

```
$ docker build . -t consumer-sms
```

## [Configuration File & Env Vars](config.yml)

* AMQP_QUEUE_NAME //required
* MESSAGE_FOOTER //optional
* STATIC_PARAMS //required
* AMQP_CONNECTIONSTRING //default "amqp://guest:guest@127.0.0.1"
