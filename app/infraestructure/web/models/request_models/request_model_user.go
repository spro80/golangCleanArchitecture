package request_models

type User struct {
	Rut       string `json:"rut"`
	UserName  string `json:"userName"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Valid     bool   `json:"valid"`
}
