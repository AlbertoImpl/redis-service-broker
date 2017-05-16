package broker_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sahilm/redis-service-broker/broker"
)

var _ = Describe("Broker", func() {
	Context("Services", func() {
		It("should return the services catalog", func() {
			redisServiceBroker := broker.RedisServiceBroker{}
			Expect(redisServiceBroker.Services(nil)).NotTo(Equal(nil))
		})
	})
})
