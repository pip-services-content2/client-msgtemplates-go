package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-content2/client-msgtemplates-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type messageTemplatesCommandableHttpClientV1Test struct {
	client  *version1.MessageTemplatesCommandableHttpClientV1
	fixture *MessageTemplatesClientFixtureV1
}

func newMessageTemplatesCommandableHttpClientV1Test() *messageTemplatesCommandableHttpClientV1Test {
	return &messageTemplatesCommandableHttpClientV1Test{}
}

func (c *messageTemplatesCommandableHttpClientV1Test) setup(t *testing.T) *MessageTemplatesClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewMessageTemplatesCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewMessageTemplatesClientFixtureV1(c.client)

	return c.fixture
}

func (c *messageTemplatesCommandableHttpClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableHttpCrudOperations(t *testing.T) {
	c := newMessageTemplatesCommandableHttpClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
