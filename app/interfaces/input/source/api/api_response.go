package input_source_api

import "encoding/json"

type ApiResponseInterface[R any] interface {
	MapToResponse(response R) (map[string]interface{}, error)
}

type ApiResponse[R any] struct {
}

func NewApiResponse[R any]() *ApiResponse[R] {
	return &ApiResponse[R]{}
}

func (r *ApiResponse[R]) MapToResponse(response R) (map[string]interface{}, error) {
	var mapResponse map[string]interface{}
	responseByte, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseByte, &mapResponse)
	if err != nil {
		return nil, err
	}

	return mapResponse, nil
}
