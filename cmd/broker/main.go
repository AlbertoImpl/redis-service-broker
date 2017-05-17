package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/brokerapi"
	"github.com/sahilm/redis-service-broker/broker"
)

func main() {
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	logger := createLogger("redis-broker")

	brokerAPI := createBroker(createCredentials("username", "password"), logger)

	h := &http.Server{Addr: ":" + os.Getenv("PORT"), Handler: brokerAPI}

	go func() {
		logger.Info("Starting Redis Broker")
		logger.Fatal("http-listen", h.ListenAndServe())
	}()

	<-stop

	logger.Info("\nShutting down the server...")
	h.Shutdown(context.Background())
	logger.Info("Server gracefully stopped")
}

func createLogger(component string) lager.Logger {
	brokerLogger := lager.NewLogger(component)
	brokerLogger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
	brokerLogger.RegisterSink(lager.NewWriterSink(os.Stderr, lager.ERROR))
	return brokerLogger
}

func createCredentials(username string, password string) brokerapi.BrokerCredentials {
	return brokerapi.BrokerCredentials{
		Username: username,
		Password: password,
	}
}

func createBroker(credentials brokerapi.BrokerCredentials, logger lager.Logger) http.Handler {
	return brokerapi.New(&broker.RedisServiceBroker{}, logger, credentials)
}
