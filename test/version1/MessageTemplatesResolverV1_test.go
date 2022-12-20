package test_version1

import (
	"context"
	"testing"

	clients1 "github.com/pip-services-content2/client-msgtemplates-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/stretchr/testify/assert"
)

// func TestResolveHardcodedTemplate(t *testing.T) {
// 	resolver := clients1.NewMessageTemplatesResolverV1()

// 	resolver.Configure(context.Background(), config.NewConfigParamsFromTuples(
// 		"message_templates.template1.subject", "Subject1",
// 		"message_templates.template1.text", "Text1",
// 		"message_templates.template1.html", "Html1",
// 	))

// 	template, err := resolver.Resolve(context.Background(), "template1")

// 	assert.Nil(t, err)
// 	assert.NotNil(t, template)
// 	assert.Equal(t, template.Subject, "Subject") // TODO: What type of fields should be? any, map or MultiString?
// 	assert.Equal(t, template.Text, "Text1")
// 	assert.Equal(t, template.Html, "Html1")
// }

func TestResolveMissingTemplate(t *testing.T) {
	resolver := clients1.NewMessageTemplatesResolverV1()

	resolver.Configure(context.Background(), config.NewConfigParamsFromTuples(
		"message_templates.template1.subject", "Subject1",
		"message_templates.template1.text", "Text1",
		"message_templates.template1.html", "Html1",
	))

	template, err := resolver.Resolve(context.Background(), "template2")

	assert.Nil(t, err)
	assert.Nil(t, template)
}

func TestResolveTemplateById(t *testing.T) {
	resolver := clients1.NewMessageTemplatesResolverV1()

	resolver.Configure(context.Background(), config.NewConfigParamsFromTuples(
		"message_templates.template1", "123",
	))

	template, err := resolver.Resolve(context.Background(), "template2")

	assert.Nil(t, err)
	assert.Nil(t, template)
}
