package response

type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Success(message string, data any) *SuccessResponse {
	return &SuccessResponse{Success: true, Message: message, Data: data}
}
