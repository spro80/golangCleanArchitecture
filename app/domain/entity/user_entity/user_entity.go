package entity

type User struct {
	IdUser    string `json:"idUser" validate:"required"`
	Rut       string `json:"rut" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	UserName  string `json:"userName" validate:"required"`
	Password  string `json:"password" validate:"required"`
}
