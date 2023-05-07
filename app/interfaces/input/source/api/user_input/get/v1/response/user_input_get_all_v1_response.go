package user_input_get_all_v1_response

func NewUserResponse() *UserResponse {
	return &UserResponse{}
}

type UserResponse struct {
	IdUser    string `json:"idUser" validate:"required"`
	Rut       string `json:"rut" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	UserName  string `json:"userName" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Valid     bool   `json:"valid" validate:"required"`
}
