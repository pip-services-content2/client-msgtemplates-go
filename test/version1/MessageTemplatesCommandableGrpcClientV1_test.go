package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-content2/client-msgtemplates-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type messageTemplatesCommandableGrpcClientV1Test struct {
	client  *version1.MessageTemplatesCommandableGrpcClientV1
	fixture *MessageTemplatesClientFixtureV1
}

func newMessageTemplatesCommandableGrpcClientV1Test() *messageTemplatesCommandableGrpcClientV1Test {
	return &messageTemplatesCommandableGrpcClientV1Test{}
}

func (c *messageTemplatesCommandableGrpcClientV1Test) setup(t *testing.T) *MessageTemplatesClientFixtureV1 {
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

	c.client = version1.NewMessageTemplatesCommandableGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewMessageTemplatesClientFixtureV1(c.client)

	return c.fixture
}

func (c *messageTemplatesCommandableGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableGrpcCrudOperations(t *testing.T) {
	c := newMessageTemplatesCommandableGrpcClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
