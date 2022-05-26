package helper

type ResponseServer struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    interface{}
}

type EmptyObj struct{}

func BuildResponse(success bool, message string, data interface{}) ResponseServer {
	res := ResponseServer{
		Success: success,
		Message: message,
		Data:    data,
	}
	return res
}
