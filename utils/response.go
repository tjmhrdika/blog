package utils

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   error  `json:"error"`
	Data    any    `json:"data"`
}

type EmptyObj struct{}

func BuildResponseSuccess(message string, data any) Response {
	res := Response{
		Status:  true,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

func BuildResponseFailed(message string, err error) Response {
	res := Response{
		Status:  false,
		Message: message,
		Error:   err,
	}
	return res
}
