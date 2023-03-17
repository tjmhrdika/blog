package utils

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    any    `json:"data"`
}

func BuildResponseSuccess(message string, data any) Response {
	res := Response{
		Status:  true,
		Message: message,
		Error:   "",
		Data:    data,
	}
	return res
}

func BuildResponseFailed(message string, err error) Response {
	res := Response{
		Status:  false,
		Message: message,
		Error:   err.Error(),
	}
	return res
}
