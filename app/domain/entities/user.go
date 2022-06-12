package entities

type User struct {
	Id        int    `json: "id"`
	Rut       string `json: "rut"`
	Digit     int    `json: "digit"`
	Username  string `json: "userName"`
	password  string `json: "password"`
	Email     string `json: "email"`
	FirstName string `json: "firstName"`
	LastName  string `json: "lastName"`
}
