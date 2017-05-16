package broker

import (
	"context"
	"os"

	"github.com/pivotal-cf/brokerapi"
)

type RedisServiceBroker struct {
}

func (*RedisServiceBroker) Services(context context.Context) []brokerapi.Service {
	return []brokerapi.Service{
		{
			ID:          "318d7965-5e7d-4d43-bf52-317bb12cdb30",
			Name:        "fake-service",
			Description: "sdkl;fmkds",
			Bindable:    true,
			Plans:       plans(),
			Metadata: &brokerapi.ServiceMetadata{
				DisplayName:         "sd;lfk",
				LongDescription:     "sdfdsf",
				DocumentationUrl:    "http://example.com",
				SupportUrl:          "http://example.com",
				ImageUrl:            "http://example.com",
				ProviderDisplayName: "sdl;fkdsl;kf",
			},
			Tags: []string{
				"redis", "pivotal",
			},
		},
	}
}

func plans() []brokerapi.ServicePlan {
	plan := brokerapi.ServicePlan{
		ID:          "d75982f4-15a5-435d-85f4-fb57117b35f5",
		Name:        "sdkfljds",
		Description: "sdklfjklsdjf",

		Metadata: &brokerapi.ServicePlanMetadata{
			DisplayName: "sd;lfl;dsk",
			Bullets: []string{
				"redis",
			},
		},
	}
	plans := []brokerapi.ServicePlan{plan}
	return plans
}

func (*RedisServiceBroker) Provision(context context.Context, instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	return brokerapi.ProvisionedServiceSpec{}, nil
}

func (*RedisServiceBroker) Deprovision(context context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	return brokerapi.DeprovisionServiceSpec{}, nil
}

func (*RedisServiceBroker) Bind(context context.Context, instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	binding := brokerapi.Binding{}

	credentialsMap := map[string]interface{}{
		"host":     os.Getenv("REDIS_HOST"),
		"port":     os.Getenv("REDIS_PORT"),
		"password": os.Getenv("REDIS_PASSWORD"),
	}

	binding.Credentials = credentialsMap

	return binding, nil
}

func (*RedisServiceBroker) Unbind(context context.Context, instanceID, bindingID string, details brokerapi.UnbindDetails) error {
	return nil
}

func (*RedisServiceBroker) Update(context context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	return brokerapi.UpdateServiceSpec{}, nil
}

func (*RedisServiceBroker) LastOperation(context context.Context, instanceID, operationData string) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, nil
}
