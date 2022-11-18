package interfaces

type UserEntityInterface interface {
	GetRut() string
	GetUserName() string
	GetPassword() string
	GetEmail() string
	GetFirstName() string
	GetLastName() string

	SetRut(rut string) error
	SetUserName(name string) error
	SetPassword(password string) error
	SetEmail(email string) error
	SetFirstName(firstName string) error
	SetLastName(lastName string) error
}
