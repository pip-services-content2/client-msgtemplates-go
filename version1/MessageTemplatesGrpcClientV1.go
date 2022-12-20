package version1

import (
	"context"

	"github.com/pip-services-content2/client-msgtemplates-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type MessageTemplatesGrpcClientV1 struct {
	*clients.GrpcClient
}

func NewMessageTemplatesGrpcClientV1() *MessageTemplatesGrpcClientV1 {
	return &MessageTemplatesGrpcClientV1{
		GrpcClient: clients.NewGrpcClient("msgtemplates_v1.MessageTemplates"),
	}
}

func (c *MessageTemplatesGrpcClientV1) GetTemplates(ctx context.Context, correlationId string, filter *data.FilterParams,
	paging *data.PagingParams) (result data.DataPage[*MessageTemplateV1], err error) {
	timing := c.Instrument(ctx, correlationId, "msgtemplates_v1.get_templates")
	defer timing.EndTiming(ctx, err)

	req := &protos.MessageTemplatePageRequest{
		CorrelationId: correlationId,
	}
	if filter != nil {
		req.Filter = filter.Value()
	}
	if paging != nil {
		req.Paging = &protos.PagingParams{
			Skip:  paging.GetSkip(0),
			Take:  (int32)(paging.GetTake(100)),
			Total: paging.Total,
		}
	}

	reply := new(protos.MessageTemplatePageReply)
	err = c.CallWithContext(ctx, "get_templates", correlationId, req, reply)
	if err != nil {
		return *data.NewEmptyDataPage[*MessageTemplateV1](), err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return *data.NewEmptyDataPage[*MessageTemplateV1](), err
	}

	result = toMessageTemplatePage(reply.Page)

	return result, nil
}

func (c *MessageTemplatesGrpcClientV1) GetTemplateById(ctx context.Context, correlationId string, id string) (result *MessageTemplateV1, err error) {
	timing := c.Instrument(ctx, correlationId, "msgtemplates_v1.get_template_by_id")
	defer timing.EndTiming(ctx, err)

	req := &protos.MessageTemplateIdRequest{
		CorrelationId: correlationId,
		TemplateId:    id,
	}

	reply := new(protos.MessageTemplateObjectReply)
	err = c.CallWithContext(ctx, "get_template_by_id", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toMessageTemplate(reply.Template)

	return result, nil
}

func (c *MessageTemplatesGrpcClientV1) GetTemplateByIdOrName(ctx context.Context, correlationId string, idOrName string) (result *MessageTemplateV1, err error) {
	timing := c.Instrument(ctx, correlationId, "msgtemplates_v1.get_template_by_id_or_name")
	defer timing.EndTiming(ctx, err)

	req := &protos.MessageTemplateNameRequest{
		CorrelationId: correlationId,
		Name:          idOrName,
	}

	reply := new(protos.MessageTemplateObjectReply)
	err = c.CallWithContext(ctx, "get_template_by_id_or_name", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toMessageTemplate(reply.Template)

	return result, nil
}

func (c *MessageTemplatesGrpcClientV1) CreateTemplate(ctx context.Context, correlationId string, template *MessageTemplateV1) (result *MessageTemplateV1, err error) {
	timing := c.Instrument(ctx, correlationId, "msgtemplates_v1.create_template")
	defer timing.EndTiming(ctx, err)

	req := &protos.MessageTemplateObjectRequest{
		CorrelationId: correlationId,
		Template:      fromMessageTemplate(template),
	}

	reply := new(protos.MessageTemplateObjectReply)
	err = c.CallWithContext(ctx, "create_template", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toMessageTemplate(reply.Template)

	return result, nil
}

func (c *MessageTemplatesGrpcClientV1) UpdateTemplate(ctx context.Context, correlationId string, template *MessageTemplateV1) (result *MessageTemplateV1, err error) {
	timing := c.Instrument(ctx, correlationId, "msgtemplates_v1.update_template")
	defer timing.EndTiming(ctx, err)

	req := &protos.MessageTemplateObjectRequest{
		CorrelationId: correlationId,
		Template:      fromMessageTemplate(template),
	}

	reply := new(protos.MessageTemplateObjectReply)
	err = c.CallWithContext(ctx, "update_template", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toMessageTemplate(reply.Template)

	return result, nil
}

func (c *MessageTemplatesGrpcClientV1) DeleteTemplateById(ctx context.Context, correlationId string, id string) (result *MessageTemplateV1, err error) {
	timing := c.Instrument(ctx, correlationId, "msgtemplates_v1.delete_template_by_id")
	defer timing.EndTiming(ctx, err)

	req := &protos.MessageTemplateIdRequest{
		CorrelationId: correlationId,
		TemplateId:    id,
	}

	reply := new(protos.MessageTemplateObjectReply)
	err = c.CallWithContext(ctx, "delete_template_by_id", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toMessageTemplate(reply.Template)

	return result, nil
}
