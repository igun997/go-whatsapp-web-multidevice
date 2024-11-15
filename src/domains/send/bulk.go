package send

type BulkMessageRequest struct {
	Message string   `json:"message" form:"message"`
	Phones  []string `json:"phones" form:"phones"`
}

type BulkMessageProgress struct {
	Total     int    `json:"total"`
	Completed int    `json:"completed"`
	Current   string `json:"current"`
	Status    string `json:"status"`
}

type BulkMessageResponse struct {
	MessageID string `json:"message_id"`
	Phone     string `json:"phone"`
	Status    string `json:"status"`
}
