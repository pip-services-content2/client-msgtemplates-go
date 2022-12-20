package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type MessageTemplatesNullClientV1 struct {
}

func NewMessageTemplatesNullClientV1() *MessageTemplatesNullClientV1 {
	return &MessageTemplatesNullClientV1{}
}

func (c *MessageTemplatesNullClientV1) GetTemplates(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (result data.DataPage[*MessageTemplateV1], err error) {
	return *data.NewEmptyDataPage[*MessageTemplateV1](), nil
}

func (c *MessageTemplatesNullClientV1) GetTemplateById(ctx context.Context, correlationId string, id string) (result *MessageTemplateV1, err error) {
	return nil, nil
}

func (c *MessageTemplatesNullClientV1) GetTemplateByIdOrName(ctx context.Context, correlationId string, idOrName string) (result *MessageTemplateV1, err error) {
	return nil, nil
}

func (c *MessageTemplatesNullClientV1) CreateTemplate(ctx context.Context, correlationId string, template *MessageTemplateV1) (result *MessageTemplateV1, err error) {
	if template.Id == "" {
		template.Id = data.IdGenerator.NextLong()
	}

	return template, nil
}

func (c *MessageTemplatesNullClientV1) UpdateTemplate(ctx context.Context, correlationId string, template *MessageTemplateV1) (result *MessageTemplateV1, err error) {
	return template, nil
}

func (c *MessageTemplatesNullClientV1) DeleteTemplateById(ctx context.Context, correlationId string, id string) (result *MessageTemplateV1, err error) {
	return nil, nil
}
