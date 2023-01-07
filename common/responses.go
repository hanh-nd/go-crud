package common

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GetListResponse struct {
	Items      interface{} `json:"items"`
	TotalItems int64       `json:"totalItems"`
}

func NewSuccessResponse(data interface{}, code ...int) successResponse {
	statusCode := 200
	if len(code) > 0 {
		statusCode = code[0]
	}
	return successResponse{
		Code: statusCode,
		Data: data,
	}
}

func NewErrorResponse(code int, message ...string) errorResponse {
	statusMessage := "An error occurred"
	if len(message) > 0 {
		statusMessage = message[0]
	}
	return errorResponse{
		Code:    code,
		Message: statusMessage,
	}
}

func NewGetListResponse(items interface{}, totalItems int64) GetListResponse {
	return GetListResponse{
		Items:      items,
		TotalItems: totalItems,
	}
}
