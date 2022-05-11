# Iletimerkezi Sms Sender - Amqp Consumer

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

This repository contains a collection of Go programs and libraries that
demonstrate the language, standard libraries, and tools.

## Clone the project

```
git clone https://github.com/TA4APG/iletimerkezi-amqp-consumer
$ cd iletimerkezi-amqp-consumer
```

## Docker build

```
$ docker build . -t consumer-sms
```

## [Configuration File & Env Vars](config.yml)

- REQUEST_URL //required, default "https://api.iletimerkezi.com/v1/send-sms/json"
- ILETIMERKEZI_SENDER //required
- ILETIMERKEZI_IYS //required
- ILETIMERKEZI_KEY //required
- ILETIMERKEZI_HASH //required
- MODEL_ADDRESSES //required
- MODEL_MESSAGES //required
- AMQP_CONNECTIONSTRING //required, default "amqp://guest:guest@127.0.0.1"
- AMQP_QUEUE_NAME //required

- MESSAGE_FOOTER //optional
- ILETIMERKEZI_IYSLIST //optional

## Allowed config path (json or yaml)

- /run/secret/
- /etc/iletimerkezi-amqp/
- $HOME/.iletimerkezi-amqp
