package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IMessageTemplatesClientV1 interface {
	GetTemplates(ctx context.Context, correlationId string, filter *data.FilterParams,
		paging *data.PagingParams) (result data.DataPage[*MessageTemplateV1], err error)

	GetTemplateById(ctx context.Context, correlationId string, id string) (result *MessageTemplateV1, err error)

	GetTemplateByIdOrName(ctx context.Context, correlationId string, idOrName string) (result *MessageTemplateV1, err error)

	CreateTemplate(ctx context.Context, correlationId string, template *MessageTemplateV1) (result *MessageTemplateV1, err error)

	UpdateTemplate(ctx context.Context, correlationId string, template *MessageTemplateV1) (result *MessageTemplateV1, err error)

	DeleteTemplateById(ctx context.Context, correlationId string, id string) (result *MessageTemplateV1, err error)
}
