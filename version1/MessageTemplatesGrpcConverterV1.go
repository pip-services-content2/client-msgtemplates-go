package version1

import (
	"encoding/json"

	"github.com/pip-services-content2/client-msgtemplates-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
)

func fromError(err error) *protos.ErrorDescription {
	if err == nil {
		return nil
	}

	desc := errors.ErrorDescriptionFactory.Create(err)
	obj := &protos.ErrorDescription{
		Type:          desc.Type,
		Category:      desc.Category,
		Code:          desc.Code,
		CorrelationId: desc.CorrelationId,
		Status:        convert.StringConverter.ToString(desc.Status),
		Message:       desc.Message,
		Cause:         desc.Cause,
		StackTrace:    desc.StackTrace,
		Details:       fromMap(desc.Details),
	}

	return obj
}

func toError(obj *protos.ErrorDescription) error {
	if obj == nil || (obj.Category == "" && obj.Message == "") {
		return nil
	}

	description := &errors.ErrorDescription{
		Type:          obj.Type,
		Category:      obj.Category,
		Code:          obj.Code,
		CorrelationId: obj.CorrelationId,
		Status:        convert.IntegerConverter.ToInteger(obj.Status),
		Message:       obj.Message,
		Cause:         obj.Cause,
		StackTrace:    obj.StackTrace,
		Details:       toMap(obj.Details),
	}

	return errors.ApplicationErrorFactory.Create(description)
}

func fromMap(val map[string]any) map[string]string {
	r := map[string]string{}

	for k, v := range val {
		r[k] = convert.StringConverter.ToString(v)
	}

	return r
}

func toMap(val map[string]string) map[string]any {
	r := map[string]any{}

	for k, v := range val {
		r[k] = v
	}

	return r
}

func toJson(value any) string {
	if value == nil {
		return ""
	}

	b, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(b[:])
}

func fromJson(value string) any {
	if value == "" {
		return nil
	}

	var m any
	json.Unmarshal([]byte(value), &m)
	return m
}

func fromMessageTemplate(template *MessageTemplateV1) *protos.MessageTemplate {
	if template == nil {
		return nil
	}

	obj := &protos.MessageTemplate{
		Id:      template.Id,
		Name:    template.Name,
		From:    template.From,
		Subject: template.Subject,
		Text:    template.Text,
		Html:    template.Html,
		Status:  template.Status,
	}

	return obj
}

func toMessageTemplate(obj *protos.MessageTemplate) *MessageTemplateV1 {
	if obj == nil {
		return nil
	}

	template := &MessageTemplateV1{
		Id:      obj.Id,
		Name:    obj.Name,
		From:    obj.From,
		Subject: obj.Subject,
		Text:    obj.Text,
		Html:    obj.Html,
		Status:  obj.Status,
	}

	return template
}

func fromMessageTemplatePage(page data.DataPage[*MessageTemplateV1]) *protos.MessageTemplatePage {
	obj := &protos.MessageTemplatePage{
		Total: int64(page.Total),
		Data:  make([]*protos.MessageTemplate, len(page.Data)),
	}

	for i, v := range page.Data {
		template := v
		obj.Data[i] = fromMessageTemplate(template)
	}

	return obj
}

func toMessageTemplatePage(obj *protos.MessageTemplatePage) data.DataPage[*MessageTemplateV1] {
	if obj == nil {
		return *data.NewEmptyDataPage[*MessageTemplateV1]()
	}

	templates := make([]*MessageTemplateV1, len(obj.Data))

	for i, v := range obj.Data {
		templates[i] = toMessageTemplate(v)
	}

	page := *data.NewDataPage(templates, int(obj.Total))

	return page
}
