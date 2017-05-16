package main

import (
	"net/http"
	"os"

	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/brokerapi"
	"github.com/sahilm/redis-service-broker/broker"
)

func main() {
	brokerLogger := lager.NewLogger("redis-broker")
	brokerLogger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
	brokerLogger.RegisterSink(lager.NewWriterSink(os.Stderr, lager.ERROR))

	brokerCredentials := brokerapi.BrokerCredentials{
		Username: "username",
		Password: "password",
	}

	brokerLogger.Info("Starting Redis Broker")

	serviceBroker := &broker.RedisServiceBroker{}

	brokerAPI := brokerapi.New(serviceBroker, brokerLogger, brokerCredentials)
	http.Handle("/", brokerAPI)

	brokerLogger.Fatal("http-listen", http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
