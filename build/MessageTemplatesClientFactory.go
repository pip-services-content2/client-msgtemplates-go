package build

import (
	clients1 "github.com/pip-services-content2/client-msgtemplates-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type MessageTemplatesClientFactory struct {
	*cbuild.Factory
}

func NewMessageTemplatesClientFactory() *MessageTemplatesClientFactory {
	c := &MessageTemplatesClientFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-msgtemplates", "client", "null", "*", "1.0")
	mockClientDescriptor := cref.NewDescriptor("service-msgtemplates", "client", "mock", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-msgtemplates", "client", "commandable-http", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("service-msgtemplates", "client", "grpc", "*", "1.0")
	cmdGrpcClientDescriptor := cref.NewDescriptor("service-msgtemplates", "client", "commandable-grpc", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewMessageTemplatesNullClientV1)
	c.RegisterType(mockClientDescriptor, clients1.NewMessageTemplatesMockClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewMessageTemplatesCommandableHttpClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewMessageTemplatesGrpcClientV1)
	c.RegisterType(cmdGrpcClientDescriptor, clients1.NewMessageTemplatesCommandableGrpcClientV1)

	return c
}
