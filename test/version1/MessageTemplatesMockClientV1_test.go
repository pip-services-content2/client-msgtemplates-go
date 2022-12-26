package test_version1

import (
	"testing"

	"github.com/pip-services-content2/client-msgtemplates-go/version1"
)

type messageTemplatesMockClientV1Test struct {
	client  *version1.MessageTemplatesMockClientV1
	fixture *MessageTemplatesClientFixtureV1
}

func newMessageTemplatesMockClientV1Test() *messageTemplatesMockClientV1Test {
	return &messageTemplatesMockClientV1Test{}
}

func (c *messageTemplatesMockClientV1Test) setup(t *testing.T) *MessageTemplatesClientFixtureV1 {
	c.client = version1.NewMessageTemplatesMockClientV1()

	c.fixture = NewMessageTemplatesClientFixtureV1(c.client)

	return c.fixture
}

func (c *messageTemplatesMockClientV1Test) teardown(t *testing.T) {
	c.client = nil
}

func TestMockOperations(t *testing.T) {
	c := newMessageTemplatesMockClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
