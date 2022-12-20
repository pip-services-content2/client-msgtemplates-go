package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-content2/client-msgtemplates-go/version1"
	"github.com/stretchr/testify/assert"
)

type MessageTemplatesClientFixtureV1 struct {
	Client version1.IMessageTemplatesClientV1
}

var TEMPLATE1 = &version1.MessageTemplateV1{
	Id:      "1",
	Name:    "template1",
	From:    "",
	Subject: map[string]string{"en": "Text 1"},
	Text:    map[string]string{"en": "Text 1"},
	Html:    map[string]string{"en": "Text 1"},
	Status:  version1.TemplateStatusCompleted,
}
var TEMPLATE2 = &version1.MessageTemplateV1{
	Id:      "2",
	Name:    "template2",
	From:    "",
	Subject: map[string]string{"en": "Text 2"},
	Text:    map[string]string{"en": "Text 2"},
	Html:    map[string]string{"en": "Text 2"},
	Status:  version1.TemplateStatusCompleted,
}

func NewMessageTemplatesClientFixtureV1(client version1.IMessageTemplatesClientV1) *MessageTemplatesClientFixtureV1 {
	return &MessageTemplatesClientFixtureV1{
		Client: client,
	}
}

func (c *MessageTemplatesClientFixtureV1) clear() {
	page, _ := c.Client.GetTemplates(context.Background(), "", nil, nil)

	for _, v := range page.Data {
		template := v
		c.Client.DeleteTemplateById(context.Background(), "", template.Id)
	}
}

func (c *MessageTemplatesClientFixtureV1) TestCrudOperations(t *testing.T) {
	c.clear()
	defer c.clear()

	// Create one template
	template, err := c.Client.CreateTemplate(context.Background(), "", TEMPLATE1)
	assert.Nil(t, err)

	assert.NotNil(t, template)
	assert.Equal(t, template.Name, TEMPLATE1.Name)
	assert.Equal(t, template.Id, TEMPLATE1.Id)

	template1 := template

	// Create another template
	template, err = c.Client.CreateTemplate(context.Background(), "", TEMPLATE2)
	assert.Nil(t, err)

	assert.NotNil(t, template)
	assert.Equal(t, template.Name, TEMPLATE2.Name)
	assert.Equal(t, template.Id, TEMPLATE2.Id)

	//template2 := template

	// Get all templates
	page, err1 := c.Client.GetTemplates(context.Background(), "", nil, nil)
	assert.Nil(t, err1)

	assert.NotNil(t, page)
	assert.True(t, len(page.Data) >= 2)

	// Get template by name
	template, err = c.Client.GetTemplateByIdOrName(context.Background(), "", TEMPLATE1.Name)
	assert.Nil(t, err)

	assert.NotNil(t, template)

	// Update the template
	template1.Text["en"] = "Updated Content 1"
	template, err = c.Client.UpdateTemplate(context.Background(), "", template1)
	assert.Nil(t, err)

	assert.NotNil(t, template)
	assert.Equal(t, template.Text["en"], "Updated Content 1")
	assert.Equal(t, template.Name, template1.Name)

	template1 = template

	// Delete template
	_, err = c.Client.DeleteTemplateById(context.Background(), "", template1.Id)
	assert.Nil(t, err)

	// Try to get deleted template
	template, err = c.Client.GetTemplateById(context.Background(), "", template1.Id)
	assert.Nil(t, err)

	assert.Nil(t, template)
}
