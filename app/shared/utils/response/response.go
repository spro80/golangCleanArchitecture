package shared_utils_response

type ResponseInterface interface {
	HandlerCreateResponseSuccess(status int, statusDes string, data string, errorDes string) ResponseStruct
}

type ResponseStruct struct {
	Status            int    `json:"status"`
	StatusDescription string `json:"statusDescription"`
	Data              string `json:"data"`
	Error             string `json:"error"`
}

func NewResponse() *ResponseStruct {
	return &ResponseStruct{}
}

func (r *ResponseStruct) HandlerCreateResponseSuccess(status int, statusDes string, data string, errorDes string) ResponseStruct {
	response := ResponseStruct{
		Status:            status,
		StatusDescription: statusDes,
		Data:              data,
		Error:             errorDes,
	}
	return response
}
