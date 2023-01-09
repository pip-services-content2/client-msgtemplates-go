package version1

import (
	"context"
	"strings"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type MessageTemplatesMockClientV1 struct {
	templates []*MessageTemplateV1
}

func NewMessageTemplatesMockClientV1() *MessageTemplatesMockClientV1 {
	return &MessageTemplatesMockClientV1{
		templates: make([]*MessageTemplateV1, 0),
	}
}

func (c *MessageTemplatesMockClientV1) matchString(value string, search string) bool {
	if value == "" && search == "" {
		return true
	}
	if value == "" || search == "" {
		return false
	}
	return strings.Contains(strings.ToLower(value), strings.ToLower(search))
}

func (c *MessageTemplatesMockClientV1) matchMultilanguageString(value map[string]string, search string) bool {
	for _, text := range value {
		if c.matchString(text, search) {
			return true
		}
	}

	return false
}

func (c *MessageTemplatesMockClientV1) matchSearch(item *MessageTemplateV1, search string) bool {
	search = strings.ToLower(search)
	if c.matchString(item.Name, search) {
		return true
	}
	if c.matchMultilanguageString(item.Subject, search) {
		return true
	}
	if c.matchMultilanguageString(item.Text, search) {
		return true
	}
	if c.matchMultilanguageString(item.Html, search) {
		return true
	}
	if c.matchString(item.Status, search) {
		return true
	}
	return false
}

func (c *MessageTemplatesMockClientV1) composeFilter(filter *data.FilterParams) func(*MessageTemplateV1) bool {
	if filter == nil {
		filter = data.NewEmptyFilterParams()
	}

	search, searchOk := filter.GetAsNullableString("search")
	id, idOk := filter.GetAsNullableString("id")
	status, statusOk := filter.GetAsNullableString("status")
	name, nameOk := filter.GetAsNullableString("name")

	return func(item *MessageTemplateV1) bool {
		if idOk && item.Id != id {
			return false
		}
		if nameOk && item.Name != name {
			return false
		}
		if statusOk && item.Status != status {
			return false
		}
		if searchOk && !c.matchSearch(item, search) {
			return false
		}
		return true
	}
}

func (c *MessageTemplatesMockClientV1) GetTemplates(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (result data.DataPage[*MessageTemplateV1], err error) {
	filterFunc := c.composeFilter(filter)

	items := make([]*MessageTemplateV1, 0)
	for _, v := range c.templates {
		item := v
		if filterFunc(item) {
			items = append(items, item)
		}
	}
	return *data.NewDataPage(items, len(c.templates)), nil
}

func (c *MessageTemplatesMockClientV1) GetTemplateById(ctx context.Context, correlationId string, id string) (result *MessageTemplateV1, err error) {
	for _, v := range c.templates {
		if v.Id == id {
			buf := *v
			result = &buf
			break
		}
	}
	return result, nil
}

func (c *MessageTemplatesMockClientV1) GetTemplateByIdOrName(ctx context.Context, correlationId string, idOrName string) (result *MessageTemplateV1, err error) {
	for _, v := range c.templates {
		if v.Id == idOrName || v.Name == idOrName {
			buf := *v
			result = &buf
			break
		}
	}
	return result, nil
}

func (c *MessageTemplatesMockClientV1) CreateTemplate(ctx context.Context, correlationId string, template *MessageTemplateV1) (result *MessageTemplateV1, err error) {
	if template.Id == "" {
		template.Id = data.IdGenerator.NextLong()
	}

	if template.Status == "" {
		template.Status = TemplateStatusNew
	}
	buf := *template
	c.templates = append(c.templates, &buf)

	return template, nil
}

func (c *MessageTemplatesMockClientV1) UpdateTemplate(ctx context.Context, correlationId string, template *MessageTemplateV1) (result *MessageTemplateV1, err error) {
	if template == nil {
		return nil, nil
	}

	var index = -1
	for i, v := range c.templates {
		if v.Id == template.Id {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}

	buf := *template
	c.templates[index] = &buf
	return template, nil
}

func (c *MessageTemplatesMockClientV1) DeleteTemplateById(ctx context.Context, correlationId string, id string) (result *MessageTemplateV1, err error) {
	var index = -1
	for i, v := range c.templates {
		if v.Id == id {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}

	var item = c.templates[index]
	if index < len(c.templates) {
		c.templates = append(c.templates[:index], c.templates[index+1:]...)
	} else {
		c.templates = c.templates[:index]
	}
	return item, nil
}
