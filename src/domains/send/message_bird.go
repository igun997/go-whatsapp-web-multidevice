package send

type MessageBirdRequest struct {
	TemplateId        string                 `json:"template_id" form:"template_id"`
	TemplateVersionId string                 `json:"template_version_id" form:"template_version_id"`
	Data              map[string]interface{} `json:"data" form:"data"`
	Phone             string                 `json:"phone" form:"phone"`
}
