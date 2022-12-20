package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type MessageTemplatesCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewMessageTemplatesCommandableHttpClientV1() *MessageTemplatesCommandableHttpClientV1 {
	return NewMessageTemplatesCommandableHttpClientV1WithConfig(nil)
}

func NewMessageTemplatesCommandableHttpClientV1WithConfig(config *cconf.ConfigParams) *MessageTemplatesCommandableHttpClientV1 {
	c := &MessageTemplatesCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/message_templates"),
	}

	if config != nil {
		c.Configure(context.Background(), config)
	}

	return c
}

func (c *MessageTemplatesCommandableHttpClientV1) GetTemplates(ctx context.Context, correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (result cdata.DataPage[*MessageTemplateV1], err error) {
	res, err := c.CallCommand(ctx, "get_templates", correlationId, cdata.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	))

	if err != nil {
		return *cdata.NewEmptyDataPage[*MessageTemplateV1](), err
	}

	return clients.HandleHttpResponse[cdata.DataPage[*MessageTemplateV1]](res, correlationId)
}

func (c *MessageTemplatesCommandableHttpClientV1) GetTemplateById(ctx context.Context, correlationId string, id string) (result *MessageTemplateV1, err error) {
	res, err := c.CallCommand(ctx, "get_template_by_id", correlationId, cdata.NewAnyValueMapFromTuples(
		"template_id", id,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*MessageTemplateV1](res, correlationId)
}

func (c *MessageTemplatesCommandableHttpClientV1) GetTemplateByIdOrName(ctx context.Context, correlationId string, idOrName string) (result *MessageTemplateV1, err error) {
	res, err := c.CallCommand(ctx, "get_template_by_id_or_name", correlationId, cdata.NewAnyValueMapFromTuples(
		"id_or_name", idOrName,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*MessageTemplateV1](res, correlationId)
}

func (c *MessageTemplatesCommandableHttpClientV1) CreateTemplate(ctx context.Context, correlationId string, template *MessageTemplateV1) (result *MessageTemplateV1, err error) {
	res, err := c.CallCommand(ctx, "create_template", correlationId, cdata.NewAnyValueMapFromTuples(
		"template", template,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*MessageTemplateV1](res, correlationId)
}

func (c *MessageTemplatesCommandableHttpClientV1) UpdateTemplate(ctx context.Context, correlationId string, template *MessageTemplateV1) (result *MessageTemplateV1, err error) {
	res, err := c.CallCommand(ctx, "update_template", correlationId, cdata.NewAnyValueMapFromTuples(
		"template", template,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*MessageTemplateV1](res, correlationId)
}

func (c *MessageTemplatesCommandableHttpClientV1) DeleteTemplateById(ctx context.Context, correlationId string, id string) (result *MessageTemplateV1, err error) {
	res, err := c.CallCommand(ctx, "delete_template_by_id", correlationId, cdata.NewAnyValueMapFromTuples(
		"template_id", id,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*MessageTemplateV1](res, correlationId)
}
