package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
)

type MessageTemplatesResolverV1 struct {
	client    IMessageTemplatesClientV1
	config    *cconf.ConfigParams
	templates map[string]*MessageTemplateV1
}

func NewMessageTemplatesResolverV1() *MessageTemplatesResolverV1 {
	return NewMessageTemplatesResolverV1WithParams(nil, nil)
}

func NewMessageTemplatesResolverV1WithParams(config *cconf.ConfigParams, references cref.IReferences) *MessageTemplatesResolverV1 {
	c := &MessageTemplatesResolverV1{
		config:    cconf.NewEmptyConfigParams(),
		templates: make(map[string]*MessageTemplateV1, 0),
	}

	if config != nil {
		c.Configure(context.Background(), config)
	}

	if references != nil {
		c.SetReferences(context.Background(), references)
	}

	return c
}

func NewMessageTemplatesResolverV1FromTuples(tuples ...any) *MessageTemplatesResolverV1 {
	result := NewMessageTemplatesResolverV1()
	if len(tuples) == 0 {
		return result
	}

	for index := 0; index < len(tuples); index += 2 {
		if index+1 >= len(tuples) {
			break
		}

		name := convert.StringConverter.ToString(tuples[index])
		template := tuples[index+1]

		result.Put(name, template)
	}
	return result
}

func (c *MessageTemplatesResolverV1) Configure(ctx context.Context, config *cconf.ConfigParams) {
	c.config = config.GetSection("message_templates")
}

func (c *MessageTemplatesResolverV1) SetReferences(ctx context.Context, references cref.IReferences) {
	res := references.GetOneOptional(cref.NewDescriptor("service-msgtemplates", "client", "*", "*", "1.0"))

	if res != nil {
		c.client = res.(IMessageTemplatesClientV1)
	}
}

func (c *MessageTemplatesResolverV1) Put(name string, template any) {
	c.config.Put(name, template)
}

func (c *MessageTemplatesResolverV1) retriveTemplate(ctx context.Context, name string, templateName string) (*MessageTemplateV1, error) {
	if c.client == nil {
		return nil, nil
	}

	template, err := c.client.GetTemplateByIdOrName(ctx, "msg_resolve", templateName)

	if err != nil {
		return nil, err
	}

	if template != nil {
		c.templates[name] = template
	}

	return template, nil
}

func (c *MessageTemplatesResolverV1) createTemplate(name string, config *cconf.ConfigParams) *MessageTemplateV1 {
	if config == nil || len(config.Keys()) == 0 {
		return nil
	}

	subject, subjectOk := config.GetAsObject("subject")
	text, textOk := config.GetAsObject("text")
	html, htmlOk := config.GetAsObject("html")

	if !(subjectOk && textOk && htmlOk) {
		return nil
	}

	template := &MessageTemplateV1{
		Name:    name,
		Subject: subject.(map[string]string),
		Text:    text.(map[string]string),
		Html:    html.(map[string]string),
	}

	c.templates[name] = template

	return template
}

func (c *MessageTemplatesResolverV1) Resolve(ctx context.Context, name string) (*MessageTemplateV1, error) {
	if name == "" {
		panic("Dependency name cannot be empty")
	}

	// Retrieve template first
	if template, ok := c.templates[name]; ok {
		return template, nil
	}

	// Get configuration
	config := c.config.GetSection(name)
	templateName, tNameOk := c.config.GetAsNullableString(name)
	if !tNameOk {
		templateName, _ = config.GetAsNullableString("template")
	}

	if templateName != "" {
		return c.retriveTemplate(ctx, name, templateName)
	} else {
		return c.createTemplate(name, config), nil
	}
}
