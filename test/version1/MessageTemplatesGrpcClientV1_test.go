package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-content2/client-msgtemplates-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type messageTemplatesGrpcClientV1Test struct {
	client  *version1.MessageTemplatesGrpcClientV1
	fixture *MessageTemplatesClientFixtureV1
}

func newMessageTemplatesGrpcClientV1Test() *messageTemplatesGrpcClientV1Test {
	return &messageTemplatesGrpcClientV1Test{}
}

func (c *messageTemplatesGrpcClientV1Test) setup(t *testing.T) *MessageTemplatesClientFixtureV1 {
	var GRPC_HOST = os.Getenv("GRPC_HOST")
	if GRPC_HOST == "" {
		GRPC_HOST = "localhost"
	}
	var GRPC_PORT = os.Getenv("GRPC_PORT")
	if GRPC_PORT == "" {
		GRPC_PORT = "8090"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", GRPC_HOST,
		"connection.port", GRPC_PORT,
	)

	c.client = version1.NewMessageTemplatesGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewMessageTemplatesClientFixtureV1(c.client)

	return c.fixture
}

func (c *messageTemplatesGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestGrpcCrudOperations(t *testing.T) {
	c := newMessageTemplatesGrpcClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
