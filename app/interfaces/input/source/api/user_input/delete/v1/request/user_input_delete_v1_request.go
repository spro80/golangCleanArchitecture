package user_input_delete_v1_request

type UserDeleteRequest struct {
	Rut string `param:"userId" validate:"required"`
}
