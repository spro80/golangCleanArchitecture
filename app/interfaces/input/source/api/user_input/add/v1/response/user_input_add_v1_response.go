package user_input_add_v1_response

type UserAddResponse struct {
	ApiResponse[UserAddResponse]
	MatchedCount  int `json:"matchedCount" validate:"required"`
	ModifiedCount int `json:"modifiedCount" validate:"required"`
}

func (uar *UserAddResponse) MapToResponse() (map[string]interface{}, error) {
	return uar.ApiResponse.MapToResponse(*r)
}
