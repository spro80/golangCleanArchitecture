package request_models

type User struct {
	Rut       string `json:"rut"`
	UserName  string `json:"userName"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Valid     bool   `json:"valid"`
	Profile   Profile
}

type Profile struct {
	ProfileId       int    `json:"profileId"`
	ProfileStatus   bool   `json:"profileStatus"`
	ProfileDateInit string `json:"profileDateInit"`
	ProfileDateEnd  string `json:"profileDateEnd"`
	ProfileAllTime  bool   `json:"profileAllTime"`
}
